package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

type Stack struct {
	arr []byte
}

func (s *Stack) Push(v byte) {
	s.arr = append(s.arr, v)
}

func (s *Stack) Pop() (byte, error) {
	if len(s.arr) == 0 {
		return 0, fmt.Errorf("empty :(")
	}
	v := s.arr[len(s.arr)-1]
	s.arr = s.arr[:len(s.arr)-1]
	return v, nil
}

func (s *Stack) Len() int {
	return len(s.arr)
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
	lines := []string{}
	captureLines("input.txt", func(v string) {
		lines = append(lines, v)
	})

	tokens := map[byte]byte{
		40:  41, // ()
		41:  40,
		123: 125, // {}
		125: 123,
		91:  93, // []
		93:  91,
		60:  62, // <>
		62:  60,
	}

	bad := []byte{}
	endings := []string{}
Main:
	for _, line := range lines {
		s := Stack{}
		for _, v := range line {
			b := byte(v)
			if b == 40 || b == 123 || b == 91 || b == 60 {
				s.Push(b)
			} else {
				v, _ := s.Pop()
				if end := tokens[v]; end != b {
					bad = append(bad, b)
					//fmt.Println("got:", string(end))
					continue Main
				}
			}
		}
		if s.Len() != 0 {
			ending := ""
			for {
				v, err := s.Pop()
				if err != nil {
					break
				}
				ending += string(tokens[v])
			}
			endings = append(endings, ending)
		}
	}

	aScore := 0
	for _, b := range bad {
		switch b {
		case 41:
			aScore += 3
		case 93:
			aScore += 57
		case 125:
			aScore += 1197
		case 62:
			aScore += 25137
		}
	}
	fmt.Println("a:", aScore)

	bScores := []int{}
	for _, w := range endings {
		tot := 0
		for _, r := range w {
			tot *= 5
			b := byte(r)
			switch b {
			case 41:
				tot += 1
			case 93:
				tot += 2
			case 125:
				tot += 3
			case 62:
				tot += 4
			}
		}
		bScores = append(bScores, tot)
	}
	sort.Ints(bScores)
	fmt.Println("b:", bScores[len(bScores)/2])
}
