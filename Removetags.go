package goreloaded

func Remove_tags(s []string) string {
	str := ""

	for i, tag := range s {
		if tag == "(cap," || tag == "(low," || tag == "(up," {
			s[i] = ""
			s[i+1] = ""
		} else if tag != "(up)" && tag != "(hex)" && tag != "(bin)" && tag != "(cap)" && tag != "(low)" && tag != "" {
			if i == 0 {
				str = str + tag
			} else {
				str = str + " " + tag
			}
		}
	}
	return str
}
