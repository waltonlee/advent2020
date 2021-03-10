package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var input []string

type contents struct {
	name   string
	amount int
}

var bags map[string][]contents

func main() {
	bags = make(map[string][]contents)
	file, _ := os.Open("input.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		words := strings.Split(scanner.Text(), " ")
		newBag := strings.Join([]string{words[0], words[1]}, " ")
		bags[newBag] = []contents{}
		if len(words) > 7 {
			for i := 0; i < (len(words)-4)/4; i++ {
				amt, _ := strconv.Atoi(words[4+i*4])
				newContents := contents{name: strings.Join([]string{words[5+i*4], words[6+i*4]}, " "), amount: amt}
				bags[newBag] = append(bags[newBag], newContents)
			}
		}
	}

	fmt.Println(count(bags["shiny gold"]))
}

func count(input []contents) (sum int) {
	for _, c := range input {
		fmt.Printf("%d %s bags\n", c.amount, c.name)
		sum += c.amount
		sum += c.amount * count(bags[c.name])
	}
	return sum
}
