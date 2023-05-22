package goreloaded

import "strings"

//  This capitalises after every whitespace or hyphen. Same as toLower Case.
func To_upper(str string) string {
	return strings.ToUpper(str)
}

//  The strings package contains a Compare method.
func Compare(a, b string) int {
	return strings.Compare(a, b)
}

//  This capitalises after every whitespace or hyphen. Same as toUpperCase.
func To_lower(str string) string {
	return strings.ToLower(str)
}

// gets the first rune of a string
func First_rune(s string) string {
	a := []rune(s)
	return string(a[0])
}

func Capitalise(str string) string {
	return strings.Title(str)
}

// seperate string by spaces and appends to string list
func Split_by_spaces(s string) []string {
	return strings.Fields(s)
}
