package main

import (
	"io/ioutil"
	"log"
)

func main() {
	d, e := ioutil.ReadFile("data.txt")

	if e != nil {
		log.Fatal(e)
	}

	items := string(d)
}