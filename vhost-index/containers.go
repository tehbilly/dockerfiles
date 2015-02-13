package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/fsouza/go-dockerclient"
	"github.com/gorilla/mux"
)

type Container struct {
	ID      string
	Image   string
	ImageID string
	Name    string
	Status  string
	VHost   string
	Running bool
}

func ContainerList(w http.ResponseWriter, r *http.Request) {
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
		// This will be replaced when we're handling routing ourselves
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
			Running: c.State.Running,
		}

		ci = append(ci, co)
	}

	json.NewEncoder(w).Encode(ci)
}

func ContainerInfo(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	c, err := client.InspectContainer(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(c)
}

func ContainerStart(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	err := client.StartContainer(id, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Fprintf(w, "Started container %s", id)
	log.Println("Started container:", id)
}

func ContainerStop(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	err := client.StopContainer(id, 10)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Fprintf(w, "Stopped container %s", id)
	log.Println("Stopped container:", id)
}

func ContainerKill(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "Not implemented", http.StatusNotImplemented)
}

func ContainerRestart(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "Not implemented", http.StatusNotImplemented)
}
