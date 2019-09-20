package main

import (
	"fmt"
	"runtime"
)

func cetak(n int, message string) {
	runtime.GOMAXPROCS(2)
	for i := 0; i <= n; i++ {
		fmt.Println((i + 1), message)
		angkaRandom := randomAngka(i, (i + 1))
		go fmt.Println(angkaRandom)
	}
}
