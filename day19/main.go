package main

import (
	"bufio"
	"fmt"
	"math"
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

type vector struct {
	x, y, z int
}

func (v *vector) diff(other vector) vector {
	dx := v.x - other.x
	dy := v.y - other.y
	dz := v.z - other.z
	return vector{x: dx, y: dy, z: dz}
}

func main() {
	scanners := [][]vector{}

	var vec []vector
	captureLines("input.txt", func(v string) {
		if strings.HasPrefix(v, "---") {
			vec = []vector{}
			return
		}
		if v == "" {
			scanners = append(scanners, vec)
			return
		}
		coord := strings.Split(v, ",")
		x, err := strconv.Atoi(coord[0])
		if err != nil {
			panic(err)
		}
		y, err := strconv.Atoi(coord[1])
		if err != nil {
			panic(err)
		}
		z, err := strconv.Atoi(coord[2])
		if err != nil {
			panic(err)
		}
		vec = append(vec, vector{x: x, y: y, z: z})
	})

	states := [][][]vector{}
	for _, v := range scanners {
		states = append(states, rotations(v))
	}

	ocean := map[vector]struct{}{}
	for _, v := range scanners[0] {
		ocean[v] = struct{}{}
	}

	seen := map[int]struct{}{}
	seen[0] = struct{}{}

	scannerPos := map[int]vector{}

	for len(seen) != len(scanners) {
	Loop:
		for i, state := range states {
			if _, ok := seen[i]; ok {
				continue
			}
			for rot, points := range state {
				trans := map[vector]int{}
				for _, p := range points {
					for k := range ocean {
						trans[p.diff(k)]++
					}
				}
				for v, hits := range trans {
					if hits >= 12 {
						for _, newp := range states[i][rot] {
							scannerPos[i] = vector{x: -1 * v.x, y: -1 * v.y, z: -1 * v.z}
							ocean[vector{x: -1*v.x + newp.x, y: -1*v.y + newp.y, z: -1*v.z + newp.z}] = struct{}{}
						}
						seen[i] = struct{}{}
						break Loop
					}
				}
			}
		}
	}
	fmt.Println("a:", len(ocean))

	max := math.MinInt32
	for i := 0; i < len(scannerPos)-1; i++ {
		for j := i + 1; j < len(scannerPos); j++ {
			p0 := scannerPos[i]
			p1 := scannerPos[j]
			dist := int(math.Abs(float64(p0.x-p1.x)) + math.Abs(float64(p0.y-p1.y)) + math.Abs(float64(p0.z-p1.z)))
			if dist > max {
				max = dist
			}
		}
	}
	fmt.Println("b:", max)
}

func rotations(coords []vector) [][]vector {
	out := make([][]vector, 24)
	for i := 0; i < 24; i++ {
		out[i] = make([]vector, len(coords))
	}
	for i, v := range coords {
		out[0][i] = vector{v.x, v.y, v.z} // äh...
		out[1][i] = vector{v.x, -v.z, v.y}
		out[2][i] = vector{v.x, -v.y, -v.z}
		out[3][i] = vector{v.x, v.z, -v.y}
		out[4][i] = vector{-v.x, -v.y, v.z}
		out[5][i] = vector{-v.x, -v.z, -v.y}
		out[6][i] = vector{-v.x, v.y, -v.z}
		out[7][i] = vector{-v.x, v.z, v.y}
		out[8][i] = vector{v.y, v.x, -v.z}
		out[9][i] = vector{v.y, -v.x, v.z}
		out[10][i] = vector{v.y, v.z, v.x}
		out[11][i] = vector{v.y, -v.z, -v.x}
		out[12][i] = vector{-v.y, v.x, v.z}
		out[13][i] = vector{-v.y, -v.x, -v.z}
		out[14][i] = vector{-v.y, -v.z, v.x}
		out[15][i] = vector{-v.y, v.z, -v.x}
		out[16][i] = vector{v.z, v.x, v.y}
		out[17][i] = vector{v.z, -v.x, -v.y}
		out[18][i] = vector{v.z, -v.y, v.x}
		out[19][i] = vector{v.z, v.y, -v.x}
		out[20][i] = vector{-v.z, v.x, -v.y}
		out[21][i] = vector{-v.z, -v.x, v.y}
		out[22][i] = vector{-v.z, v.y, v.x}
		out[23][i] = vector{-v.z, -v.y, -v.x}
	}
	return out
}
