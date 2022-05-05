package main

import "fmt"

const englishHelloPrefix = "Hello "

func Hello(n string) string {
	if n == "" {
		return "Hello, World"
	}
	return englishHelloPrefix + n
}

func main() {
	fmt.Println(Hello("Luke"))
}