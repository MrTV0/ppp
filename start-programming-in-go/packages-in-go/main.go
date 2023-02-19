package main

import (
	"log"

	"github.com/mrtv0/myniceprogram/helpers"
)

func main() {
	var myVar helpers.SomeType
	myVar.TypeName = "Some name"
	log.Println(myVar.TypeName)
}