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

	items := strings.Split(string(d), "\n")

	sum := 0
	sumG :=0

	for _, c := range items {
		firstRuckSack := c[:len(c)/2]
		secondRuckSack := c[len(c)/2:]
		inFirstRuckSack := make(map[rune]bool)

		for _, s := range firstRuckSack {
			inFirstRuckSack[s] = true
		}

		for _, s := range secondRuckSack {
			if inFirstRuckSack[s] {
				if s >= 'a' && s <= 'z' {
					sum += int(s-'a') + 1
				} else {
					sum += int(s-'A') + 27
				}

				break
			}
		}
	}

	for i:=0; i < len(items)-1; i += 3 {
		newMap := make(map[rune]int)

		for j:=0; j<3; j++ {
			currentMap := make(map[rune]bool)
			
			for _, k := range items[i+j] {
				if currentMap[k] {
					continue
				}

				currentMap[k] = true
				newMap[k]++
			}
		}

		for s, v := range newMap {
			if v == 3 {
				if s >= 'a' && s <= 'z' {
					sumG += int(s-'a') + 1
				} else {
					sumG += int(s-'A') + 27
				}
			}
		}
	}

	fmt.Println("Total of priorities: ", sum)
	fmt.Println("Total of priorities by group: ", sumG)
}