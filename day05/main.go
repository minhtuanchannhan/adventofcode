package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func main() {
	d, e := ioutil.ReadFile("data.example.txt")

	if e != nil {
		log.Fatal(e)
	}

	items := strings.Split(string(d), "\n")


	for _, item := range items {
		moves := strings.Split(item, " ")
		count, _ := strconv.Atoi(moves[1])
		from, _ := strconv.Atoi(moves[3])
		to, _ := strconv.Atoi(moves[5])

		fmt.Println(count, from, to)

		
	}
}