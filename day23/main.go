package main

import (
	"container/heap"
	"fmt"
	"math"
	"sort"
	"strings"
)

type coord struct {
	X int
	Y int
}

func (c *coord) IsHallway() bool {
	return c.Y == 0
}

func (c *coord) RoomToHallway(grid map[coord]int) (int, bool) {
	for i := c.Y - 1; i > 0; i-- {
		if _, ok := grid[coord{X: c.X, Y: i}]; ok {
			return 0, false
		}
	}
	return c.Y, true
}

func (c *coord) HallwayToRoom(roomSize int, typ int, grid map[coord]int) (int, bool) {
	depth := roomSize
	for i := roomSize; i > 0; i-- {
		if gridTyp, ok := grid[coord{X: c.X, Y: i}]; ok {
			if gridTyp != typ {
				return 0, false
			}
			depth--
		}
	}
	return depth, true
}

func (c *coord) Hallway(roomSize int, target int, grid map[coord]int) ([]int, int) {
	// 0, 1, 3, 5, 7, 9, 10
	typ := grid[*c]
	possible := []int{}
	checkRoom := false

	// left
	for i := c.X - 1; i >= 0; i-- {
		if _, ok := grid[coord{X: i, Y: 0}]; ok {
			break
		}
		if i == target {
			checkRoom = true
		} else if i == 0 || i == 1 || i == 3 || i == 5 || i == 7 {
			possible = append(possible, i)
		}
	}

	// right
	for i := c.X + 1; i <= 10; i++ {
		if _, ok := grid[coord{X: i, Y: 0}]; ok {
			break
		}
		if i == target {
			checkRoom = true
		} else if i == 3 || i == 5 || i == 7 || i == 9 || i == 10 {
			possible = append(possible, i)
		}
	}

	room := -1
	if checkRoom {
		roomEntry := coord{X: target, Y: 0}
		depth, ok := roomEntry.HallwayToRoom(roomSize, typ, grid)
		if ok {
			room = depth
		}
	}

	return possible, room
}

type state struct {
	Grid map[coord]int
	Cost int
}

func (s *state) IsDone() bool {
	return s.Grid[coord{X: 2, Y: 1}] == 1 &&
		s.Grid[coord{X: 2, Y: 2}] == 1 &&
		s.Grid[coord{X: 4, Y: 1}] == 10 &&
		s.Grid[coord{X: 4, Y: 2}] == 10 &&
		s.Grid[coord{X: 6, Y: 1}] == 100 &&
		s.Grid[coord{X: 6, Y: 2}] == 100 &&
		s.Grid[coord{X: 8, Y: 1}] == 1000 &&
		s.Grid[coord{X: 8, Y: 2}] == 1000
}

func clone(pos coord, grid map[coord]int) map[coord]int {
	newMap := map[coord]int{}
	for k, v := range grid {
		if k == pos {
			continue
		}
		newMap[k] = v
	}
	return newMap
}

func hash(grid map[coord]int) string {
	out := []string{}
	for k, v := range grid {
		out = append(out, fmt.Sprintf("%d%d%d", k.X, k.Y, v))
	}
	sort.Strings(out)
	return strings.Join(out, "#")
}

// meh..
func moves(roomSize int, costSoFar int, current map[coord]int) []state {
	newStates := []state{}

	for k, v := range current {
		var target int
		switch v {
		case 1:
			target = 2
		case 10:
			target = 4
		case 100:
			target = 6
		case 1000:
			target = 8
		default:
			panic("not possible")
		}
		possible, room := k.Hallway(roomSize, target, current)

		if k.IsHallway() {
			if room != -1 {
				dist := int(math.Abs(float64(k.X) - float64(target)))
				newState := clone(k, current)
				newState[coord{X: target, Y: room}] = v
				newStates = append(newStates, state{
					Grid: newState,
					Cost: costSoFar + (dist+room)*v,
				})
			}
			continue
		}

		roomToHallwayDist, ok := k.RoomToHallway(current)
		if !ok {
			continue
		}

		if room != -1 {
			dist := int(math.Abs(float64(k.X) - float64(target)))
			newState := clone(k, current)
			newState[coord{X: target, Y: room}] = v
			newStates = append(newStates, state{
				Grid: newState,
				Cost: costSoFar + (roomToHallwayDist+dist+room)*v,
			})
		}

		for _, newPos := range possible {
			dist := int(math.Abs(float64(k.X) - float64(newPos)))
			newState := clone(k, current)
			newState[coord{X: newPos, Y: 0}] = v
			newStates = append(newStates, state{
				Grid: newState,
				Cost: costSoFar + (roomToHallwayDist+dist)*v,
			})
		}
	}

	return newStates
}

func main() {
	inputs := []struct {
		Part string
		Size int
		Init map[coord]int
	}{
		{
			Part: "a",
			Size: 2,
			Init: map[coord]int{
				{X: 8, Y: 1}: 1,
				{X: 8, Y: 2}: 1,
				{X: 4, Y: 1}: 10,
				{X: 6, Y: 1}: 10,
				{X: 2, Y: 2}: 100,
				{X: 4, Y: 2}: 100,
				{X: 2, Y: 1}: 1000,
				{X: 6, Y: 2}: 1000,
			},
		},
		{
			Part: "b",
			Size: 4,
			Init: map[coord]int{
				{X: 8, Y: 1}: 1,
				{X: 8, Y: 2}: 1,
				{X: 6, Y: 3}: 1,
				{X: 8, Y: 4}: 1,
				{X: 4, Y: 1}: 10,
				{X: 6, Y: 1}: 10,
				{X: 6, Y: 2}: 10,
				{X: 4, Y: 3}: 10,
				{X: 2, Y: 4}: 100,
				{X: 4, Y: 2}: 100,
				{X: 4, Y: 4}: 100,
				{X: 8, Y: 3}: 100,
				{X: 2, Y: 1}: 1000,
				{X: 2, Y: 2}: 1000,
				{X: 2, Y: 3}: 1000,
				{X: 6, Y: 4}: 1000,
			},
		},
	}

	for _, input := range inputs {
		seen := map[string]int{}
		pq := make(PriorityQueue, 1)
		pq[0] = state{
			Grid: input.Init,
			Cost: 0,
		}

		for {
			cur := heap.Pop(&pq).(state)

			id := hash(cur.Grid)
			if _, ok := seen[id]; ok {
				continue
			}
			seen[id] = cur.Cost

			if cur.IsDone() {
				fmt.Println(fmt.Sprintf("%s: %d", input.Part, cur.Cost))
				break
			}

			for _, item := range moves(input.Size, cur.Cost, cur.Grid) {
				heap.Push(&pq, item)
			}
		}

	}
}
