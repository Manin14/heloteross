package main

//test git stash
import (
	m "math"
)

//jika tidak medhod tidak ada dalam block interface maka pemanggilan method
// ke tipe asli variabel konkritnya
type hitung interface {
	luas() float64
	keliling() float64
}

type lingkaran2 struct {
	diameter float64
}
type persegi2 struct {
	panjang float64
	lebar   float64
	sisi    float64
}

func (l lingkaran2) jariJari() float64 {
	return l.diameter / 2
}
func (l lingkaran2) luas() float64 {
	return m.Pi * m.Pow(l.jariJari(), 2)
}

func (l lingkaran2) keliling() float64 {
	return m.Pi * l.diameter
}

func (p persegi2) luas() float64 {
	return p.panjang * p.lebar
}
func (p persegi2) keliling() float64 {
	return 4 * p.sisi
}
func (p persegi2) carisisi() float64 {
	return p.keliling() / 4
}

//embeded interface
type hitung2d interface {
	luas() float64
	keliling() float64
}
type hitung3d interface {
	volume() float64
}
type hitung2 interface {
	hitung2d
	hitung3d
}
type kubus struct {
	sisi float64
}

func (k kubus) volume() float64 {
	return m.Pow(k.sisi, 3)
}
func (k kubus) luas() float64 {
	return m.Pow(k.sisi, 2) * 6
}
func (k kubus) keliling() float64 {
	return k.sisi * 12
}
