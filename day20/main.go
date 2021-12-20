package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
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

	pixels := map[coord]struct{}{}
	lookup := map[int]byte{}

	isImage := false
	row := 0
	line := ""
	captureLines("input.txt", func(v string) {
		if v == "" {
			isImage = true
			return
		}
		if isImage {
			for i, r := range v {
				if r == '.' {
					continue
				}
				pixels[coord{X: i, Y: row}] = struct{}{}
			}
			row++
			return
		}
		line += v
	})

	for i, r := range line {
		var x byte
		if r == '#' {
			x = 1
		} else {
			x = 0
		}
		lookup[i] = x
	}

	for i := 0; i < 50; i++ {
		startX := math.MaxInt32
		startY := math.MaxInt32
		endX := math.MinInt32
		endY := math.MinInt32

		for p := range pixels {
			if p.X < startX {
				startX = p.X
			}
			if p.X > endX {
				endX = p.X
			}
			if p.Y < startY {
				startY = p.Y
			}
			if p.Y > endY {
				endY = p.Y
			}
		}

		inf := 0
		// meh..
		if lookup[0] == 1 && lookup[511] == 0 {
			inf = i % 2
		}

		output := map[coord]struct{}{}

		for x := startX - 1; x <= endX+1; x++ {
			for y := startY - 1; y <= endY+1; y++ {
				sum := 0
				for ii, v := range []coord{
					{X: x - 1, Y: y - 1},
					{X: x, Y: y - 1},
					{X: x + 1, Y: y - 1},
					{X: x - 1, Y: y},
					{X: x, Y: y},
					{X: x + 1, Y: y},
					{X: x - 1, Y: y + 1},
					{X: x, Y: y + 1},
					{X: x + 1, Y: y + 1},
				} {
					isZero := false
					if v.X < startX || v.X > endX || v.Y < startY || v.Y > endY {
						isZero = inf == 0
					} else {
						_, ok := pixels[v]
						isZero = !ok
					}
					if !isZero {
						sum += int(math.Pow(2, float64(8-ii)))
					}
				}
				if v, ok := lookup[sum]; ok && v == 1 {
					output[coord{X: x, Y: y}] = struct{}{}
				}
			}
		}

		pixels = output

		if i == 1 {
			fmt.Println("a:", len(pixels))
		} else if i == 49 {
			fmt.Println("b:", len(pixels))
		}
	}
}
