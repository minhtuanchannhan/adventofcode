package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"sort"
	"strconv"
	"strings"
)


func main() {
	d, e := ioutil.ReadFile("data.txt")

	if e != nil {
		log.Fatal(e)
	}

	c := string(d)
	cl := [][]int{}
	sumCl := []int{}

	for _, g := range strings.Split(c, "\n\n") {
		r := []int{}

		for _, s := range strings.Split(g, "\n") {
			s, _ := strconv.Atoi(s)
			r = append(r, s)
		}

		cl = append(cl, r)
	}

	for _, v := range cl {
		sum := 0

		for _, z := range v {
			sum = sum + z
		}

		sumCl = append(sumCl, sum)
	}

	sort.Sort(sort.Reverse(sort.IntSlice(sumCl)))

	fmt.Println("Total calories: ", sumCl[0])
	fmt.Printf("Total top 3 calories of elves: %d", sumCl[0] + sumCl[1] + sumCl[2])
}