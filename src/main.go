package main

import (
	"fmt"
	"strings"
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

	//2 nilai
	var nilai1, nilai2 = duaReturn()
	fmt.Println("nilai1 = ", nilai1)
	fmt.Println("nilai2 = ", nilai2)

	//lingkaran
	var luasLingkaran, kelilingLingkaran = lingkaran(4)
	fmt.Println("luas lingkaran : ", luasLingkaran)
	fmt.Println("keliling lingkaran : ", kelilingLingkaran)

	//variadic
	//dengan map
	var deretAngka = []int{1, 2, 3, 4, 5, 6, 7, 8}
	var ratarata = rataRata(deretAngka...)
	fmt.Println(ratarata)

	//return string
	var stringnya = returnString("satu", "satu", "dua")
	fmt.Println(stringnya)

	//closure with variadic
	var listAngka = []int{1, 2, 3}
	// var mouse = ([]int{1, 2, 3})
	var mouse2 = (listAngka)
	fmt.Println(mouse2)

	//closure style
	fmt.Println("isi dari iife fnc => ", iife)

	//coba fungsi params
	var iniData = []string{
		"data1",
		"data2",
		"data3",
	}
	var RunFcFilter = filter(iniData, func(each string) bool {
		fmt.Println("isi dari each => ", each)
		return strings.Contains(each, "1")
	})
	fmt.Println(RunFcFilter)

	//coba pointer
	fmt.Println("------------------------------------------ pointer")
	cobaPointer()

	//coba struct
	cobaStruct()
	cobaStruct2()
	cobaStruct3()
	cobaEmbeded()
	cobaAnonStruct()
	withSliceAndMap()
	sliceAnonStruct()
	nestedStruct()
}
