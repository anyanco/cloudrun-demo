package main

import (
    "fmt"
    "log"
    "net/http"
    "os"
    "time"
    "context"
    "cloud.google.com/go/firestore"
)

// get Firestore data
func GetFirestoreData(w http.ResponseWriter, r *http.Request) {

    log.Print("getFirestoreData received a request.")
    ctx := r.Context()

    lastAccessDate, currentAccessDate := GetAccessString(ctx)

    fmt.Fprintf(w, "before: %s\n", lastAccessDate)
    fmt.Fprintf(w, "latest: %s", currentAccessDate)
}

func GetAccessString(ctx context.Context) (lastAccessString string, currentAccessString string) {

    lastAccessString = "Error"

    client, err := firestore.NewClient(ctx, os.Getenv("PROJECT_ID"))
    if err != nil {
      lastAccessString = "Can't get Client ..."
      return
    }
    access := client.Doc("demo/demo-access")

    // 前回アクセス日時を取得
    lastAccess, isFirstAccess := GetLastAccess(ctx, access)

    // 最終アクセス日時の更新
    nowDate := UpdateLastAccess(ctx, access)
    currentAccessString = TimeToSring(nowDate)

    // 初回ログインの場合は現在日付を前回アクセス日時として使用
    if isFirstAccess {
      lastAccess = nowDate
    }

    lastAccessString = TimeToSring(lastAccess)

    defer client.Close()
    return
}

func GetLastAccess(ctx context.Context, doc *firestore.DocumentRef) (lastAccess time.Time, isFirstAccess bool){

    // 初回ログインかどうか
    isFirstAccess = false

    docsnap, err := doc.Get(ctx)
    if err != nil {
      log.Fatal(err)
      return
    }

    dataMap := docsnap.Data()

    lastAccessDate, isThere := dataMap["last-access"]
    if isThere {
      // time型に変換できるか確認
      lastAccessDateTime, isTime := lastAccessDate.(time.Time)
      if isTime {
        lastAccess = lastAccessDateTime
        isFirstAccess = false
        return
      }
    }

    return
}


func UpdateLastAccess(ctx context.Context, doc *firestore.DocumentRef) (now time.Time){
    now = time.Now()
    _, err := doc.Update(ctx, []firestore.Update{
      {Path: "last-access", Value: now},
    })
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
