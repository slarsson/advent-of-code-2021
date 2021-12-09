package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
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

func walk(p int, w int, seen map[int]struct{}, matrix []int) int {
	if _, ok := seen[p]; ok {
		return 0
	}
	seen[p] = struct{}{}

	tot := 1
	if matrix[p+1] < 9 {
		tot += walk(p+1, w, seen, matrix)
	}
	if matrix[p-1] < 9 {
		tot += walk(p-1, w, seen, matrix)
	}
	if matrix[p+w] < 9 {
		tot += walk(p+w, w, seen, matrix)
	}
	if matrix[p-w] < 9 {
		tot += walk(p-w, w, seen, matrix)
	}
	return tot
}

func main() {
	matrix := []int{}
	w := 100

	padding := []int{}
	for i := 0; i < w+2; i++ {
		padding = append(padding, 9)
	}

	matrix = append(matrix, padding...)
	captureLines("input.txt", func(v string) {
		matrix = append(matrix, 9)
		for _, r := range v {
			v, err := strconv.Atoi(string(r))
			if err != nil {
				panic(err)
			}
			matrix = append(matrix, v)
		}
		matrix = append(matrix, 9)
	})
	matrix = append(matrix, padding...)

	lowPointIndex := []int{}
	a := 0
	for i := w + 3; i < len(matrix)-(w+3); i++ {
		if matrix[i] < matrix[i-1] && matrix[i] < matrix[i+1] && matrix[i] < matrix[i-w-2] && matrix[i] < matrix[i+w+2] {
			a += matrix[i] + 1
			lowPointIndex = append(lowPointIndex, i)
		}
	}
	fmt.Println("a:", a)

	max := []int{0, 0, 0}
	for _, p := range lowPointIndex {
		size := walk(p, w+2, map[int]struct{}{}, matrix)
		if size > max[0] {
			max[0] = size
			sort.Ints(max)
		}
	}
	fmt.Println("b:", max[0]*max[1]*max[2])
}
