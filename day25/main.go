package main

import (
	"bufio"
	"fmt"
	"os"
)

type coord struct {
	X int
	Y int
}

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

	coords := map[coord]string{}

	w := 139
	h := 137

	row := 0
	captureLines("input.txt", func(v string) {
		for i, r := range v {
			coords[coord{
				X: i,
				Y: row,
			}] = string(r)
		}
		row++
	})

	moves := 0
	count := 0

	for {
		start := moves
		newCoords := map[coord]string{}

		for i := 0; i < h; i++ {
			for j := 0; j < w; j++ {
				if coords[coord{X: j, Y: i}] != ">" {
					continue
				}
				if v, ok := coords[coord{X: j + 1, Y: i}]; ok {
					if v == "." {
						newCoords[coord{X: j, Y: i}] = "."
						newCoords[coord{X: j + 1, Y: i}] = ">"
						moves++
					}
				} else if v, ok := coords[coord{X: 0, Y: i}]; ok {
					if v == "." {
						newCoords[coord{X: j, Y: i}] = "."
						newCoords[coord{X: 0, Y: i}] = ">"
						moves++
					}
				}
			}
		}

		for k, v := range newCoords {
			coords[k] = v
		}

		for i := 0; i < h; i++ {
			for j := 0; j < w; j++ {
				if coords[coord{X: j, Y: i}] != "v" {
					continue
				}
				if v, ok := coords[coord{X: j, Y: i + 1}]; ok {
					if v == "." {
						newCoords[coord{X: j, Y: i}] = "."
						newCoords[coord{X: j, Y: i + 1}] = "v"
						moves++
					}
				} else if v, ok := coords[coord{X: j, Y: 0}]; ok {
					if v == "." {
						newCoords[coord{X: j, Y: i}] = "."
						newCoords[coord{X: j, Y: 0}] = "v"
						moves++
					}
				}
			}
		}

		for k, v := range newCoords {
			coords[k] = v
		}

		count++
		if start == moves {
			break
		}
	}

	fmt.Println("a", count)
}
