package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func main() {
	d, e := ioutil.ReadFile("./data.txt")

	if e != nil {
		log.Fatal(e)
	}

	var st int
	var px [240]bool
	x := 1
	pc := 1
	ct := string(d)
	sp := strings.Split(ct, "\n")

	for _, s := range sp[:len(sp)-1] {
		if pc%40 == 20 {
			st += x * pc
		}

		if x >= ((pc-1)%40)-1 && x <= ((pc-1)%40)+1 {
			px[pc-1] = true
		}

		str := strings.Split(string(s), " ")
		switch str[0] {
		case "addx":
			val, _ := strconv.Atoi(str[1])
			pc++
			if pc%40 == 20 {
				st += x * pc
			}
			if x >= ((pc-1)%40)-1 && x <= ((pc-1)%40)+1 {
				px[pc-1] = true
			}
			pc++
			x += val
		case "noop":
			pc++
		}
	}

	fmt.Println(st)

	for i, v := range px {
		if v {
			fmt.Printf("#")
		} else {
			fmt.Printf(".")
		}

		if i%40==39 {
			fmt.Println()
		}
	}
}