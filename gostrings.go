package main

import (
	"strings"
)

var iniStringnya string = "jakarta adalah ibukota indonesia" //hmmm
//mirip explode di php duarrrrrrr *** XD
var pecahDuarr = strings.Split(iniStringnya, " ") //di pisah pake sperator spasi

//mirip implode di php hhmmm
var gabungIn = strings.Join(pecahDuarr, ",") //test gabungin pake koma
