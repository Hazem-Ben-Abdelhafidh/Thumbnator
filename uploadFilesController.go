package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

type UploadFileResponse struct {
	Thumbnails []string `json:"thumbnails"`
}

func uploadFile(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	if r.Method == "POST" {
		r.ParseMultipartForm(10 << 20)
		file, handler, err := r.FormFile("Video")
		if err != nil {
			log.Println("Error while uploading file: ", err)
			return
		}
		log.Println("Handler: ", handler.Filename)
		defer file.Close()
		fileNameSplitted := strings.Split(handler.Filename, ".")
		var fileName strings.Builder
		fileName.WriteString(fileNameSplitted[0])
		fileName.WriteString("-*.")
		fileName.WriteString(fileNameSplitted[1])
		tempFile, err := ioutil.TempFile("temp-images", fileName.String())
		if err != nil {
			log.Println("error creating this:  ", err)
		}
		defer tempFile.Close()
		// read all of the contents of our uploaded file into a
		// byte array
		fileBytes, err := ioutil.ReadAll(file)
		if err != nil {
			log.Println(err)
		}
		// write this byte array to our temporary file
		tempFile.Write(fileBytes)
		_, err = SplitVideoToframes(tempFile.Name())
		if err != nil {
			log.Println("error splitting Video: ", err)
			return
		}
		files, err := os.ReadDir("./temp-images")
		if err != nil {
			log.Println("Couldn't Read Directory")
			return
		}

		images := []string{}
		for _, file := range files {
			fileNameSplitted := strings.Split(file.Name(), ".")
			if fileNameSplitted[len(fileNameSplitted)-1] == "png" {
				images = append(images, "http://localhost:8080/temp-images/"+file.Name())

			}
		}

		responseBody := UploadFileResponse{
			Thumbnails: images,
		}
		json.NewEncoder(w).Encode(responseBody)

	}

}
