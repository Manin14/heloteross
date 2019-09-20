package main

import (
	"fmt"
	"os"
)

//meskipun ada statemen yang di defer jika ada baris exit maka statement setelahnya tidak akan di exsekusi
//
func cobaExit() {
	defer fmt.Println("ini adalah yang terakhir :( ")
	fmt.Println("percobaan exit ,, emmm")
	os.Exit(1)
	fmt.Println("halo hari ini adalah hari ...")

}
