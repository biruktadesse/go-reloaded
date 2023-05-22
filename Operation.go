package goreloaded

import (
	"fmt"
	"os"
	"strconv"
)

// reads file and saves content to 'data' var
func Operation() {
	data, err := os.ReadFile(os.Args[1])
	Check(err)
	input := string(data)
	result := Split_by_spaces(input)

	// runs range loop to modify result
	for i, v := range result {
		// replaces the word before with its decimal version
		if Compare(v, "(hex)") == 0 {
			j, _ := strconv.ParseInt(result[i-1], 16, 64)
			result[i-1] = fmt.Sprint(j)
		}
		// replaces the word before with its decimal version
		if Compare(v, "(bin)") == 0 {
			j, _ := strconv.ParseInt(result[i-1], 2, 64)
			result[i-1] = fmt.Sprint(j)
		}
		// converts the word before to lowercase
		if Compare(v, "(low)") == 0 {
			result[i-1] = To_lower(result[i-1])
		}
		// converts the number of words before to lowercase
		if Compare(v, "(low,") == 0 {
			result[i-1] = To_lower(result[i-1])

			le := len(result[i+1])
			numb := result[i+1][:le-1]
			nu, err := strconv.Atoi(numb)
			Check(err)

			for j := 1; j <= nu; j++ {
				result[i-j] = To_lower(result[i-j])
			}
		}
		// converts the word before to uppercase
		if Compare(v, "(up)") == 0 {
			result[i-1] = To_upper(result[i-1])
		}
		// converts the number of words before to uppercase
		if Compare(v, "(up,") == 0 {
			result[i-1] = To_upper(result[i-1])

			le := len(result[i+1])
			numb := result[i+1][:le-1]
			nu, err := strconv.Atoi(numb)
			Check(err)

			for j := 1; j <= nu; j++ {
				result[i-j] = To_upper(result[i-j])
			}
		}
		// Capitalises the word before
		if Compare(v, "(cap)") == 0 {
			result[i-1] = Capitalise(result[i-1])
		}
		// Capitalises the number of words before
		if Compare(v, "(cap,") == 0 {
			result[i-1] = Capitalise(result[i-1])

			le := len(result[i+1])
			numb := result[i+1][:le-1]
			nu, err := strconv.Atoi(numb)
			Check(err)

			for j := 1; j <= nu; j++ {
				result[i-j] = Capitalise(result[i-j])
			}
		}
		// converts 'a' into 'an' when the next word begins with a vowel or 'h'.
		if (Compare(v, "a") == 0 && First_rune(result[i+1]) == "a") || (Compare(v, "A") == 0 && First_rune(result[i+1]) == "a") {
			result[i] += "n"
		}
		if (Compare(v, "a") == 0 && First_rune(result[i+1]) == "e") || (Compare(v, "A") == 0 && First_rune(result[i+1]) == "e") {
			result[i] += "n"
		}
		if (Compare(v, "a") == 0 && First_rune(result[i+1]) == "i") || (Compare(v, "A") == 0 && First_rune(result[i+1]) == "i") {
			result[i] += "n"
		}
		if (Compare(v, "a") == 0 && First_rune(result[i+1]) == "o") || (Compare(v, "A") == 0 && First_rune(result[i+1]) == "o") {
			result[i] += "n"
		}
		if (Compare(v, "a") == 0 && First_rune(result[i+1]) == "u") || (Compare(v, "A") == 0 && First_rune(result[i+1]) == "u") {
			result[i] += "n"
		}
		if (Compare(v, "a") == 0 && First_rune(result[i+1]) == "h") || (Compare(v, "A") == 0 && First_rune(result[i+1]) == "h") {
			result[i] += "n"
		}
	}

	// calls Remove_tags() and Split_by_spaces() and gets a new result variable
	Tag_Removed_Result := Remove_tags(result)
	New_Result := Split_by_spaces(Tag_Removed_Result)

	str := ""
	for _, word := range New_Result {
		str = str + word + " "
	}
	// remove spaces from string
	str = Remove_spaces(str)

	word := ""
	for i, char := range str {
		if i == len(str)-1 {
			if char == '.' || char == ',' || char == '!' || char == '?' || char == ';' || char == ':' {
				if str[i-1] == ' ' {
					word = word[:len(word)-1]
					word = word + string(char)
				} else {
					word = word + string(char)
				}
			} else {
				word = word + string(char)
			}
		} else if char == '.' || char == ',' || char == '!' || char == '?' || char == ';' || char == ':' {
			if str[i-1] == ' ' {
				word = word[:len(word)-1]
				word = word + string(char)
			} else {
				word = word + string(char)
			}
			if str[i+1] != ' ' && str[i+1] != '.' && str[i+1] != ',' && str[i+1] != '!' && str[i+1] != '?' && str[i+1] != ';' && str[i+1] != ':' {
				word = word + " "
			}
		} else {
			word = word + string(char)
		}
	}
	// calls Remove_spaces_withQuote() to remove additional spaces
	word = Remove_spaces_withQuote(word)
	output := []byte(word)
	OurData := os.Args[2]
	words := os.WriteFile(OurData, output, 0o644)
	Check(words)
	fmt.Println("Done! you can see the file modifid now")
}
