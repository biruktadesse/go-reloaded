package goreloaded

// quits if there's an error
func Check(err error) {
	if err != nil {
		panic(err)
	}
}
