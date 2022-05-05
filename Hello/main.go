package main

import "fmt"

const englishHelloPrefix = "Hello "
const spanishHelloPrefix = "Hola "
const frenchHelloPrefix = "Bonjour "

func Hello(n, ln string) string {
	if n == "" {
		return "Hello, World"
	}
	switch ln {
	case "Spanish":
		return spanishHelloPrefix + n
	case "French":
		return frenchHelloPrefix + n
	default:
		return englishHelloPrefix + n
	}
}

func main() {
	fmt.Println(Hello("Luke", ""))
}
