package main

func Sum(xi []int) int {
	total := 0
	for _, v := range xi {
		total += v
	}
	return total
}

func SumAll(xxi ...[]int) []int {
	var sums []int
	for _, v := range xxi {
		sums = append(sums, Sum(v))
	}
	return sums
}
func SumAllTails(xxi ...[]int) []int {
	var sums []int
	for _, v := range xxi {
		if len(v) == 0 {
			sums = append(sums, 0)
		} else {
			sums = append(sums, Sum(v[1:]))
		}
	}
	return sums
}