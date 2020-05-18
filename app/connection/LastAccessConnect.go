package connection

import (
	"context"
	"log"
	"time"

	"google.golang.org/api/iterator"
)

/*
GetLastAccess
  最終ログイン時刻を取得する
*/
func GetLastAccess(ctx context.Context, info LastAccessInfo) (lastAccess time.Time, leadErr error) {

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
		// log.Fatal(leadErr)
		return nil, leadErr
	}
	lastAccess = accessData.LastAccess
	return
}

/*
  SetLastAccess
  データ更新（Documentが存在しない場合は新規作成する）
  ※ただし、不要なデータが存在する場合、全て消去してしまうため、
    状況に応じてMargeAllの設定が必要。
*/
func SetLastAccess(ctx context.Context, info LastAccessInfo) (err error) {

	_, err = getCollection(ctx, COLLECTION_NAME).Doc(info.UserName).Set(ctx, info)
	if err != nil {
		// log.Fatal(err)
		return err
	}

	return
}
