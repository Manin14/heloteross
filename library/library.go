package librarylagi

import (
	f "fmt"
)

//fungsi ini akan otomatis di panggil pertama kali ketika aplikasi ini runing
//
func init() {
	f.Println("@./library : fungsi ini akan otomatis di panggil pertama kali ketika aplikasi ini runing")
}

func Helloworld() {
	f.Println("Helloworld@librarylagi")
}
