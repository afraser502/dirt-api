package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"github.com/spf13/viper"
)

/*type ImageReturn struct {
	test string
}*/

var repo = "alpine"
var tag = "latest"
var imgDigest = ""

//Handler for downloads page - calls ImagePull function
func imgPull(w http.ResponseWriter, r *http.Request) {

	img := Image{}

	err := r.ParseForm()

	if err != nil {
		fmt.Println(fmt.Errorf("Error: %v", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	img.Repository = r.Form.Get("Repository")
	//img.Id =
	img.Tag = r.Form.Get("Tag")
	//img.Size =
	//img.Created =
	//img.Requestor =
	//img.Status =

	test, err := ImagePull(repo, tag)

	fmt.Println(test)

	http.Redirect(w, r, "/assets/", http.StatusFound)

	fmt.Fprintln(w, "Download images from here !")
}

//ImagePull retrieves an external image
func ImagePull(repo string, tag string) (string, error) {
	//Set main variables
	viper.SetConfigFile("./configs/env.json") //Set config file

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}

	ctx := context.Background()
	cli, err := client.NewEnvClient()

	if err != nil {
		panic(err)
	}

	//Struct for download response
	type ImagePullResponse struct {
		ID             string `json"id"`
		Status         string `json:"status"`
		ProgressDetail struct {
			Current int64 `json:"current"`
			Total   int64 `json:"total"`
		} `json:"progressDetail"`
		Progress string `json:"progress"`
	}

	imageName := repo
	imageTag := tag

	//imgDigest := ""
	imgMatch := false
	imgLocalDigest := ""
	//splitImageName := strings.Split(imageName, "/")

	//intSlice := splitImageName
	//last := intSlice[len(splitImageName)-1]

	imageFullname := ""
	if imageTag != "" {
		imageFullname = imageName + `:` + imageTag
	} else {
		imageTag = "latest"
		imageFullname = imageName + `:` + imageTag

	}

	reader, err := cli.ImagePull(ctx, imageFullname, types.ImagePullOptions{})
	if err != nil {
		panic(err)
	}
	defer reader.Close()

	//termFd, isTerm := term.GetFdInfo(os.Stderr)
	//jsonmessage.DisplayJSONMessagesStream(reader, os.Stderr, termFd, isTerm, nil)

	d := json.NewDecoder(reader)
	for {
		var pullResult ImagePullResponse
		if err := d.Decode(&pullResult); err != nil {
			// handle the error
			break
		}

		//Print all output
		//fmt.Println(pullResult)

		if strings.Contains(pullResult.Status, "Digest") {
			result := strings.Split(pullResult.Status, " ")
			/*
				for i := range result {
					println(result[i])
				}
			*/

			//Will return length of array
			//fmt.Println(len(result))
			//fmt.Println(strings.Split(pullResult.Status, " "))

			//Print last entry for the SHA digest
			//fmt.Println(result[1])
			imgDigest = result[1]
		}

	}

	//Poll local docker images
	images, err := cli.ImageList(context.Background(), types.ImageListOptions{})
	if err != nil {
		panic(err)
	}

	//Check for local digest using repodigests
	for _, image := range images {
		//Second range iteration to split multi slice into single lines
		for l := range image.RepoDigests {
			//Check whether pulled digest is similar to local digest.
			if strings.Contains(image.RepoDigests[l], imgDigest) {
				//fmt.Println(image.RepoDigests[l])
				imgLocalDigest = image.RepoDigests[l]
				imgMatch = true
			}
		}

	}
	//START

	fmt.Println(imgMatch)

	fmt.Println("Image Name: ", imageName)
	fmt.Println("Image Tag : ", imageTag)
	fmt.Println("Digest of pulled image: ", imgDigest)
	fmt.Println("Digest of local image matched: ", imgLocalDigest)

	/*if imgMatch {

	}*/

	return imgDigest, err
}

//Function to retag image
/*func reTagImage(imageFullname string, imageTag string, newRepo string, splitImageName string) {
	ctx := context.Background()
	cli, err := client.NewEnvClient()
	if err != nil {
		panic(err)
	}

	newImageFullName := newRepo + splitImageName + `:` + imageTag
	fmt.Println(newImageFullName)
	fmt.Println(imageFullname)
	err1 := cli.ImageTag(ctx, imageFullname, newImageFullName)
	if err1 != nil {
		panic(err1)
	}

	ImagePush(newImageFullName)
}

//ImagePush of retagged image to new registry
func ImagePush(newImageFullName string) {
	ctx := context.Background()
	cli, err := client.NewEnvClient()

	authConfig := types.AuthConfig{
		Username: "username",
		Password: "password",
	}
	encodedJSON, err := json.Marshal(authConfig)
	if err != nil {
		panic(err)
	}
	authStr := base64.URLEncoding.EncodeToString(encodedJSON)

	out, err := cli.ImagePush(ctx, newImageFullName, types.ImagePushOptions{RegistryAuth: authStr})
	if err != nil {
		panic(err)
	}

	//Parse the responses of image push
	responses, err := ioutil.ReadAll(out)
	fmt.Println(string(responses))

	defer out.Close()

}
*/
