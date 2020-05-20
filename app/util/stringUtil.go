package util

/*
  EmptyString:string初期値を返却する
*/
var EmptyString = func() string {
	return ""
}

func IsEmpty(text string) bool {
	if len(text) == 0 {
		return true
	}
	if text == " " {
		return true
	}
	if text == "　" {
		return true
	}
	return false
}
