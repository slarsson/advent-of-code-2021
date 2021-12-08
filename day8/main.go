package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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

func score(a []byte, b []byte) int {
	tot := 0
	for _, v1 := range a {
		for _, v2 := range b {
			if v1 == v2 {
				tot++
				break
			}
		}
	}
	return tot
}

func toByte(str string) []byte {
	chars := []byte{}
	for _, v := range str {
		chars = append(chars, byte(v))
	}
	return chars
}

func main() {
	output := [][]string{}
	input := [][]string{}
	captureLines("input.txt", func(v string) {
		line := strings.Split(v, " | ")
		input = append(input, strings.Split(line[0], " "))
		output = append(output, strings.Split(line[1], " "))
	})

	decode := map[int]int{
		2336: 0,
		2335: 3,
		1326: 6,
		2436: 9,
	}

	aScore := 0
	bScore := 0

	for i, v1 := range input {
		known := map[int][]byte{}
		unknown := [][]byte{}

		for _, w := range v1 {
			length := len(w)
			switch length {
			case 2:
				known[1] = toByte(w)
			case 3:
				known[7] = toByte(w)
			case 4:
				known[4] = toByte(w)
			case 7:
				known[8] = toByte(w)
			case 5, 6:
				unknown = append(unknown, toByte(w))
			default:
				panic("wtf")
			}
		}

		is2or5 := [][]byte{}
		for _, v := range unknown {
			sum := score(v, known[1]) * 1000
			sum += score(v, known[4]) * 100
			sum += score(v, known[7]) * 10
			sum += score(v, known[8])
			if k, ok := decode[sum]; ok {
				known[k] = v
			} else {
				is2or5 = append(is2or5, v)
			}
		}

		for _, v := range is2or5 {
			switch score(v, known[4]) {
			case 2:
				known[2] = v
			case 3:
				known[5] = v
			default:
				panic("wtf")
			}
		}

		dict := map[string]int{}
		for k, v := range known {
			sort.Slice(v, func(i int, j int) bool { return v[i] < v[j] })
			dict[string(v)] = k
		}

		tot := 0
		for _, out := range output[i] {
			bytes := toByte(out)
			sort.Slice(bytes, func(i int, j int) bool { return bytes[i] < bytes[j] })
			val := dict[string(bytes)]
			tot *= 10
			tot += val

			if val == 1 || val == 4 || val == 7 || val == 8 {
				aScore++
			}
		}

		bScore += tot
	}

	fmt.Println("a:", aScore)
	fmt.Println("b:", bScore)
}
