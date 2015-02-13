package main

import (
	"encoding/json"
	"net/http"

	"github.com/fsouza/go-dockerclient"
	"github.com/gorilla/mux"
)

type Image struct {
	Tags        []string
	ID          string
	ParentID    string
	Size        int64
	VirtualSize int64
}

func ImageList(w http.ResponseWriter, r *http.Request) {
	imgs, err := client.ListImages(docker.ListImagesOptions{})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	images := make([]Image, 0)
	for _, i := range imgs {
		img := Image{
			Tags:        i.RepoTags,
			ID:          i.ID,
			ParentID:    i.ParentID,
			Size:        i.Size,
			VirtualSize: i.VirtualSize,
		}

		images = append(images, img)
	}

	json.NewEncoder(w).Encode(images)
}

func ImageInfo(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	image, err := client.InspectImage(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(image)
}
