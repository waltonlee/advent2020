package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	input := []int{}

	file, _ := os.Open("input.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line, _ := strconv.Atoi(scanner.Text())
		input = append(input, line)
	}

	for i, v := range input {
		for j, w := range input {
			if i != j && v+w == 2020 {
				fmt.Printf("%v\n", v*w)
				return
			}
		}
	}
}
