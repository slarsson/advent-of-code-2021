package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
	"unicode"
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

func isUpper(str string) bool {
	for _, r := range str {
		return unicode.IsUpper(r)
	}
	return false
}

// :(
func haveDuplicate(arr []string) bool {
	tmp := append([]string{}, arr...)
	sort.Strings(tmp)
	for i := 1; i < len(tmp); i++ {
		if tmp[i] == tmp[i-1] {
			return true
		}
	}
	return false
}

type caves struct {
	list map[string][]string
}

func newCave() caves {
	return caves{
		list: make(map[string][]string),
	}
}

func (c *caves) Add(p1 string, p2 string) {
	c.list[p1] = append(c.list[p1], p2)
	c.list[p2] = append(c.list[p2], p1)
}

func (c *caves) Count(current string, path []string, seen []string, isPartB bool) int {
	if current == "end" {
		return 1
	}

	tot := 0

	next, ok := c.list[current]
	if !ok {
		panic("wtf")
	}

Loop:
	for _, v := range next {
		if v == current || v == "start" {
			continue
		}

		cpy := append([]string{}, seen...)
		if !isUpper(v) {
			for _, s := range seen {
				if v == s {
					if isPartB && !haveDuplicate(seen) {
						break
					}
					continue Loop
				}
			}
			cpy = append(cpy, v)
		}
		tot += c.Count(v, append(path, v), cpy, isPartB)
	}
	return tot
}

func main() {
	h := newCave()

	captureLines("input.txt", func(v string) {
		path := strings.Split(v, "-")
		h.Add(path[0], path[1])
	})

	fmt.Println("a:", h.Count("start", []string{"start"}, []string{}, false))
	fmt.Println("b:", h.Count("start", []string{"start"}, []string{}, true))
}
