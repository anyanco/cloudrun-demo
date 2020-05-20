package util

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEmptyString(t *testing.T) {
	assert := assert.New(t)
	assert.Equal("", EmptyString(), fmt.Sprintf("%s is Error.", "EmptyTesting"))
}

func TestIsEmpty(t *testing.T) {

	patterns := []struct {
		exp       bool   // 期待値
		paramText string // 引数
		title     string // テスト内容
	}{
		{true, "", "len 0"},
		{false, "0", "0"},
		{true, " ", "Space"},
		{true, "　", "2bite Space"},
		{false, "あいうえお", "Collect 2bite"},
		{false, "a16gdsa", "Collect alphaNumeric"},
	}

	assert := assert.New(t)
	for _, p := range patterns {
		assert.Equal(p.exp, IsEmpty(p.paramText), fmt.Sprintf("%s Pattern is Error.", p.title))
	}
}
