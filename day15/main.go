package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
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

type Position struct {
	X int
	Y int
}

type Edge struct {
	Node   Position
	Weight int
}

type Vertex struct {
	Prev  *Position
	Score int
}

func main() {
	pos := map[Position]int{}
	row := 0
	size := 100
	captureLines("input.txt", func(v string) {
		for col, x := range v {
			val := int(x - '0')
			for i := 0; i < 5; i++ {
				for j := 0; j < 5; j++ {
					next := val + i + j
					if next > 9 {
						next -= 9
					}
					pos[Position{X: col + i*size, Y: row + j*size}] = next
				}
			}
		}
		row++
	})

	edges := map[Position][]Edge{}
	vertices := map[Position]Vertex{}

	for from := range pos {
		vertices[from] = Vertex{Prev: nil, Score: math.MaxInt32}
		edges[from] = []Edge{}
		if weight, ok := pos[Position{X: from.X - 1, Y: from.Y}]; ok {
			edges[from] = append(edges[from], Edge{
				Node:   Position{X: from.X - 1, Y: from.Y},
				Weight: weight,
			})
		}
		if weight, ok := pos[Position{X: from.X + 1, Y: from.Y}]; ok {
			edges[from] = append(edges[from], Edge{
				Node:   Position{X: from.X + 1, Y: from.Y},
				Weight: weight,
			})
		}
		if weight, ok := pos[Position{X: from.X, Y: from.Y - 1}]; ok {
			edges[from] = append(edges[from], Edge{
				Node:   Position{X: from.X, Y: from.Y - 1},
				Weight: weight,
			})
		}
		if weight, ok := pos[Position{X: from.X, Y: from.Y + 1}]; ok {
			edges[from] = append(edges[from], Edge{
				Node:   Position{X: from.X, Y: from.Y + 1},
				Weight: weight,
			})
		}
	}

	startPos := Position{X: 0, Y: 0}

	queue := []Position{startPos}
	vertices[startPos] = Vertex{
		Prev:  nil,
		Score: 0,
	}

	seen := map[Position]struct{}{}

	for {
		if len(queue) == 0 {
			break
		}

		sort.Slice(queue, func(i, j int) bool {
			return vertices[queue[i]].Score < vertices[queue[j]].Score
		})
		cur := queue[0]
		queue = queue[1:]

		if _, ok := seen[cur]; ok {
			break
		}
		seen[cur] = struct{}{}

		curScore := vertices[cur].Score
		for _, edge := range edges[cur] {
			tot := curScore + edge.Weight
			target := vertices[edge.Node]
			if tot < target.Score {
				vertices[edge.Node] = Vertex{Prev: &cur, Score: tot}
				queue = append(queue, Position{X: edge.Node.X, Y: edge.Node.Y})
			}
		}
	}

	fmt.Println("a:", vertices[Position{X: 99, Y: 99}].Score)
	fmt.Println("b:", vertices[Position{X: 499, Y: 499}].Score)
}
