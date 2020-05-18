package util

import (
	"log"
	"net/url"
)

/*
GetParamValue
  URLパラメータから指定したKeyの値を取得する
*/
func GetParamValue(reqUrl string, key string) (userName string) {

	log.Printf("start GetParamValue : %s", key)
	if len(key) == 0 {
		// 対象のkeyが設定されていない場合はnilを返却
		return nil
	}

	urlString, err := url.Parse(reqUrl)
	if err != nil {
		// URLがパースできなかった場合はnilを返却
		return nil
	}

	for key, value := range urlString.Query() {
		if key == key {
			log.Printf("UserName : %s", value)
			userName = value[0]
			return
		}
	}
	return nil
}
