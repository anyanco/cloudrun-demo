package util

import (
	"log"
	"net/url"
	"strings"
)

type PathSchemaType int

const (
	UnknownSchema = iota
	GCS
	Local
)

type FileInfo struct {
	FilePath   string
	SchemaType PathSchemaType
}

/*
  GetAppFilePath
  ファイルパスを取得する。
  @return ファイルパス
    「gs://〜」「file://~」で始まる場合はそのまま返却。
    そうでなければ「file://~」にして返却
*/
func GetAppFilePath(targetPath string) FileInfo {

	log.Printf("targetPath : %s\n", targetPath)

	u, err := url.Parse(targetPath)
	if err != nil {
		// パースでエラーが発生した場合は空文字を返却
		return FileInfo{EmptyString(), UnknownSchema}
	}

	// schemaの有無を確認
	if u.IsAbs() {
		log.Printf("scheme : %s\n", u.Scheme)

		switch u.Scheme {
		case "gs":
			return FileInfo{targetPath, GCS}

		case "file":
			return FileInfo{targetPath, Local}
		}
	}

	// スキーマが設定されていない、もしくは別のschemaが設定されていれば[file://~]として扱う
	u.Scheme = "file"

	return FileInfo{u.String(), Local}
}

type GcsObjectInfo struct {
	BucketName string
	ObjectName string
}

func (i *GcsObjectInfo) IsSet() bool {
	if len(i.BucketName) > 0 && len(i.ObjectName) > 0 {
		return true
	}
	return false
}

/*
  GetGcsObjectInfo
  ファイルパスからGCSのファイル情報を取得する
*/
func GetGcsObjectInfo(targetPath string) GcsObjectInfo {
	u, err := url.Parse(targetPath)
	if err != nil || u.Scheme != "gs" {
		return GcsObjectInfo{}
	}

	obj := strings.TrimLeft(u.Path, "/")
	return GcsObjectInfo{u.Host, obj}
}

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
