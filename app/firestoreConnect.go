package main

import (
    "fmt"
    "log"
    "net/http"
    "os"
    "cloud.google.com/go/firestore"
)
// get Firestore data
func GetFirestoreData(w http.ResponseWriter, r *http.Request) {

    log.Print("getFirestoreData received a request.")
    ctx := r.Context()
    client, err := firestore.NewClient(ctx, os.Getenv("PROJECT_ID"))
    if err != nil {
       log.Fatal(err)
    }

    access := client.Doc("demo/demo-access")
    docsnap, err := access.Get(ctx)
    if err != nil {
      log.Fatal(err)
    }
    dataMap := docsnap.Data()
    fmt.Println(dataMap)
    fmt.Fprintf(w, "%s", dataMap)
}
