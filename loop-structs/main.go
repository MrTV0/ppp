package main

import "log"

func main() {
	var firstLine = "Once upon a timeeerrrreeeee"

	for i, l := range firstLine {
		log.Println(i, ":", l)
	}
}
