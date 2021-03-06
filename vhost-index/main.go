package main

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"math/rand"
	"net/http"
	"os"
	"strings"

	"github.com/codegangsta/negroni"
	"github.com/fsouza/go-dockerclient"
	"github.com/gorilla/mux"
	ngzip "github.com/phyber/negroni-gzip/gzip"
)

var (
	client *docker.Client
)

func init() {
	// Try to set up the client!
	dc, err := getDockerClient()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	client = dc
}

func main() {
	n := negroni.New()
	n.Use(ngzip.Gzip(ngzip.DefaultCompression))
	n.Use(negroni.NewStatic(http.Dir("public")))

	// Set up a Basic Authentication handler unless explicitly disabled
	if os.Getenv("AUTH_OFF") != strings.ToLower("true") {
		ah := NewAuthHandler()
		var user string
		if os.Getenv("AUTH_USER") != "" {
			user = os.Getenv("AUTH_USER")
		} else {
			user = "vadmin"
		}
		var pass string
		if os.Getenv("AUTH_PASS") != "" {
			pass = os.Getenv("AUTH_PASS")
		} else {
			src := "abcdefghjkmnpqrstuvwxyzABCDEFGHJKLMNPQRSTUVWXYZ1234567890"
			pwb := make([]byte, 12)
			for i := range pwb {
				pwb[i] = src[rand.Int()%len(src)]
			}
			pass = string(pwb)
		}
		fmt.Printf("Auth User: %s\nAuth Pass: %s\n", user, pass)
		ah.AddUser(user, pass)
		n.Use(ah)
	}

	// Gotta have a router
	r := mux.NewRouter()
	r.HandleFunc("/containers/list", ContainerList)
	r.HandleFunc("/containers/{id}/info", ContainerInfo)
	r.HandleFunc("/containers/{id}/start", ContainerStart)
	r.HandleFunc("/containers/{id}/stop", ContainerStop)
	r.HandleFunc("/containers/{id}/kill", ContainerKill)
	r.HandleFunc("/containers/{id}/restart", ContainerRestart)
	r.HandleFunc("/images/list", ImageList)
	r.HandleFunc("/images/{id}/info", ImageInfo)
	r.HandleFunc("/dockerinfo", DockerServerInfo)
	n.UseHandler(r)

	// So `docker logs` shows something
	fmt.Println("Starting VHost server.")
	err := http.ListenAndServe(":3000", n)

	if err != nil {
		fmt.Println("Unable to start http server:", err)
		os.Exit(1)
	}
}

// Simple authentication for a simple application.
// AuthHandler contains users we trust.
type AuthHandler struct {
	// A string of users that are authorized to access the application.
	// Users[username] = base64.StdEncoding.EncodeToString([]byte(username+":"+password))
	Users map[string]string
}

// Create a new (empty) AuthHandler
func NewAuthHandler() *AuthHandler {
	a := &AuthHandler{
		Users: make(map[string]string),
	}
	return a
}

// Add a user that's allowed to log in
func (a *AuthHandler) AddUser(name, pass string) {
	a.Users[name] = base64.StdEncoding.EncodeToString([]byte(name + ":" + pass))
}

func (a *AuthHandler) ServeHTTP(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	h := r.Header.Get("Authorization")
	for _, as := range a.Users {
		if "Basic "+as == h {
			next(w, r)
			return
		}
	}
	// Still here? Let's go ahead and prompt for authorization
	w.Header().Set("WWW-Authenticate", "Basic realm=\"VHost Index\"")
	http.Error(w, "Not Authorized", http.StatusUnauthorized)
}

func DockerServerInfo(w http.ResponseWriter, r *http.Request) {
	info, err := client.Info()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(info)
}

func getDockerClient() (*docker.Client, error) {
	if envHost := os.Getenv("DOCKER_HOST"); envHost != "" {
		return docker.NewClient(envHost)
	}

	if envHost := os.Getenv("DOCKER_ENDPOINT"); envHost != "" {
		fmt.Println("WARNING: DOCKER_ENDPOINT will be deprecated in the future. Use the standard DOCKER_HOST environment variable going forward.")
		return docker.NewClient(envHost)
	}

	unixSockets := []string{"/docker.sock", "/run/docker.sock", "/var/run/docker.sock"}
	for _, s := range unixSockets {
		if _, err := os.Stat(s); err == nil {
			return docker.NewClient(fmt.Sprintf("unix://%s", s))
		}
	}

	// Still here? Not good.
	return nil, errors.New("Unable to find a docker endpoint to use.")
}
