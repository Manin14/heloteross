package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

func checkNumber() {
	var input string
	fmt.Print("type some number : ")
	fmt.Scanln(&input)

	var number int
	var err error
	number, err = strconv.Atoi(input)
	if err == nil {
		fmt.Println(number, "is a number")
	} else {
		fmt.Println(input, "is not a number")
		//throw error here
		fmt.Println("ini error nya", err.Error())
	}
}
func cobaCustomError(input string) (bool, error) {
	if strings.TrimSpace(input) == "" {
		return false, errors.New("tidak boleh kosong ya")
	}
	return true, nil
}
