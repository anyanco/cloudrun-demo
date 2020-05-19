package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/aki36-an/cloudrun-demo/app/connection"
	"github.com/aki36-an/cloudrun-demo/app/controller"
)

// routing
func handler(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	switch path {
	case "/storage":
		connection.GetStorageText(w, r)
	case "/firestore":
		connection.GetFirestoreData(w, r)
	case "/login":
		controller.Login(w, r)
	case "/":
		fmt.Fprintf(w, "Hello Docker World")
	default:
	}
}

func main() {
	http.HandleFunc("/", handler)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}
