package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
)

var input []string
var ids []int

func main() {
	file, _ := os.Open("input.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		input = append(input, scanner.Text())
	}

	for _, v := range input {
		n := row(v[0:7])*8 + column(v[7:10])
		ids = append(ids, n)
	}

	sort.Ints(ids)

	for i, v := range ids {
		if i == len(ids)-1 {
			break
		}
		if ids[i+1] == v+2 {
			fmt.Println(v + 1)
		}
	}
}

func row(code string) int {
	var value int
	for i, v := range code {
		var mult float64
		if string(v) == "B" {
			mult = 1
		}
		inc := int(math.Pow(2, (6-float64(i))) * mult)
		// fmt.Println(inc)
		value += inc
	}
	// fmt.Println(value)
	return value
}

func column(code string) int {
	var value int
	for i, v := range code {
		var mult float64
		if string(v) == "R" {
			mult = 1
		}
		inc := int(math.Pow(2, (2-float64(i))) * mult)
		// fmt.Println(inc)
		value += inc
	}
	// fmt.Println(value)
	return value
}
