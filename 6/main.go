package main

import (
	"bufio"
	"fmt"
	"os"
)

var input []string

func main() {
	file, _ := os.Open("input.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		input = append(input, scanner.Text())
	}

	var count, groupCount int
	yesses := make(map[string]int)

	for _, v := range input {
		if v == "" {
			for _, y := range yesses {
				if y == groupCount {
					count++
				}
			}
			groupCount = 0
			yesses = make(map[string]int)
		} else {
			groupCount++
			for _, c := range v {
				yesses[string(c)]++
			}
		}
	}
	for _, y := range yesses {
		if y == groupCount {
			count++
		}
	}

	fmt.Println(count)
}
