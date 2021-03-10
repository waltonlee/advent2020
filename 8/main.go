package main

import (
	"bufio"
	"fmt"
	"os"
)

var input []string

var acc int

func main() {
	file, _ := os.Open("input.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		input = append(input, scanner.Text())
	}

	var op string
	var arg int
	seen := make(map[int]bool)
	i := 0
	for {
		if seen[i] {
			fmt.Println(acc);
			return
		}
		seen[i] = true

		_, err := fmt.Sscanf(input[i], "%s %d", &op, &arg)
		if err != nil {
			fmt.Printf("err: scan %d, %s\n", i, err)
		}

		switch(op) {
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