package main

import (
	"fmt"
)

func main() {
	fmt.Print("hellow")
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

}
