package util

import (
	"log"
	"net/url"
)

/*
	GetParamValue
  URLパラメータから指定したKeyの値を取得する
	@return 取得できなかった場合は空文字を返却
*/
func GetParamValue(reqUrl string, paramKey string) (paramValue string) {

	// パラメータからユーザ名を取得できなかった場合の初期値
	paramValue = ""
	log.Printf("start GetParamValue Key : %s", paramKey)
	if len(paramKey) == 0 {
		return
	}

	urlString, err := url.Parse(reqUrl)
	if err != nil {
		return
	}

	for key, value := range urlString.Query() {
		if key == paramKey {
			log.Printf("paramValue : %s", value)
			paramValue = value[0]
			return
		}
	}
	return
}
