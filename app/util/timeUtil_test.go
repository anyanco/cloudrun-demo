package util

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestInitTime(t *testing.T) {
	assert := assert.New(t)
	assert.Equal(time.Date(1, time.January, 1, 0, 0, 0, 0, time.UTC), InitTime(), fmt.Sprintf("%s Pattern is Error.", "InitTime"))
}

func TestTimeToSring(t *testing.T) {
	jst := time.FixedZone("Asia/Tokyo", 9*60*60)

	patterns := []struct {
		exp   string    // 期待値
		param time.Time // 引数
		title string    // テスト内容
	}{
		{"2020/01/02 23:00:00", time.Date(2020, time.January, 2, 23, 0, 0, 0, jst), "Correct"},
		{"0001/01/01 09:00:00", InitTime(), "time.Zero"},
	}

	assert := assert.New(t)
	for _, p := range patterns {
		assert.Equal(p.exp, TimeToSring(p.param), fmt.Sprintf("%s Pattern is Error.", p.title))
	}
}
