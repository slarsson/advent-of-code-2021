package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
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

func getRune(str string, pos int) rune {
	var r rune
	for i, v := range str {
		if i == pos {
			return v
		}
	}
	return r
}

func runes(str string) []rune {
	rs := []rune{}
	for _, r := range str {
		rs = append(rs, r)
	}
	return rs
}

func minMax(dist map[rune]int64) (int64, int64) {
	max := int64(math.MinInt64)
	min := int64(math.MaxInt64)
	for _, v := range dist {
		if v > max {
			max = v
		}
		if v < min {
			min = v
		}
	}
	return min, max
}

type pair struct {
	First  rune
	Second rune
}

func main() {
	r := []rune{}
	rules := map[pair]rune{}
	captureLines("input.txt", func(v string) {
		if len(r) == 0 {
			r = runes(v)
			return
		}
		if v == "" {
			return
		}
		rule := strings.Split(v, " -> ")
		rules[pair{First: getRune(rule[0], 0), Second: getRune(rule[0], 1)}] = getRune(rule[1], 0)
	})

	for _, itrs := range []int{10, 40} {
		current := map[pair]int64{}
		for i := 0; i < len(r)-1; i++ {
			current[pair{First: r[i], Second: r[i+1]}]++
		}

		for i := 0; i < itrs; i++ {
			x := map[pair]int64{}
			for k, v := range current {
				x[k] = v
			}
			for k, v := range current {
				r, ok := rules[k]
				if !ok {
					continue
				}
				x[pair{First: k.First, Second: r}] += v
				x[pair{First: r, Second: k.Second}] += v
				x[k] -= v
			}
			current = x
		}

		dist := map[rune]int64{}
		for k, v := range current {
			dist[k.First] += v
			dist[k.Second] += v
		}
		dist[r[len(r)-1]]++
		min, max := minMax(dist)
		fmt.Println(max/2 - min/2) // meh..
	}
}
