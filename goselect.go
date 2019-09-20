package main

import "fmt"

func findMax(isi chan int) {
	var deret = []int{1, 2, 0, 4, 5}
	var max = deret[0]

	for _, dangka := range deret {
		fmt.Println("max = ", max, "<", dangka, (max < dangka))

		if max < dangka {
			max = dangka
		}
	}
	isi <- max
}
func RemoveIndex(s []int, index int) []int {
	return append(s[:index], s[index+1:]...)
}
func findNol() []int {
	var deret = []int{1, 5, 0, 3}
	for _, each := range deret {
		if each == 0 {
			// fmt.Println("found ! at index ", index)
			RemoveIndex(deret, 0)
		}
	}
	return deret
}
