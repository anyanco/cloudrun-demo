package main

    import (
        "log"
        "net/http"
        "html/template"
    )

    func Login(w http.ResponseWriter, r *http.Request) {

        log.Print("Login received a request.")
        ctx := r.Context()

        lastAccessString, _ := GetAccessString(ctx)
        html := GetHtmlText(ctx, "index.html")

        //埋め込み変数
        params := map[string]string{
            "beforeAccessDate": lastAccessString,
        }

        tpl, _ := template.New("index").Parse(html)
        tpl.Execute(w, params)

    }
