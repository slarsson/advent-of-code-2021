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

func inputToIndex(v string) (int, error) {
	switch v {
	case "w":
		return 0, nil
	case "x":
		return 1, nil
	case "y":
		return 2, nil
	case "z":
		return 3, nil
	default:
		return 0, fmt.Errorf("is number")
	}
}

type Op struct {
	Type       string
	LeftIndex  int
	RightIndex *int
	Value      *int
}

func apply(values [4]int, op Op) [4]int {
	var v int
	if op.RightIndex != nil {
		v = values[*op.RightIndex]
	} else if op.Value != nil {
		v = *op.Value
	} else {
		panic("bad input")
	}

	switch op.Type {
	case "add":
		values[op.LeftIndex] += v
	case "mul":
		values[op.LeftIndex] *= v
	case "div":
		if v == 0 {
			panic("wtf1")
		}
		values[op.LeftIndex] /= v
	case "mod":
		if values[op.LeftIndex] < 0 || v <= 0 {
			panic("wtf")
		}
		values[op.LeftIndex] %= v
	case "eql":
		if values[op.LeftIndex] == v {
			values[op.LeftIndex] = 1
		} else {
			values[op.LeftIndex] = 0
		}
	}
	return values
}

type State struct {
	Number [2]int
	Values [4]int
}

func main() {
	ops := []Op{}

	captureLines("input.txt", func(v string) {
		input := strings.Split(v, " ")

		leftIndex, err := inputToIndex(input[1])
		if err != nil {
			panic(err)
		}

		op := Op{
			Type:      input[0],
			LeftIndex: leftIndex,
		}

		if len(input) > 2 {
			rightIndex, err := inputToIndex(input[2])
			if err == nil {
				op.RightIndex = &rightIndex
			} else {
				x, err := strconv.Atoi(input[2])
				if err != nil {
					panic(err)
				}
				op.Value = &x
			}
		}

		ops = append(ops, op)
	})

	states := []State{{
		Number: [2]int{0, 0},
		Values: [4]int{0, 0, 0, 0},
	}}

	for _, op := range ops {
		if op.Type == "inp" {
			fmt.Println("progress:", len(states))

			seen := map[[4]int][2]int{}

			for _, state := range states {
				for i := 1; i <= 9; i++ {
					newMaxValue := state.Number[0]*10 + i
					newMinValue := state.Number[1]*10 + i
					newValues := [4]int{i, state.Values[1], state.Values[2], state.Values[3]}
					if current, ok := seen[newValues]; ok {
						max := current[0]
						min := current[1]
						if newMaxValue > max {
							max = newMaxValue
						}
						if newMinValue < min {
							min = newMinValue
						}
						seen[newValues] = [2]int{max, min}
					} else {
						seen[newValues] = [2]int{newMaxValue, newMinValue}
					}
				}
			}

			states = []State{}
			for k, v := range seen {
				states = append(states, State{
					Number: v,
					Values: k,
				})
			}
			continue
		}
		for i := range states {
			states[i].Values = apply(states[i].Values, op)
		}
	}

	max := math.MinInt
	min := math.MaxInt
	for _, state := range states {
		if state.Values[3] == 0 {
			if state.Number[0] > max {
				max = state.Number[0]
			}
			if state.Number[1] < min {
				min = state.Number[1]
			}
		}
	}

	fmt.Println("a:", max)
	fmt.Println("b:", min)
}
