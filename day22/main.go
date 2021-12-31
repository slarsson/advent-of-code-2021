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

type action struct {
	state string
	x     [2]int
	y     [2]int
	z     [2]int
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

func getIndex(value int, arr []int) int {
	for i, v := range arr {
		if v == value {
			return i
		}
	}
	panic("not possible")
}

func removeDups(cur []int) []int {
	arr := []int{}
	seen := map[int]struct{}{}
	for _, v := range cur {
		if _, ok := seen[v]; ok {
			continue
		}
		seen[v] = struct{}{}
		arr = append(arr, v)
	}
	return arr
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	parts := []struct {
		Part string
		Max  int
		Min  int
	}{
		{
			Part: "a",
			Max:  50,
			Min:  -50,
		},
		{
			Part: "b",
			Max:  math.MaxInt,
			Min:  math.MinInt,
		},
	}

	for _, part := range parts {
		actions := []action{}
		xs := []int{}
		ys := []int{}
		zs := []int{}

		captureLines("input.txt", func(v string) {
			a := action{}
			tmp := strings.Fields(v)
			a.state = tmp[0]
			for i, v := range strings.Split(tmp[1], ",") {
				r := strings.Split(v[2:], "..")
				start, err := strconv.Atoi(r[0])
				if err != nil {
					panic(err)
				}
				start = max(start, part.Min)
				end, err := strconv.Atoi(r[1])
				if err != nil {
					panic(err)
				}
				end = min(end, part.Max)
				end++ // +1 to handle inclusive intervals
				switch i {
				case 0:
					a.x = [2]int{start, end}
					xs = append(xs, start, end)
				case 1:
					a.y = [2]int{start, end}
					ys = append(ys, start, end)
				case 2:
					a.z = [2]int{start, end}
					zs = append(zs, start, end)
				}
			}
			actions = append(actions, a)
		})

		sort.Ints(xs)
		sort.Ints(ys)
		sort.Ints(zs)

		xs = removeDups(xs)
		ys = removeDups(ys)
		zs = removeDups(zs)

		fmt.Printf("memory: %f GB\n", float64(len(xs)*len(ys)*len(zs)*4)/1000000000)

		grid := [][][]bool{}
		for i := 0; i < len(xs); i++ {
			grid = append(grid, [][]bool{})
			for j := 0; j < len(ys); j++ {
				grid[i] = append(grid[i], []bool{})
				for k := 0; k < len(zs); k++ {
					grid[i][j] = append(grid[i][j], false)
				}
			}
		}

		for _, a := range actions {
			x0 := getIndex(a.x[0], xs)
			x1 := getIndex(a.x[1], xs)
			y0 := getIndex(a.y[0], ys)
			y1 := getIndex(a.y[1], ys)
			z0 := getIndex(a.z[0], zs)
			z1 := getIndex(a.z[1], zs)

			on := a.state == "on"

			for i := x0; i < x1; i++ {
				for j := y0; j < y1; j++ {
					for k := z0; k < z1; k++ {
						grid[i][j][k] = on
					}
				}
			}
		}

		sum := int64(0)
		for i := 0; i < len(xs)-1; i++ {
			for j := 0; j < len(ys)-1; j++ {
				for k := 0; k < len(zs)-1; k++ {
					if !grid[i][j][k] {
						continue
					}
					sum += int64(xs[i+1]-xs[i]) * int64(ys[j+1]-ys[j]) * int64(zs[k+1]-zs[k])
				}
			}
		}

		fmt.Printf("%s: %d\n", part.Part, sum)
	}
}
