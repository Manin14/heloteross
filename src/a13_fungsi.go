package main

import (
	"fmt"
	"strings"
)

func setStringMessages(arr []string) {
	//ini mirip implode di php cuk XD
	var localString = strings.Join(arr, " ")
	fmt.Println(localString)
}
