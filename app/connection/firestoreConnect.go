package connection

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"cloud.google.com/go/firestore"
	"github.com/aki36-an/cloudrun-demo/app/util"
)

// GetFirestoreData
func GetFirestoreData(w http.ResponseWriter, r *http.Request) {

	log.Print("getFirestoreData received a request.")
	// sample
	userName := "test"
	ctx := r.Context()
	info := InitLastAccessInfo(userName, time.Now())
	lastAccessDate := util.TimeToSring(GetLastAccess(ctx, info))
	fmt.Fprintf(w, "before: %s\n", lastAccessDate)
	fmt.Fprintf(w, "latest: %s", TimeToSring(time.Now()))
}

func getCollection(ctx context.Context, collectionName string) (collection *firestore.CollectionRef) {
	fmt.Printf("getCollection/collectionName : %s\n", collectionName)
	client, err := firestore.NewClient(ctx, os.Getenv("PROJECT_ID"))
	if err != nil {
		log.Fatal(err, ERROR_NO_FS_CLIENT)
		return
	}
	collection = client.Collection(collectionName)
	return
}
