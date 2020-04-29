package main

    import (
        "fmt"
        "log"
        "net/http"
        "os"
    )

    // routing
    func handler(w http.ResponseWriter, r *http.Request) {
        path := r.URL.Path
        switch path {
          case "/storage" :
            getStorageText(w, r)
          case "/firestore" :
            GetFirestoreData(w, r)
          case "/" :
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
