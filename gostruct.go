package main

import (
	"fmt"
)

func cobaStruct() {

	type mahasiswa struct {
		nim   int
		nama  string
		kelas string
	}

	var mhs mahasiswa
	mhs.nim = 12345
	mhs.nama = "rifky azmi"
	mhs.kelas = "S7N"

	fmt.Println("nim\t", mhs.nim)
	fmt.Println("nama\t", mhs.nama)
	fmt.Println("kelas\t", mhs.kelas)

}

func cobaStruct2() {

	type mahasiswa struct {
		nim   int
		nama  string
		kelas string
	}

	var mhs1 = mahasiswa{245, "rifky", "s2n"}
	var mhs2 = mahasiswa{246, "rifkya", "s3n"}
	var mhs3 = mahasiswa{247, "rifkyb", "s4n"}
	fmt.Println("|--------------------------------------|")
	fmt.Println("| NIM\t | nama\t| kelas \t|")
	fmt.Println("|--------------------------------------|")
	fmt.Println(mhs1.nim, "\t", mhs1.nama, "\t", mhs1.kelas)
	fmt.Println(mhs2.nim, "\t", mhs2.nama, "\t", mhs2.kelas)
	fmt.Println(mhs3.nim, "\t", mhs3.nama, "\t", mhs3.kelas)
	fmt.Println("|--------------------------------------|")
}

func cobaStruct3() {

	type mahasiswa struct {
		nim   int
		nama  string
		kelas string
	}

	var data1 = mahasiswa{
		nim:   1234,
		nama:  "qnoy",
		kelas: "s2g",
	}
	var data2 *mahasiswa = &data1

	fmt.Println(data1.nama)
	fmt.Println(data2)
}

func cobaEmbeded() {
	type alamat struct {
		rt   int
		rw   int
		kota string
	}
	type keluarga struct {
		kepalaKeluarga string
		pekerjaan      string
		noTlp          int
		alamat
	}
	var klg = keluarga{}
	klg.kepalaKeluarga = "bapak"
	klg.noTlp = 123456
	klg.alamat.rt = 5
	klg.alamat.rw = 4
	klg.alamat.kota = "jakarta"

	fmt.Println(klg)

	var infoAlamat = alamat{
		rt:   2,
		rw:   8,
		kota: "bandung",
	}
	var keluargaBaru = keluarga{
		kepalaKeluarga: "bapak2",
		alamat:         infoAlamat,
	}
	fmt.Println(keluargaBaru)
}

func cobaAnonStruct() {
	var anon = struct {
		//deklarasinya disini
		key1 int
		key2 string
		key3 bool
	}{
		//disini properti nya
	}
	fmt.Println(anon)
}
func withSliceAndMap() {
	type strucnya struct {
		key11 string
		key22 bool
		key33 int
	}
	var listData = []strucnya{
		{
			key11: "ini isi key11",
			key22: true,
			key33: 201,
		},
		{
			key11: "ini isi key12",
			key22: true,
			key33: 202,
		},
		{
			key11: "ini isi key13",
			key22: true,
			key33: 203,
		},
	}
	// fmt.Println(listData)
	for _, each := range listData {
		fmt.Println(each.key11)
	}
}
