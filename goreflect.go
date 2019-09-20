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
	fmt.Println("value nilaiVarnya=>", nilaiVarnya)
}

type Students struct {
	Name  string
	Grade int
}

func (s *Students) getPropertyInfo() {
	var reflectValue = reflect.ValueOf(s)
	fmt.Println(reflectValue)
	if reflectValue.Kind() == reflect.Ptr {
		reflectValue = reflectValue.Elem()
	}

	var reflectType = reflectValue.Type()
	for i := 0; i < reflectValue.NumField(); i++ {
		fmt.Println("nama\t:", reflectType.Field(i).Name)
		fmt.Println("tipe data\t:", reflectType.Field(i).Type)
		fmt.Println("nilai\t:", reflectValue.Field(i).Interface())

	}
}
func (s *Students) SetName(name string, grade int) {
	s.Name = name
	s.Grade = grade
}
