package main

import "fmt"

func cobaDefer() {
	fmt.Println("ini text pertama ya ")
	defer fmt.Println("ini text kedua") //akan di eksekusi paling terakhir di dalem fungsi ini
	fmt.Println("ini text ke tiga ya")
}

func checkSaya(nama string) {
	defer fmt.Println("nama kamu rifky")
	if nama == "imza" {
		fmt.Println("nama yang di input adalah :", nama)
		return

	}
}

//dijalankan berurutan dari statment yang di defer paling akhir/paling bawah ke atas
func cobaBanyakDefer() {
	defer fmt.Println("defer 3")
	for i := 0; true; i++ {
		if i == 10 {
			break
		}
		fmt.Println("defer banyak :D ", i)
	}
	defer fmt.Println("defer 1")
	defer fmt.Println("defer 2")
}
