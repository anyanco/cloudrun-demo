package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"cloud.google.com/go/storage"
)

// get Text from GCS
func getStorageText(w http.ResponseWriter, r *http.Request) {

	log.Print("getStorageText received a request.")
	ctx := r.Context()
	client, err := storage.NewClient(ctx)
	if err != nil {
		log.Fatal(err)
	}

	// connect to GCS
	bucketName := os.Getenv("BUCKET_NAME")
	fmt.Printf("bucketname : %s\n", bucketName)
	objectPath := "demo-test.txt"
	obj := client.Bucket(bucketName).Object(objectPath)
	reader, err := obj.NewReader(ctx)
	if err != nil {
		log.Fatal(err)
	}

	msg := "Error"
	// read File
	txt, err := ioutil.ReadAll(reader)
	if err != nil {
		log.Fatal(err)
	} else {
		msg = string(txt)
	}
	defer reader.Close()

	fmt.Fprintf(w, msg)
}

func GetHtmlText(ctx context.Context, targetHtml string) (html string) {

	client, err := storage.NewClient(ctx)

	if err != nil {
		log.Fatal(err)
		return
	}

	// connect to GCS
	bucketName := os.Getenv("BUCKET_NAME")
	fmt.Printf("bucketname : %s\n", bucketName)
	obj := client.Bucket(bucketName).Object(targetHtml)
	reader, err := obj.NewReader(ctx)
	defer reader.Close()

	if err != nil {
		log.Fatal(err)
		return
	}

	// read File
	htmlBite, err := ioutil.ReadAll(reader)
	if err != nil {
		log.Fatal(err)
		return
	} else {
		html = string(htmlBite)
	}

	return
}
