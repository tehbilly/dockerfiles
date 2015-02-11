package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

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
	mux.Handle("/", http.FileServer(http.Dir("./site")))
	mux.HandleFunc("/containers", ListContainers)
	mux.HandleFunc("/container/", ContainerInfo)
	mux.HandleFunc("/images", ListImages)
	mux.HandleFunc("/image/", ImageInfo)
	mux.HandleFunc("/dockerinfo", DockerServerInfo)

	fmt.Println("Starting http server...")
	err := http.ListenAndServe(":3000", mux)

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
