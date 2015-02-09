package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/fsouza/go-dockerclient"
)

var (
	endpoint = "unix:///docker.sock"
	client   *docker.Client
)

func init() {
	if os.Getenv("DOCKER_ENDPOINT") != "" {
		endpoint = os.Getenv("DOCKER_ENDPOINT")
	}

	// Try to set up the client!
	dc, err := docker.NewClient(endpoint)
	if err != nil {
		fmt.Println("Unable to contact docker daemon at:", endpoint)
		os.Exit(1)
	}
	client = dc
}

func main() {
	mux := http.NewServeMux()
	mux.Handle("/", http.FileServer(http.Dir("site")))
	mux.HandleFunc("/containers", ListContainers)
	mux.HandleFunc("/container/", ContainerInfo)
	mux.HandleFunc("/dockerinfo", DockerServerInfo)

	fmt.Println("Starting http server...")
	err := http.ListenAndServe(":3000", mux)

	if err != nil {
		fmt.Println("Unable to start http server:", err)
		os.Exit(1)
	}
}

type Container struct {
	ID      string
	Image   string
	ImageID string
	Name    string
	Status  string
	VHost   string
}

func ListContainers(w http.ResponseWriter, r *http.Request) {
	// Let's get a list of containers
	containers, err := client.ListContainers(docker.ListContainersOptions{All: true})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	ci := make([]*Container, 0)
	for _, cv := range containers {
		c, err := client.InspectContainer(cv.ID)
		if err != nil {
			fmt.Println("Error inspecting container:", err)
			continue
		}

		// Check for VHOST environment variable!
		vhost := ""
		for _, e := range c.Config.Env {
			if strings.HasPrefix(e, "VIRTUAL_HOST") {
				vhost = strings.TrimPrefix(e, "VIRTUAL_HOST=")
			}
		}

		co := &Container{
			ID:      c.ID,
			Image:   c.Config.Image,
			ImageID: c.Image,
			Name:    strings.TrimPrefix(c.Name, "/"),
			Status:  cv.Status,
			VHost:   vhost,
		}

		ci = append(ci, co)
	}

	json.NewEncoder(w).Encode(ci)
}

func ContainerInfo(w http.ResponseWriter, r *http.Request) {
	parts := strings.Split(r.URL.RequestURI(), "/")
	id := parts[len(parts)-1]

	c, err := client.InspectContainer(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(c)
}

func DockerServerInfo(w http.ResponseWriter, r *http.Request) {
	info, err := client.Info()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(info)
}
