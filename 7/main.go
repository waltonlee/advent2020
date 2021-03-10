package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var input []string

var bags map[string][]string

func main() {
	bags = make(map[string][]string)
	file, _ := os.Open("input.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		// input = append(input, scanner.Text())
		words := strings.Split(scanner.Text(), " ")
		newBag := strings.Join([]string{words[0], words[1]}, " ")
		bags[newBag] = []string{}
		if len(words) > 7 {
			for i := 0; i < (len(words)-4)/4; i++ {
				contents := strings.Join([]string{words[5+i*4], words[6+i*4]}, " ")
				bags[newBag] = append(bags[newBag], contents)
			}
		}
	}

	var count int

	for _, v := range bags {
		if hasGold(v) {
			count++
		}
	}
	fmt.Println(count)
}

func hasGold(input []string) bool {
	if len(input) == 0 {
		return false
	}
	for _, v := range input {
		if v == "shiny gold" {
			return true
		}
		if hasGold(bags[v]) {
			return true
		}
	}
	return false
}
