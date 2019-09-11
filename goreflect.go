package main

import (
	"fmt"
	"reflect"
)

func cobaReflect() {
	var varnya = 10
	var checkRef = reflect.ValueOf(varnya)
	//Type() return string
	fmt.Println(">>", checkRef.Type())
	if checkRef.Kind() == reflect.Int {
		fmt.Println("=>", checkRef.Int())
	}
	fmt.Println("", checkRef.Interface())
	var nilaiVarnya = checkRef.Interface().(int)
	fmt.
}
