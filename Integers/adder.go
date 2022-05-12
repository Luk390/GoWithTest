package main

func Adder(xi ...int) int {
	sum := 0
	for _, v := range xi {
		sum += v
	}
	return sum
}
