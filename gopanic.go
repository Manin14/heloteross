package main

import (
	"fmt"
)

func cobaPanic() error {

	if false {

	} else {
		defer fmt.Println("nongol donk duluan")

		panic("nih")
		fmt.Println("panik ganyn")
	}
	return nil
}

func catch() {
	if r := recover(); r != nil {
		fmt.Println("error occured", r)
	} else {
		fmt.Println("Applicaion running perfectly")
	}
}
