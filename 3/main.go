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

	fmt.Print(slopeCalc(1, 1) * slopeCalc(3, 1) * slopeCalc(5, 1) * slopeCalc(7, 1) * slopeCalc(1, 2))
}

func slopeCalc(x int, y int) int {
	var count int
	for i := 1; i*y < len(input); i++ {
		if string(input[i*y][(i*x)%len(input[0])]) == "#" {
			count++
		}
	}
	return count
}
