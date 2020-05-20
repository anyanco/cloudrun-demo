package connection

import (
	"context"
	"log"
	"time"

	"github.com/aki36-an/cloudrun-demo/app/common"
	"github.com/aki36-an/cloudrun-demo/app/entity"
	"github.com/aki36-an/cloudrun-demo/app/util"
	"google.golang.org/api/iterator"
)

/*
GetLastAccess
  最終ログイン時刻を取得する
*/
func GetLastAccess(ctx context.Context, info entity.LastAccessInfo) (lastAccess time.Time, leadErr error) {

	collection := getCollection(ctx, common.COLLECTION_NAME)

	query := collection.Where(common.USER_NAME, "==", info.UserName)
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
	var accessData entity.LastAccessInfo
	if leadErr := doc.DataTo(&accessData); leadErr != nil {
		initTime := util.InitTime()
		return initTime, leadErr
	}
	lastAccess = accessData.LastAccess
	log.Printf("lastAccess >> %s", lastAccess)
	return
}

/*
  SetLastAccess
  データ更新（Documentが存在しない場合は新規作成する）
  ※ただし、不要なデータが存在する場合、全て消去してしまうため、
    状況に応じてMargeAllの設定が必要。
*/
func SetLastAccess(ctx context.Context, info entity.LastAccessInfo) (err error) {

	_, err = getCollection(ctx, common.COLLECTION_NAME).Doc(info.UserName).Set(ctx, info)
	if err != nil {
		// log.Fatal(err)
		return err
	}

	return
}
