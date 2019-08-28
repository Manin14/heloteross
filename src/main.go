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
	var count int = 100
	for index := 1; index < count; index++ {
		angkaRandom := randomAngka(index, index)
		fmt.Println("dapet berapa ya => ", angkaRandom)
	}

}
