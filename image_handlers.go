package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

//Image struct to define entries for images table in the DB
type Image struct {
	Repository string `json:"repository"`
	Id         string `json:"id"`
	Tag        string `json:"tag"`
	Size       string `json:"size"`
	Created    string `json:"created"`
	Requestor  string `json:"requestor"`
	Status     string `json:"status"`
}

//ImageList displays images in the DB
func ImageList(w http.ResponseWriter, r *http.Request) {

	images, err := store.GetImageDownloads()

	imageListBytes, err := json.Marshal(images)

	if err != nil {
		fmt.Println(fmt.Errorf("Error: %v", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Write(imageListBytes)
}

//ImageDownload initiates a download into the DB
func ImageDownload(w http.ResponseWriter, r *http.Request) {
	img := Image{}

	err := r.ParseForm()

	if err != nil {
		fmt.Println(fmt.Errorf("Error: %v", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	img.Repository = r.Form.Get("Repository")
	img.Id = r.Form.Get("ID")
	img.Tag = r.Form.Get("Tag")
	img.Size = r.Form.Get("Size")
	img.Created = r.Form.Get("Created")
	img.Requestor = r.Form.Get("Requestor")
	img.Status = r.Form.Get("Status")

	// The only change we made here is to use the `CreateBird` method instead of
	// appending to the `bird` variable like we did earlier
	err = store.CreateDownload(&img)
	if err != nil {
		fmt.Println(err)
	}

	http.Redirect(w, r, "/assets/", http.StatusFound)

	fmt.Fprintln(w, "Download images from here !")
}
