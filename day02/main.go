package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

func main() {
	d, e := ioutil.ReadFile("data.txt")

	if e != nil {
		log.Fatal(e)
	}

	firstSum := 0
	secondSum := 0
	firstPoints := map[string]int{"A X": 4, "A Y": 8, "A Z": 3, "B X": 1, "B Y": 5, "B Z": 9, "C X": 7, "C Y": 2, "C Z": 6}
	secondPoints := map[string]int{"A X": 3, "A Y": 4, "A Z": 8, "B X": 1, "B Y": 5, "B Z": 9, "C X": 2, "C Y": 6, "C Z": 7}

	for _, c := range strings.Split(string(d), "\n") {
		firstSum += firstPoints[c]
		secondSum += secondPoints[c]
	}

	fmt.Println("Total score: ", firstSum)
	fmt.Println("Exactly total score: ", secondSum)
}