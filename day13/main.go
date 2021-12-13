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

type coord struct {
	X int
	Y int
}

func main() {
	coords := []coord{}
	folds := []coord{}

	captureLines("input.txt", func(v string) {
		if strings.HasPrefix(v, "fold along x") {
			arr := strings.Split(v, "=")
			x, err := strconv.Atoi(arr[1])
			if err != nil {
				panic(err)
			}
			folds = append(folds, coord{X: x, Y: 0})
			return
		}

		if strings.HasPrefix(v, "fold along y") {
			arr := strings.Split(v, "=")
			y, err := strconv.Atoi(arr[1])
			if err != nil {
				panic(err)
			}
			folds = append(folds, coord{X: 0, Y: y})
			return
		}

		arr := strings.Split(v, ",")
		if len(arr) < 2 {
			return
		}

		x, err := strconv.Atoi(arr[0])
		if err != nil {
			panic(err)
		}

		y, err := strconv.Atoi(arr[1])
		if err != nil {
			panic(err)
		}

		coords = append(coords, coord{X: x, Y: y})
	})

	partA := true
	for _, fold := range folds {
		newCoords := map[coord]struct{}{}
		if fold.Y != 0 {
			for _, c := range coords {
				if c.Y < fold.Y {

					newCoords[c] = struct{}{}
					continue
				}
				if c.Y == fold.Y {
					continue
				}
				newCoords[coord{X: c.X, Y: fold.Y - (c.Y - fold.Y)}] = struct{}{}
			}
		} else if fold.X != 0 {
			for _, c := range coords {
				if c.X < fold.X {
					newCoords[c] = struct{}{}
					continue
				}
				newCoords[coord{X: fold.X - (c.X - fold.X), Y: c.Y}] = struct{}{}
			}
		}

		if partA {
			fmt.Println("a:", len(newCoords))
			partA = false
		}

		arr := []coord{}
		for k := range newCoords {
			arr = append(arr, k)
		}
		coords = arr
	}

	w := 0
	h := 0
	for _, v := range coords {
		if v.X > w {
			w = v.X
		}
		if v.Y > h {
			h = v.Y
		}
	}

	for i := 0; i <= h; i++ {
	Loop:
		for j := 0; j <= w; j++ {
			for _, coord := range coords {
				if coord.X == j && coord.Y == i {
					fmt.Print("#")
					continue Loop
				}
			}
			fmt.Print(".")
		}
		fmt.Println("")
	}
}
