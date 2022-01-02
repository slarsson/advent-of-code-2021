package main

import (
	"fmt"
	"math"
)

func update(x, y, dx, dy int) (int, int, int, int) {
	x += dx
	y += dy
	dy--
	if dx != 0 {
		if dx < 0 {
			dx++
		} else {
			dx--
		}
	}
	return x, y, dx, dy
}

func main() {
	// target area: x=138..184, y=-125..-71
	targetX := [2]int{138, 184}
	targetY := [2]int{-125, -71}

	maxY := math.MinInt32
	hits := 0

	// just brute force it..
	for x := -1500; x <= 1500; x++ {
		for y := -1500; y <= 1500; y++ {
			dx := x
			dy := y
			curX := 0
			curY := 0
			curMaxY := math.MinInt32

			for {
				curX, curY, dx, dy = update(curX, curY, dx, dy)

				if curY > curMaxY {
					curMaxY = curY
				}

				if dy < 0 && curY < targetY[0] {
					break
				}
				if dx > 0 && curX > targetX[1] {
					break
				}
				if dx < 0 && curX < targetX[0] {
					break
				}

				if curX >= targetX[0] && curX <= targetX[1] && curY >= targetY[0] && curY <= targetY[1] {
					if curMaxY > maxY {
						maxY = curMaxY
					}
					hits++
					break
				}
			}
		}
	}

	fmt.Println("a:", maxY)
	fmt.Println("b:", hits)
}
