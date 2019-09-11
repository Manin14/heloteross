package main

import "fmt"

func CobaMap() {
	var chicken map[string]int
	var mapString = map[string]string{
		"one":    "1",
		"two":    "2",
		"theree": "3",
		"four":   "4",
	}

	chicken = map[string]int{}

	chicken["keys1"] = 50
	chicken["keys2"] = 40

	for key, value := range mapString {
		fmt.Println(key, value)
	}
	//before i delete this keys
	fmt.Println(mapString)
	//lets delete this and guys
	delete(mapString, "one")
	fmt.Println(mapString)

	//check coba ya
	//bingung dah
	var value, check = mapString["one"]

	if check {
		fmt.Println("You found the value is ", value, "and check is : ", check)
	} else {
		fmt.Println("Realy ? no ", value, "and check is : ", check)
	}

	var data = []map[string]string{
		{"key1.1": "val1"},
		{"key1.2": "val2"},
		{"key1.3": "val3"},
	}
	fmt.Println(data)
	//netsted freakin data structure XD
	for _, values := range data {
		for keys22, values22 := range values {
			fmt.Println(keys22, "=>", values22)

		}
	}

}
