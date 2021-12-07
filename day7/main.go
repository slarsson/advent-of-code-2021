package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
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

func max(x, y int64) int64 {
	if x < y {
		return y
	}
	return x
}

func main() {
	input := []int64{}

	captureLines("input.txt", func(v string) {
		for _, v := range strings.Split(v, ",") {
			n, err := strconv.Atoi(v)
			if err != nil {
				panic(err)
			}
			input = append(input, int64(n))
		}
	})

	sort.Slice(input, func(i, j int) bool {
		return input[i] < input[j]
	})

	aLowest := int64(math.MaxInt64)
	bLowest := int64(math.MaxInt64)

	for i := input[0]; i < input[len(input)-1]; i++ {
		aScore := int64(0)
		bScore := int64(0)
		for j := 0; j < len(input); j++ {
			x := max(i-input[j], input[j]-i)
			aScore += x
			bScore += x * (x + 1) / 2 // sum of 1+2+3+..+n => n*(n+1)/2
		}
		if aScore < aLowest {
			aLowest = aScore
		}
		if bScore < bLowest {
			bLowest = bScore
		}
	}

	fmt.Println("a:", aLowest)
	fmt.Println("b:", bLowest)
}
