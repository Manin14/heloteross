package main

import (
	"fmt"
	// "strings"

	// . "github.com/mqnoy/belajar-golang-level-akses/library"

	// "math"
	// "reflect"
	"runtime"
	// "./library"
)

func main() {
	// fmt.Print("helloworld XD")
	// var slicenih = []string{
	// 	"aduh",
	// 	"katanya",
	// 	"lol",
	// }
	// for _, tampung := range slicenih {
	// 	fmt.Println(tampung)
	// }

	// CobaMap()
	// var arrString = []string{
	// 	"aku",
	// 	"saya",
	// 	"gua",
	// 	"gue",
	// }
	// //join array string with function in another file
	// setStringMessages(arrString)

	// // mau coba fungsi pake return value nih ,
	// hasilnya := OperasiMtkTambahan(1, 1)

	// fmt.Println(hasilnya) //harusnya si hasilnya 2 XD

	// //coba random angka
	// // var count int = 10
	// // for index := 1; index <= count; index++ {
	// // 	angkaRandom := randomAngka(index, index)
	// // 	fmt.Println(angkaRandom)
	// // 	if index == count {
	// // 		break
	// // 	}
	// // }

	// //coba math pow
	// Powpow()
	// //coba akar
	// var angkanya float64 = 16
	// kuadrat := Sqrtsqrt(angkanya)
	// fmt.Println("hasil sqrt(", angkanya, ")", kuadrat)

	// //tipedata , bingung ya ? wkkwkw
	// var tipe1 uint8 //255
	// tipe1 = 255
	// fmt.Println(tipe1)

	// //2 nilai
	// var nilai1, nilai2 = duaReturn()
	// fmt.Println("nilai1 = ", nilai1)
	// fmt.Println("nilai2 = ", nilai2)

	// //lingkaran
	// var luasLingkaran, kelilingLingkaran = lingkaran(4)
	// fmt.Println("luas lingkaran : ", luasLingkaran)
	// fmt.Println("keliling lingkaran : ", kelilingLingkaran)

	// //variadic
	// //dengan map
	// var deretAngka = []int{1, 2, 3, 4, 5, 6, 7, 8}
	// var ratarata = rataRata(deretAngka...)
	// fmt.Println(ratarata)

	// //return string
	// var stringnya = returnString("satu", "satu", "dua")
	// fmt.Println(stringnya)

	// //closure with variadic
	// var listAngka = []int{1, 2, 3}
	// // var mouse = ([]int{1, 2, 3})
	// var mouse2 = (listAngka)
	// fmt.Println(mouse2)

	// //closure style
	// fmt.Println("isi dari iife fnc => ", iife)

	// //coba fungsi params
	// var iniData = []string{
	// 	"data1",
	// 	"data2",
	// 	"data3",
	// }
	// var RunFcFilter = filter(iniData, func(each string) bool {
	// 	fmt.Println("isi dari each => ", each)
	// 	return strings.Contains(each, "1")
	// })
	// fmt.Println(RunFcFilter)

	// //coba pointer
	// fmt.Println("------------------------------------------ pointer")
	// cobaPointer()

	// //coba struct
	// cobaStruct()
	// cobaStruct2()
	// cobaStruct3()
	// cobaEmbeded()
	// cobaAnonStruct()
	// withSliceAndMap()
	// sliceAnonStruct()
	// nestedStruct()

	// //asdsd

	// var mahasiswa = structnya2{
	// 	nama:  "rifky",
	// 	kelas: "S7N",
	// 	hobi:  []string{"hobi12", "hobi22"},
	// }
	// var Lnm = listDatanm{
	// 	{
	// 		[]string{
	// 			"mouse",
	// 			"keyboard",
	// 			"monitor",
	// 		},
	// 	}, {
	// 		[]string{
	// 			"mouse2",
	// 			"keyboard2",
	// 			"monitor2",
	// 		},
	// 	},
	// }
	// var nmMhs = mahasiswa.getName()
	// fmt.Println("nama mahasiswanya adalah : ", nmMhs)

	// var namaBaru = "rifky azmi"
	// fmt.Println("coba ganti nama ", mahasiswa.nama, "ke", namaBaru)
	// mahasiswa.changeName(namaBaru)
	// fmt.Println("nama mahasiswanya adalah : ", mahasiswa.nama)
	// mahasiswa.getOffsets()
	// //slice struct
	// fmt.Println(Lnm)

	// //strings here :D
	// fmt.Println(pecahDuarr)
	// for _, eachDuar := range pecahDuarr {
	// 	fmt.Println(eachDuar)
	// }
	// ///
	// totArr := len(pecahDuarr)
	// fmt.Println("total", totArr)
	// for i := 0; i < totArr; i++ {
	// 	fmt.Println(pecahDuarr[i])
	// }

	// //
	// fmt.Println("setelah di gabungin \n", gabungIn)

	// //yg ada di a13
	// sayHello()

	// //panggil semua yg ada di library taro sini
	// SayHello()

	// //
	// librarylagi.Helloworld()

	// //coba interface
	// var bangungDatar hitung
	// bangungDatar = lingkaran2{10}
	// var jrjr = bangungDatar.luas()
	// fmt.Println(jrjr)
	// fmt.Println("jari jari di panggil dari struct lingkaran", bangungDatar.(lingkaran2).jariJari())

	// bangungDatar = persegi2{10, 20, 4}
	// // bangungDatar.panjang = 10
	// // bangungDatar.lebar = 20
	// // bangungDatar.sisi = 4
	// var lsp = bangungDatar.luas()
	// var kel = bangungDatar.keliling()
	// var carsisi = bangungDatar.(persegi2).carisisi()
	// fmt.Println(lsp)
	// fmt.Println(kel)
	// fmt.Println(carsisi)

	// var bangunRuang hitung2
	// bangunRuang = kubus{4}
	// var vol = bangunRuang.volume()
	// fmt.Println("volume kubus = ", vol)
	// fmt.Println("luas kubus = ", bangunRuang.luas())
	// fmt.Println("keliling kubus = ", bangunRuang.keliling())

	// //interface kosong

	// //aka dynamic typing
	// var bebas interface{}

	// bebas = "ini string"
	// fmt.Println(bebas)

	// bebas = 7.0

	// fmt.Println(bebas)
	// fmt.Println(math.Pow(bebas.(float64), 2)) //pangkat2

	// var dataMhs map[string]interface{}
	// dataMhs = map[string]interface{}{
	// 	"name":  "rifky azmi",
	// 	"kelas": "s7n",
	// 	"smt":   7,
	// 	"hobi":  []string{"hhmm", "hheemm"},
	// }
	// bebas = []string{
	// 	"aku",
	// 	"adalah",
	// 	"ceret",
	// }
	// fmt.Println(bebas)
	// fmt.Println(strings.Join(bebas.([]string), " "))
	// fmt.Println(dataMhs)
	// /*for keys, each := range dataMhs {
	// 	fmt.Println(keys, "=>", each)
	// 	// for _, hobies := range each["hobi"] {
	// 	fmt.Println(each["hobi"])
	// 	// }
	// }*/
	// type orang struct {
	// 	nama string
	// 	umur int
	// }
	// var bebas2 interface{} = orang{"azmi", 24}
	// var namanya = bebas2.(orang).nama
	// fmt.Println(namanya)

	// var aduh int = 100
	// var bebas3 interface{} = &aduh
	// var tampunganAduh *int = &aduh
	// fmt.Println(bebas3)
	// fmt.Println(tampunganAduh)

	// //reflect
	// var s11 = &Students{
	// 	Name:  "rifky",
	// 	Grade: 10,
	// }
	// s11.getPropertyInfo()
	// var reflectValue = reflect.ValueOf(s11)
	// var method = reflectValue.MethodByName("SetName")
	// fmt.Println([]reflect.Value{reflect.ValueOf("asd")})
	// method.Call(
	// 	[]reflect.Value{
	// 		//harus urut posisinya mengikuti parameter di method tersebut
	// 		reflect.ValueOf("azmi"),
	// 		reflect.ValueOf(100),
	// 	},
	// )
	// s11.getPropertyInfo()

	//goroutine
	//runtime.GOMAXPROCS(2)

	//go cetak(1, "halo")
	//cetak(2, "apa kabar")
	//fmt.Scanln()

	//coba chanel
	runtime.GOMAXPROCS(2)
	var pesanan = make(chan int)

	// go sendMessageCh(pesanan)
	go sendDataTm(pesanan)
	retreiveData(pesanan)
	// printMessageCh(pesanan)
	// var messages = make(chan string)
	// var sayHelloTo = func(who string) {
	// 	var data = fmt.Sprintf("hellow %s", who)
	// 	messages <- data
	// }
	// go sayHelloTo("rifky1")
	// go sayHelloTo("azmi2")
	// go sayHelloTo("rifky azmi3")

	// var message1 = <-messages
	// fmt.Println(message1)
	// var message2 = <-messages
	// fmt.Println(message2)
	// var message3 = <-messages
	// fmt.Println(message3)

	var pesan = make(chan string)

	for _, each := range []string{"satu", "dua", "tiga"} {
		go func(who string) {
			var data = fmt.Sprintf("hellow %s ", who)
			pesan <- data
		}(each)
	}
	for i := 0; i < 3; i++ {
		fmt.Println(pesan)

	}
	//coba chan buf
	cobaBuff()

	//coba select + testing chan
	var isian = make(chan int)
	go findMax(isian)
	select {
	case maxi := <-isian:
		fmt.Println(maxi)

	}
	var hasilnya = findNol()
	fmt.Println(hasilnya)
	var ituh []int

	for id := 0; id <= 10; id++ {
		ituh = append(ituh, id)
	}
	fmt.Println(ituh)

	// coba defer
	cobaDefer()
	checkSaya("imza")
	cobaBanyakDefer()

	//coba exit
	//cobaExit()
	cetakanEh("eh eh ")

	//coba error
	checkNumber()
}
