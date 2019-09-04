package main

import (
	"fmt"
	"math"
)

//fungsi ini akan otomatis di panggil pertama kali ketika aplikasi ini runing
//
func init() {
	fmt.Println("@gomath : fungsi ini akan otomatis di panggil pertama kali ketika aplikasi ini runing")
}
func Powpow() {

	var a, b float64
	a = 2
	b = 3
	fmt.Println(math.Pow(a, b))
}

func Sqrtsqrt(angka float64) float64 {
	return math.Sqrt(angka)
}

func duaReturn() (string, string) {
	localVar1 := "aku"
	localVar2 := "dan"
	return localVar1, localVar2
}

func lingkaran(r float64) (float64, float64) {
	var diameter = r * 2
	//Luas = π × r²
	luas := math.Pi * math.Pow(r, 2)
	// Keliling = π × d
	keliling := math.Pi * diameter
	return luas, keliling
}
