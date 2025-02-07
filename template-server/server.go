package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", handler)
	serverAddress := "localhost:5000"
	fmt.Println("Server Listening at:", serverAddress)
	fmt.Printf("Visit http://%s in your browser\n", serverAddress)
	log.Fatal(http.ListenAndServe(serverAddress, nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	htmlPage, err := os.ReadFile("./templates/index.html")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Fprintf(w, "%s", string(htmlPage))
}
