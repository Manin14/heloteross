package main

import (
	"fmt"
)

func cobaPointer() {

	var varNilai string = "ini isi dari varNilai"
	var penampungan *string // tipe data untuk  penampung pointer

	fmt.Println("varNilai=", varNilai)
	fmt.Println("penampungan=", penampungan)

	//hint
	// gunakan tanda/symbol ampersand & pada nama variabel
	// untuk mengambil alamat/pointer dari variabel tersebut
	fmt.Println("ambil alamat varNilai masukan ke dalam penampungan")
	penampungan = &varNilai

	fmt.Println("alamat varNilai:", penampungan, "berisi=>", varNilai)
	fmt.Println("penampungan=", penampungan)
	fmt.Println("&varNilai", &varNilai)

}

//pointer with function
