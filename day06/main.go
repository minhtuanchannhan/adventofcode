package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

func main() {
	d, e := ioutil.ReadFile("data.txt")

	if e != nil {
		log.Fatal(e)
	}

	items := d[:len(d)-1]

	fmt.Println("First marker: ", checkMark(items, 4))
	fmt.Println("Second marker: ", checkMark(items, 14))
}

func checkMark(data []byte, cI int) int {
	for i := 0; i < len(data) - cI; i++ {
		mark := make(map[byte]bool)

		for j := 0; j < cI; j++ {
			mark[data[i+j]] = true
		}

		if len(mark) == cI {
			return i + cI
		}
	}

	return -1
}