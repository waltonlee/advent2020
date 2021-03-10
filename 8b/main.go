package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var input []string

func main() {
	file, _ := os.Open("input.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		input = append(input, scanner.Text())
	}

	// fmt.Println(testLoop(input))

	for i := 0; i < len(input); i++ {
		testInput := make([]string, len(input))
		copy(testInput, input)
		op, _ := getOp(testInput[i])
		switch op {
		case "nop":
			testInput[i] = strings.Replace(testInput[i], "nop", "jmp", 1)
		case "jmp":
			testInput[i] = strings.Replace(testInput[i], "jmp", "nop", 1)
		default:
			continue
		}

		valid, acc := testLoop(testInput)
		if valid {
			fmt.Println(acc)
		}
	}
}

func testLoop(testInput []string) (bool, int) {
	var acc int
	seen := make(map[int]bool)
	i := 0
	for {
		if i >= len(testInput) {
			return true, acc
		}
		if seen[i] {
			return false, acc
		}
		seen[i] = true

		op, arg := getOp(testInput[i])

		switch op {
		case "nop":
			i++
		case "jmp":
			i += arg
		case "acc":
			acc += arg
			i++
		}
	}
}

func getOp(ops string) (string, int) {
	var op string
	var arg int
	_, err := fmt.Sscanf(ops, "%s %d", &op, &arg)
	if err != nil {
		fmt.Printf("err: %s\n", err)
	}
	return op, arg
}
