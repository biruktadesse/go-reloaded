package goreloaded

func Remove_spaces_withQuote(s string) string {
	str := ""
	var removeSpace bool
	for i, char := range s {
		if char == 39 && s[i-1] == ' ' {
			if removeSpace {
				str = str[:len(str)-1]
				str = str + string(char)
				removeSpace = false
			} else {
				str = str + string(char)
				removeSpace = true
			}
		} else if i > 1 && s[i-2] == 39 && s[i-1] == ' ' {
			if removeSpace {
				str = str[:len(str)-1]
				str = str + string(char)
			} else {
				str = str + string(char)
			}
		} else {
			str = str + string(char)
		}
	}
	return str
}

func Remove_spaces(s string) string {
	len := len(s) - 1
	if s[len-1] == ' ' {
		return Remove_spaces(s[:len])
	}
	return s[:len]
}
