package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func setStringMessages(arr []string) {
	//ini mirip implode di php cuk XD
	var localString = strings.Join(arr, " ")
	fmt.Println(localString)
}

/**
* @param1 int
* @param2 int
* @returnn int
 */
func OperasiMtkTambahan(a int, b int) int {
	var hasil = a + b
	return hasil

}

//coba keyword lain yak
//ref https://golang.cafe/blog/golang-random-number-generator.html
func randomAngka(min int, max int) int {
	rand.Seed(time.Now().UnixNano())
	Angka := rand.Intn(max - min + 1 + min)
	return Angka

}

//coba return map ya
func retrunMap() map[string]string {
	var mapString = map[string]string{
		"one":    "1",
		"two":    "2",
		"theree": "3",
		"four":   "4",
	}
	return mapString
}

//coba return array ya
func retrunArr() []string {
	var arrString = []string{
		"aku",
		"saya",
		"gua",
		"gue",
	}
	return arrString
}

func filter(data []string, callback func(string) bool) []string {
	var result []string
	for keys, each := range data {
		filtered := callback(each)
		fmt.Print(keys, each, " ")

		fmt.Println("search ", filtered)
		if filtered {
			fmt.Println("got it !")
			result = append(result, each)
		}

	}
	return result
}
