package main

import (
	"strings"
)

// mirip  slice
func rataRata(deret ...int) float64 {
	var total int = 0
	for _, baris := range deret {
		total += baris
	}
	var avg = float64(total) / float64(len(deret))
	return avg
}

func returnString(kumpulan ...string) string {
	return strings.Join(kumpulan, ",")
}

//closure with variadic
var mouse = func(numbers []int) []int {
	// var arr = numbers
	return numbers
}

//closure with variadic
var mouse2 = func(numbers ...int) []int {
	// var arr = numbers
	return numbers
}
