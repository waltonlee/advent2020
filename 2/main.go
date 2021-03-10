package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	input := []string{}

	file, _ := os.Open("input.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		input = append(input, scanner.Text())
	}

	var count int

	for _, v := range input {
		var min, max int
		var key, password string
		fmt.Sscanf(v, "%d-%d %1s: %s", &min, &max, &key, &password)

		if (string(password[min-1]) == key) != (string(password[max-1]) == key) {
			count++
		}
	}
	fmt.Print(count)
}
