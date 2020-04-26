package utils

import "strings"

// QuoteString escapes and quotes a string making it safe for interpolation into an SQL string.
func QuoteString(str string) string {
	return "'" + strings.Replace(str, "'", "''", -1) + "'"
}
