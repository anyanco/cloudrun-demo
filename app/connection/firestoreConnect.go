package connection

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"cloud.google.com/go/firestore"
	"github.com/aki36-an/cloudrun-demo/app/common"
	"github.com/aki36-an/cloudrun-demo/app/entity"
	"github.com/aki36-an/cloudrun-demo/app/util"
)

// GetFirestoreData
func GetFirestoreData(w http.ResponseWriter, r *http.Request) {

	log.Print("getFirestoreData received a request.")
	// sample
	userName := "test"
	ctx := r.Context()
	info := entity.InitLastAccessInfo(userName, time.Now())
	lastAccess, err := GetLastAccess(ctx, info)
	if err != nil {
		fmt.Fprintf(w, "before: %s\n", "Nothing")
		return
	}
	lastAccessDate := util.TimeToSring(lastAccess)
	fmt.Fprintf(w, "before: %s\n", lastAccessDate)
	fmt.Fprintf(w, "latest: %s", util.TimeToSring(time.Now()))
}

func getCollection(ctx context.Context, collectionName string) (collection *firestore.CollectionRef) {
	fmt.Printf("getCollection/collectionName : %s\n", collectionName)
	client, err := firestore.NewClient(ctx, os.Getenv("PROJECT_ID"))
	if err != nil {
		log.Fatal(err, common.ERROR_NO_FS_CLIENT)
		return
	}
	collection = client.Collection(collectionName)
	return
}
