package main

import (
	"fmt"
)

type structnya2 struct {
	nama  string
	kelas string
	hobi  []string
}

type listDatanm []struct {
	listnama []string
}

func (s structnya2) getName() string {
	return s.nama
}
func (s *structnya2) changeName(newName string) {
	s.nama = newName
}
func (s *structnya2) getOffsets() {
	var tampungan *string
	tampungan = &s.nama
	fmt.Println(tampungan)
}
