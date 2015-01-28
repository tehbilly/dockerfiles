package main

import (
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
	http.HandleFunc("/", redirectToList)
	http.HandleFunc("/list", ListVHostContainers)

	fmt.Println("Starting http server...")
	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		fmt.Println("Unable to start http server:", err)
		os.Exit(1)
	}
}

func redirectToList(rw http.ResponseWriter, req *http.Request) {
	http.Redirect(rw, req, "/list", http.StatusTemporaryRedirect)
}

func ListVHostContainers(rw http.ResponseWriter, req *http.Request) {
	// Let's get a list of containers
	containers, err := client.ListContainers(docker.ListContainersOptions{All: true})
	if err != nil {
		http.Error(rw, "Error retrieving list of containers: "+err.Error(), http.StatusInternalServerError)
		return
	}
	rw.Header().Set("Content-Type", "text/html")

	fmt.Fprint(rw, "<table>")
	for _, cv := range containers {
		c, err := client.InspectContainer(cv.ID)
		if err != nil {
			fmt.Println("Error inspecting container:", err)
			continue
		}

		for _, e := range c.Config.Env {
			if strings.HasPrefix(e, "VIRTUAL_HOST=") {
				name := strings.TrimPrefix(c.Name, "/")
				vhost := strings.TrimPrefix(e, "VIRTUAL_HOST=")
				status := "Stopped"
				if c.State.Running {
					status = "Running"
				}
				fmt.Fprintf(rw, "<tr><td>%s</td><td><a href='http://%s'>http://%s</a></td><td>%s</td></tr>", name, vhost, vhost, status)
			}
		}
	}
	fmt.Fprint(rw, "</table>")
}
