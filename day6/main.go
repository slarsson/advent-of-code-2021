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
	occr := make(map[int]uint64)
	for i := 0; i <= 9; i++ {
		occr[i] = 0
	}

	captureLines("input.txt", func(v string) {
		for _, v := range strings.Split(v, ",") {
			i, err := strconv.Atoi(v)
			if err != nil {
				panic(err)
			}
			occr[i]++
		}
	})

	for i := 0; i < 256; i++ {
		fishes := occr[0]
		for i := 1; i < 9; i++ {
			occr[i-1] = occr[i]
		}
		occr[6] += fishes
		occr[8] = fishes

		if i == 79 {
			fmt.Println("a:", sum(occr))
		}
	}
	fmt.Println("b:", sum(occr))
}

func sum(m map[int]uint64) uint64 {
	tot := uint64(0)
	for _, v := range m {
		tot += v
	}
	return tot
}
