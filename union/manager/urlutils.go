package manager

import "strings"

func SubstringAfterLast(str, delimiter string) string {
	newStr := strings.Split(str, delimiter)
	return newStr[len(newStr) -1]
} 

func GetExtension(str string) string {
	return SubstringAfterLast(str, ".")
}