package db

import "strings"

// bottomLineToUpper 讓 _ 之後的字母大寫
func bottomLineToUpper(str string) string {
	names := strings.Split(str, "_")

	var name string
	for _, val := range names {
		name += strings.Title(val)
	}

	return name
}
