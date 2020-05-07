package main

    import (
        "log"
        "fmt"
        "time"
        "net/http"
        "html/template"
    )

    func Login(w http.ResponseWriter, r *http.Request) {

        log.Print("Login received a request.")
        ctx := r.Context()

        // URLパラメータからユーザ名を取得
        userName := GetUserName(r.URL.String())
        log.Printf("GetUserName/UserName >> %s\n", userName)
        if len(userName) == 0 {
          fmt.Fprint(w, ERROR_USER_NONE)
          return
        }

        nowDate := time.Now()
        accessInfo := InitLastAccessInfo(userName, nowDate)

        // 最終アクセス日時取得
        lastAccess := GetLastAccess(ctx, accessInfo)
        if lastAccess.IsZero() {
          // データが存在しない場合は今回のアクセス日時を最新とする
          lastAccess = nowDate
        }

        // 最終アクセス日時を設定
        execErr := SetLastAccess(ctx, accessInfo)
        if execErr != nil {
          fmt.Fprint(w, ERROR_MISS_WRITE)
          return
        }

        // テンプレート取得
        html := GetHtmlText(ctx, "index.html")

        //埋め込み変数
        params := map[string]string{
            "beforeAccessDate": TimeToSring(lastAccess),
        }

        tpl, _ := template.New("index").Parse(html)
        tpl.Execute(w, params)

    }
