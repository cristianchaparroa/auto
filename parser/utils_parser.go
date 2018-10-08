package parser

import "regexp"

// FilterSpaces removes spaces in a line
func FilterSpaces(line string) string {
	leadCloseRe := regexp.MustCompile(`^[\s\p{Zs}]+|[\s\p{Zs}]+$`)
	insideRe := regexp.MustCompile(`[\s\p{Zs}]{2,}`)
	filtered := leadCloseRe.ReplaceAllString(line, "")
	filtered = insideRe.ReplaceAllString(filtered, " ")
	return filtered
}
