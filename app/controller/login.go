package controller

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/aki36-an/cloudrun-demo/app/common"
	"github.com/aki36-an/cloudrun-demo/app/connection"
	"github.com/aki36-an/cloudrun-demo/app/entity"
	"github.com/aki36-an/cloudrun-demo/app/util"
)

func Login(w http.ResponseWriter, r *http.Request) {

	log.Print("Login received a request.")
	ctx := r.Context()

	// URLパラメータからユーザ名を取得
	userName := util.GetParamValue(r.URL.String(), common.USER_PARAM_KEY)
	log.Printf("GetUserName/UserName >> %s\n", userName)
	if len(userName) == 0 {
		fmt.Fprint(w, common.ERROR_USER_NONE)
		return
	}

	nowDate := util.CurrentTime()
	accessInfo := entity.InitLastAccessInfo(userName, nowDate)

	// 最終アクセス日時取得
	lastAccess, err := connection.GetLastAccess(ctx, accessInfo)
	if lastAccess.IsZero() || err != nil {
		// データが存在しない場合は今回のアクセス日時を最新とする
		lastAccess = nowDate
	}

	// 最終アクセス日時を設定
	execErr := connection.SetLastAccess(ctx, accessInfo)
	if execErr != nil {
		fmt.Fprint(w, common.ERROR_MISS_WRITE)
		return
	}

	// テンプレート取得
	html := connection.GetHtmlText(ctx, "index.html")

	//埋め込み変数
	params := map[string]string{
		"beforeAccessDate": util.TimeToSring(lastAccess),
	}

	tpl, _ := template.New("index").Parse(html)
	tpl.Execute(w, params)

}
