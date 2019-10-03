package main

import (
	"fmt"
)

type dataMhs struct {
	nama   string
	kelas  string
	matkul []string
	age    int
	lulus  bool
}

func showdataMhs() {
	var data = dataMhs{
		nama:   "rifky",
		kelas:  "s7n",
		matkul: []string{"filsafat", "riset"},
		age:    25,
		lulus:  true,
	}

	fmt.Printf("%b\n", data.age) //decimal to binner (cast string)
	/*
		128 64 32 16  8 4 2 1
		0 0 1 1 0 0 1
	*/
	//%c numeric to unicode
	fmt.Printf("%c\n", 82) //output R

	//%d numeric to string
	fmt.Printf("%d\n", data.age) //(cast string)

	//%e dan %E decimal to numeric standart scientic
	fmt.Printf("%E\n", 182.5) //1.825000E+02 => 182 x 10^2
	fmt.Printf("%e\n", 3.14)

	//%f dan %F untuk memformat angka di belakang koma dari sebuah nilai decimal (default 6 digit )
	fmt.Printf("%f\n", 182.5111111111111)
	fmt.Printf("%.2f\n", 182.5) //2 angka di belakang koma
	fmt.Printf("%.2F\n", 182.5) //2 angka di belakang koma

	//%g dan %G sama seperti e dan f perbedaannya angka di belakang koma bisa banyak tapi tidak bisa di custom spt di format e
	fmt.Printf("%g\n", 192.5111111111111)
	fmt.Printf("%G\n", 192.5000000000000)

	//decimal to octal (cast string) output : 31
	fmt.Printf("%o\n", data.age)
	/*
		mod = hasil bagi
		25 mod 8 = 3 sisa 1
		3 mod 8 = 0 sisa 3
		urutkan dari bawah ke atas
		31
	*/

	//%p memformat data pointer , output hexadecimal dengan prefix 0x
	// output : 0xc000056178
	fmt.Printf("%p\n", &data.age)

	//%q untuk escape string
	fmt.Printf("%q\n", `rifky\ "," azmi`)

	//%s memformat data string
	fmt.Printf("%s\n", data.nama)

	//%t memformat data boolean , cast bool
	fmt.Printf("%t\n", data.lulus)

	//%T mengambil tipedata dari data yg di format ,simplenya untuk check tipe datanya XD
	fmt.Printf("%T\n", data.lulus)
	fmt.Printf("%T\n", data.nama)
	fmt.Printf("%T\n", data.matkul)

	//%v menampilkan nilai properti dari tipe data apa saja secara berurutan
	fmt.Printf("%v\n", data)

	//%+v mirip seperti %v perbedaannya menampilkan nama propertinya juga hanya tipe data struct saja
	fmt.Printf("%+v\n", data)

	//%#v mirip seperti %+v perbedaannya structur nya juga di tampilkan
	fmt.Printf("%#v\n", data)

	var data2 = struct {
		name   string
		height float64
	}{
		name:   "wick",
		height: 182.5,
	}

	fmt.Printf("%#v\n", data2)

	//%x dan %X decimal to hexadecimal ,perbedaannya di outputnya yaitu x = huruf kecil ,X huruf besar
	fmt.Printf("%x\n", data.age)
	/*
		mod = hasil bagi
		25 mod 16 = 1 sisa 9
		1 mod 16 = 0 sisa 1
		urutkan dari bawah ke atas
		19
	*/
	fmt.Printf("%x\n", 26) // output 1a
	fmt.Printf("%X\n", 26) //output 1A

	//%% untuk menulis caaracter pada string format
	fmt.Printf("%%\n")
}
