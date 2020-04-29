package main

    import (
        "fmt"
        "log"
        "net/http"
        "os"
        "io/ioutil"
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
        bucketName := os.Getenv("BUCKETNAME")
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
