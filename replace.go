package replace

import "strings"

func ForFileName(str string) string {
	str = strings.Replace(str, "。", ".", -1)
	str = strings.Replace(str, "，", ",", -1)
	str = strings.Replace(str, "《", "(", -1)
	str = strings.Replace(str, "》", ")", -1)
	str = strings.Replace(str, "【", "(", -1)
	str = strings.Replace(str, "】", ")", -1)
	//（ ）
	str = strings.Replace(str, "（", "(", -1)
	str = strings.Replace(str, "）", ")", -1)
	str = strings.Replace(str, "「", "(", -1)
	str = strings.Replace(str, "」", ")", -1)
	str = strings.Replace(str, "+", "_", -1)
	str = strings.Replace(str, "`", "", -1)
	str = strings.Replace(str, " ", "", -1)
	str = strings.Replace(str, "\u00A0", "", -1)
	str = strings.Replace(str, "\u0000", "", -1)
	str = strings.Replace(str, "·", "", -1)
	str = strings.Replace(str, "\uE000", "", -1)
	str = strings.Replace(str, "\u000D", "", -1)
	str = strings.Replace(str, "、", "", -1)
	str = strings.Replace(str, "/", "", -1)
	str = strings.Replace(str, "！", "", -1)
	str = strings.Replace(str, "|", "_", -1)
	str = strings.Replace(str, "｜", "_", -1)
	str = strings.Replace(str, ":", "_", -1)
	str = strings.Replace(str, " ", "_", -1)

	return str
}
func Replace(str string) string {
	str = strings.Replace(str, "\n", "", -1)
	str = strings.Replace(str, " ", "_", -1)
	str = strings.Replace(str, "，", "_", -1)
	str = strings.Replace(str, "。", "_", -1)
	str = strings.Replace(str, ",", "_", -1)
	str = strings.Replace(str, "《", "", -1)
	str = strings.Replace(str, "》", "", -1)
	str = strings.Replace(str, "【", "", -1)
	str = strings.Replace(str, "】", "", -1)
	str = strings.Replace(str, "(", "", -1)
	str = strings.Replace(str, "+", "_", -1)
	str = strings.Replace(str, ")", "", -1)
	str = strings.Replace(str, "`", "", -1)
	str = strings.Replace(str, " ", "", -1)
	str = strings.Replace(str, "\u00A0", "", -1)
	str = strings.Replace(str, "\u0000", "", -1)
	str = strings.Replace(str, "·", "", -1)
	str = strings.Replace(str, "\uE000", "", -1)
	str = strings.Replace(str, "\u000D", "", -1)
	str = strings.Replace(str, "、", "", -1)
	str = strings.Replace(str, "！", "", -1)
	str = strings.Replace(str, "|", "_", -1)
	str = strings.Replace(str, "｜", "_", -1)

	return str
}
