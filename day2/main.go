package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func captureLines(path string, f func(v string)) {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		f(scanner.Text())
	}
}

func main() {
	// part a
	depth := 0
	horizontal := 0
	aim := 0

	captureLines("input.txt", func(v string) {
		line := strings.Split(v, " ")
		x, err := strconv.Atoi(line[1])
		if err != nil {
			panic(err)
		}

		switch line[0] {
		case "forward":
			horizontal += x
		case "down":
			depth += x
		case "up":
			depth -= x
		}
		//fmt.Println(line, depth*horizontal, horizontal, depth)
	})
	fmt.Println("a:", depth*horizontal)

	// part b
	depth = 0
	horizontal = 0
	captureLines("input.txt", func(v string) {
		line := strings.Split(v, " ")
		x, err := strconv.Atoi(line[1])
		if err != nil {
			panic(err)
		}

		switch line[0] {
		case "forward":
			horizontal += x
			depth += aim * x
		case "down":
			aim += x
		case "up":
			aim -= x
		}
		// fmt.Println(line)
		// fmt.Println(depth*horizontal, horizontal, depth, aim)
	})
	fmt.Println("b:", depth*horizontal)
}
