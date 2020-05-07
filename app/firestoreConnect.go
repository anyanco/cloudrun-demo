package main

import (
    "fmt"
    "log"
    "net/http"
    "os"
    "time"
    "context"
    "net/url"
    "cloud.google.com/go/firestore"
    "google.golang.org/api/iterator"
)


// get Firestore data
func GetFirestoreData(w http.ResponseWriter, r *http.Request) {

    log.Print("getFirestoreData received a request.")
    // sample
    userName := "test"
    ctx := r.Context()
    info := InitLastAccessInfo(userName, time.Now())
    lastAccessDate := TimeToSring(GetLastAccess(ctx, info))
    fmt.Fprintf(w, "before: %s\n", lastAccessDate)
    fmt.Fprintf(w, "latest: %s", TimeToSring(time.Now()))
}

/*
  URLパラメータからユーザ名を取得する
*/
func GetUserName(reqUrl string)(userName string){
  fmt.Printf("start GetUserName URL : %s\n", reqUrl)
  // log.Print("start GetUserName : " + reqUrl)
  urlString, err := url.Parse(reqUrl)
  if err != nil {
    log.Fatal(err, ERROR_USER_NONE)
    return
  }
  for key, value := range urlString.Query() {
    if key == USER_PARAM_KEY {
      fmt.Printf("UserName : %s\n", value)
      // log.Printf("UserName : %s", value)
      userName = value[0]
      return
    }
  }
  return
}

func getCollection(ctx context.Context, collectionName string)(collection *firestore.CollectionRef){
  fmt.Printf("getCollection/collectionName : %s\n", collectionName)
  client, err := firestore.NewClient(ctx, os.Getenv("PROJECT_ID"))
  if err != nil {
    log.Fatal(err, ERROR_NO_FS_CLIENT)
    return
  }
  collection = client.Collection(collectionName)
  return
}

/*
  最終ログイン時刻を取得する
*/
func GetLastAccess(ctx context.Context, info LastAccessInfo) (lastAccess time.Time){

    collection := getCollection(ctx, COLLECTION_NAME)

    query := collection.Where(USER_NAME, "==", info.UserName)
    log.Printf("Query : %v", query)

    iter := query.Documents(ctx)

    // １件だけの取得なので、forは回さない
    doc, err := iter.Next()
    // var Done = errors.New("no more items in iterator")
    if err == iterator.Done {
      log.Print("新規作成 -----")
      return
    }

    // 読み込み（LastAccessInfoに設定）
    var accessData LastAccessInfo
    if leadErr := doc.DataTo(&accessData); leadErr != nil {
      log.Fatal(leadErr)
      return
    }
    lastAccess = accessData.LastAccess
    return
}

/*
  データ更新（Documentが存在しない場合は新規作成する）
  ※ただし、不要なデータが存在する場合、全て消去してしまうため、
    状況に応じてMargeAllの設定が必要。
*/
func SetLastAccess(ctx context.Context, info LastAccessInfo) (err error){

      _, err = getCollection(ctx, COLLECTION_NAME).Doc(info.UserName).Set(ctx, info)
      if err != nil {
        log.Fatal(err)
      }

    return
}


func TimeToSring(baseTime time.Time) (timeString string){
  const dateFormat = "2006/01/02 15:04:05" // 24h表現、0埋めあり
  jst := time.FixedZone("Asia/Tokyo", 9*60*60)
  timeJST := baseTime.In(jst)
  timeString = timeJST.Format(dateFormat)
  return
}
