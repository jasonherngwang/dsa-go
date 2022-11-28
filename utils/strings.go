package utils

import (
	"sort"
	"strings"
)

// Split string into chars
func SortChars(str string) string {
	chars := strings.Split(str, "")
	sort.Strings(chars)
	return strings.Join(chars, "")
}

// Lots todo here
