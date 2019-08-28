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
