package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/fsouza/go-dockerclient"
	"github.com/gorilla/mux"
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
	r := mux.NewRouter()
	r.NotFoundHandler = http.FileServer(http.Dir("site"))
	r.HandleFunc("/containers/list", ContainerList)
	r.HandleFunc("/containers/{id}/info", ContainerInfo)
	r.HandleFunc("/containers/{id}/start", ContainerStart)
	r.HandleFunc("/containers/{id}/stop", ContainerStop)
	r.HandleFunc("/containers/{id}/kill", ContainerKill)
	r.HandleFunc("/containers/{id}/restart", ContainerRestart)
	r.HandleFunc("/images/list", ImageList)
	r.HandleFunc("/images/{id}/info", ImageInfo)
	r.HandleFunc("/dockerinfo", DockerServerInfo)

	fmt.Println("Starting http server...")
	err := http.ListenAndServe(":3000", r)

	if err != nil {
		fmt.Println("Unable to start http server:", err)
		os.Exit(1)
	}
}

func DockerServerInfo(w http.ResponseWriter, r *http.Request) {
	info, err := client.Info()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(info)
}
