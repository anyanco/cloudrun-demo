package util

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetAppFilePath(t *testing.T) {

	gsSample := "gs://sample/index.html"
	fileSample := "file://sample/index.html"
	notSetSample := "sample/index.html"
	otherSample := "http://sample/index.html"

	patterns := []struct {
		exp   FileInfo // 期待値
		param string   // 引数
		title string   // テスト内容
	}{
		// パターンを把握して追加・削除を行うのが楽になる
		{FileInfo{gsSample, GCS}, gsSample, "gs"},
		{FileInfo{fileSample, Local}, fileSample, "file"},
		{FileInfo{fileSample, Local}, notSetSample, "Not Set"},
		{FileInfo{fileSample, Local}, otherSample, "Other schema"},
	}

	assert := assert.New(t)
	for _, p := range patterns {
		assert.Equal(p.exp, GetAppFilePath(p.param), fmt.Sprintf("%s is Error.", p.title))
	}
}

func TestIsSet(t *testing.T) {
	patterns := []struct {
		objInfo GcsObjectInfo // 設定値
		exp     bool          // 期待値
		title   string        // テスト内容
	}{
		{GcsObjectInfo{"sandbox", "index.html"}, true, "Correct Both"},
		{GcsObjectInfo{"", "index.html"}, false, "Empty Bucket"},
		{GcsObjectInfo{"sandbox", ""}, false, "Empty Object"},
		{GcsObjectInfo{"", ""}, false, "Empty Both"},
	}

	assert := assert.New(t)
	for _, p := range patterns {
		assert.Equal(p.exp, p.objInfo.IsSet(), fmt.Sprintf("%s is Error.", p.title))
	}
}

func TestGetGcsObjectInfo(t *testing.T) {

	patterns := []struct {
		exp       GcsObjectInfo // 期待値
		paramPath string        // 引数（パス）
		title     string        // テスト内容
	}{
		{GcsObjectInfo{"sandbox", "index.html"}, "gs://sandbox/index.html", "Correct Non Folder"},
		{GcsObjectInfo{"sandbox", "test/index.html"}, "gs://sandbox/test/index.html", "Correct under Folder"},
		{GcsObjectInfo{"", ""}, "http://sandbox/index.html", "MisMatch Schema"},
	}

	assert := assert.New(t)
	for _, p := range patterns {
		assert.Equal(p.exp, GetGcsObjectInfo(p.paramPath), fmt.Sprintf("%s is Error.", p.title))
	}

}

func TestGetParamValue(t *testing.T) {

	patterns := []struct {
		exp       string // 期待値
		paramPath string // 引数（パス）
		paramKey  string // 引数（クエリKey）
		title     string // テスト内容
	}{
		{"", "http://localhost", "u", "NoParam"},
		{"hoge", "http://localhost?u=hoge", "u", "CorrenctParam"},
		{"", "http://localhost?u=hoge", "", "Nothing Key "},
	}

	assert := assert.New(t)
	for _, p := range patterns {
		assert.Equal(p.exp, GetParamValue(p.paramPath, p.paramKey), fmt.Sprintf("%s Pattern is Error.", p.title))
	}
}
