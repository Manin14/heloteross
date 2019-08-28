package main

import (
	"fmt"
)

func main() {
	fmt.Print("helloworld XD")
	var slicenih = []string{
		"aduh",
		"katanya",
		"lol",
	}
	for _, tampung := range slicenih {
		fmt.Println(tampung)
	}

	CobaMap()
	var arrString = []string{
		"aku",
		"saya",
		"gua",
		"gue",
	}
	//join array string with function in another file
	setStringMessages(arrString)

	// mau coba fungsi pake return value nih ,
	hasilnya := OperasiMtkTambahan(1, 1)

	fmt.Println(hasilnya) //harusnya si hasilnya 2 XD

	//coba random angka
	// var count int = 10
	// for index := 1; index <= count; index++ {
	// 	angkaRandom := randomAngka(index, index)
	// 	fmt.Println(angkaRandom)
	// 	if index == count {
	// 		break
	// 	}
	// }

	//coba math pow
	Powpow()
	//coba akar
	var angkanya float64 = 16
	kuadrat := Sqrtsqrt(angkanya)
	fmt.Println("hasil sqrt(", angkanya, ")", kuadrat)

	//tipedata , bingung ya ? wkkwkw
	var tipe1 uint8 //255
	tipe1 = 255
	fmt.Println(tipe1)
}
