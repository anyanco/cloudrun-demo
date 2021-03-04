package connection

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"cloud.google.com/go/storage"
	"github.com/aki36-an/cloudrun-demo/app/common"
	"github.com/aki36-an/cloudrun-demo/app/util"
)

// get Text from GCS
func GetStorageText(w http.ResponseWriter, r *http.Request) {

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

func GetHtmlText(ctx context.Context) (html string) {

	html = util.EmptyString()
	fileInfo := common.TemplatePathInfo()

	switch fileInfo.SchemaType {
	case util.GCS:
		gcs := util.GetGcsObjectInfo(fileInfo.FilePath)
		html = GetGcsObjectString(ctx, gcs)
		return
	case util.Local:
		// Localのファイル読み込みの挙動は未実装
		return
	}

	return
}

/*
	GetGcsObjectString
	Google Cloud Strage からファイルを取得し、テキスト形式で返却する
*/
func GetGcsObjectString(ctx context.Context, gcsInfo util.GcsObjectInfo) (text string) {

	text = util.EmptyString()

	if !gcsInfo.IsSet() {
		return
	}

	client, err := storage.NewClient(ctx)
	if err != nil {
		return
	}

	obj := client.Bucket(gcsInfo.BucketName).Object(gcsInfo.ObjectName)
	reader, err := obj.NewReader(ctx)
	defer reader.Close()
	if err != nil {
		return
	}

	htmlBite, err := ioutil.ReadAll(reader)
	if err != nil {
		return
	}

	text = string(htmlBite)
	return
}

func GetFixWords(w http.ResponseWriter, r *http.Request) {

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
