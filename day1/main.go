package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	depths := []int{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		i, err := strconv.Atoi(scanner.Text())
		if err != nil {
			panic(err)
		}
		depths = append(depths, i)
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	a := 0
	aPrev := depths[0]
	b := 0
	bPrev := depths[0]
	for i, v := range depths {
		sum := 0
		for _, x := range depths[i : i+3] {
			sum += x
		}
		if v > aPrev {
			a++
		}
		if sum > bPrev {
			b++
		}
		aPrev = v
		bPrev = sum
	}

	fmt.Println("a:", a)
	fmt.Println("b:", b)
}
