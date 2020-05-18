package util

import "time"

var CurrentTime = func() time.Time {
	return time.Now()
}

func TimeToSring(baseTime time.Time) (timeString string) {
	const dateFormat = "2006/01/02 15:04:05" // 24h表現、0埋めあり
	jst := time.FixedZone("Asia/Tokyo", 9*60*60)
	timeJST := baseTime.In(jst)
	timeString = timeJST.Format(dateFormat)
	return
}
