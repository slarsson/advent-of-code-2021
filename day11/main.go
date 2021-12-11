package main

import (
	"fmt"
	"math"
)

type stack struct {
	arr []int
}

func (s *stack) Push(v int) {
	s.arr = append(s.arr, v)
}

func (s *stack) Pop() (int, error) {
	if len(s.arr) == 0 {
		return 0, fmt.Errorf("empty :(")
	}
	v := s.arr[len(s.arr)-1]
	s.arr = s.arr[:len(s.arr)-1]
	return v, nil
}

type grid struct {
	matrix []int
	size   int
}

func newGrid(list string) grid {
	size := int(math.Sqrt(float64(len(list))))

	padding := []int{}
	for i := 0; i < size+2; i++ {
		padding = append(padding, math.MinInt32)
	}

	matrix := []int{}
	matrix = append(matrix, padding...)
	for i, r := range list {
		if i%size == 0 {
			matrix = append(matrix, math.MinInt32)
		}
		matrix = append(matrix, int(r-'0'))
		if i%size == (size - 1) {
			matrix = append(matrix, math.MinInt32)
		}
	}
	matrix = append(matrix, padding...)

	return grid{
		size:   size,
		matrix: matrix,
	}
}

func (g *grid) Flash() int {
	seen := map[int]struct{}{}
	s := stack{}

	for i := g.size + 3; i < len(g.matrix)-(g.size+3); i++ {
		g.matrix[i]++
		if g.matrix[i] <= 9 {
			continue
		}
		s.Push(i)
		seen[i] = struct{}{}
	}

	for {
		i, err := s.Pop()
		if err != nil {
			break
		}

		adjecent := []int{
			i - g.size - 2 - 1,
			i - g.size - 2,
			i - g.size - 2 + 1,
			i - 1,
			i + 1,
			i + g.size + 2 - 1,
			i + g.size + 2,
			i + g.size + 2 + 1,
		}
		for _, v := range adjecent {
			g.matrix[v]++
			if _, ok := seen[v]; ok {
				continue
			}
			if g.matrix[v] > 9 {
				s.Push(v)
				seen[v] = struct{}{}
			}
		}
	}

	flashes := 0
	for i := g.size + 3; i < len(g.matrix)-(g.size+3); i++ {
		if g.matrix[i] > 9 {
			g.matrix[i] = 0
			flashes++
		} else if g.matrix[i] < 0 {
			g.matrix[i] = math.MinInt32
		}
	}
	return flashes
}

func main() {
	input := "3172537688456648312563745126538321148885434274775813621885827582213132688787526876351127877242787273"
	//input := "5483143223274585471152645561736141336146635738547841675246452176841721688288113448468485545283751526"

	aGrid := newGrid(input)
	tot := 0
	for i := 0; i < 100; i++ {
		tot += aGrid.Flash()
	}
	fmt.Println("a:", tot)

	bGrid := newGrid(input)
	i := 0
	for {
		if bGrid.Flash() == len(input) {
			fmt.Println("b:", i+1)
			break
		}
		i++
	}
}
