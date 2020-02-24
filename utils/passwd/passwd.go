package passwd

import "regexp"

// 密码脱敏
func Desensitization(str string, x string) string {
	re, _ := regexp.Compile("(\\S{1})(\\S{1,20})(\\S{1})")
	return re.ReplaceAllString(str, "$1"+x+"$3")
}
