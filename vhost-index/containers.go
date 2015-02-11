package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/fsouza/go-dockerclient"
)

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
