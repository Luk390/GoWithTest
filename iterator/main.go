package main

func Repeat(s string) string {
	var repeatedString string
	for i := 0; i < 5; i++ {
		repeatedString += s
	}
	return repeatedString
}