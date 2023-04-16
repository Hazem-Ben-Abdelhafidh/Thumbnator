package main

import (
	"log"
	"net/http"
)

func setupRoutes() {
	log.Println("listening on port 8080")
	http.HandleFunc("/uploadFile", uploadFile)
	//serve the images
	http.Handle("/temp-images/", http.StripPrefix("/temp-images/", http.FileServer(http.Dir("./temp-images"))))

	http.ListenAndServe(":8080", nil)
}

func main() {

	setupRoutes()
}
