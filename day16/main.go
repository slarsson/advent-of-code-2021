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

func main() {
	msg := []byte{}
	captureLines("input.txt", func(v string) {
		for _, r := range v {
			msg = append(msg, hexToBinary(r)...)
		}
	})

	a, b, _ := parse(msg)

	fmt.Println("a:", a)
	fmt.Println("b:", b)
}

func parse(msg []byte) (int, int, []byte) {
	if toDecimal(msg) == 0 {
		return 0, 0, []byte{}
	}

	version, typeID, newMsg := header(msg)
	msg = newMsg
	//fmt.Printf("version=%d, typeID=%d\n", version, typeID)

	if typeID == 4 {
		cursor := 0
		tot := []byte{}

		for {
			tot = append(tot, msg[cursor+1], msg[cursor+2], msg[cursor+3], msg[cursor+4])
			cursor += 5
			if msg[cursor-5] == 0 {
				break
			}
		}

		literal := toDecimal(tot)

		if cursor == len(msg) {
			return version, literal, []byte{}
		}
		return version, literal, msg[cursor:]
	}

	lengthTypeID := msg[0]
	msg = msg[1:]

	totVersion := version
	values := []int{}

	if lengthTypeID == 0 {
		size := toDecimal(msg[0:15])
		msg = msg[15:]
		targetLength := len(msg) - size
		for {
			v, literal, newMsg := parse(msg)
			totVersion += v
			msg = newMsg
			values = append(values, literal)
			if len(msg) == targetLength {
				break
			}
		}
	} else {
		nPackets := toDecimal(msg[0:11])
		msg = msg[11:]
		for i := 0; i < nPackets; i++ {
			v, literal, newMsg := parse(msg)
			totVersion += v
			msg = newMsg
			values = append(values, literal)
		}
	}

	totSum := 0
	switch typeID {
	case 0:
		for _, v := range values {
			totSum += v
		}
	case 1:
		if len(values) == 1 {
			totSum = values[0]
		} else {
			totSum = 1
			for _, v := range values {
				totSum *= v
			}
		}
	case 2:
		totSum = math.MaxInt32
		for _, v := range values {
			if v < totSum {
				totSum = v
			}
		}
	case 3:
		for _, v := range values {
			if v > totSum {
				totSum = v
			}
		}
	case 5:
		if len(values) != 2 {
			panic("not good")
		}
		if values[0] > values[1] {
			totSum = 1
		}
	case 6:
		if len(values) != 2 {
			panic("not good")
		}
		if values[0] < values[1] {
			totSum = 1
		}
	case 7:
		if len(values) != 2 {
			panic("not good")
		}
		if values[0] == values[1] {
			totSum = 1
		}
	default:
		panic("not possible")
	}

	return totVersion, totSum, msg
}

func header(msg []byte) (int, int, []byte) {
	return toDecimal(msg[0:3]), toDecimal(msg[3:6]), msg[6:]
}

func hexToBinary(r rune) []byte {
	switch r {
	case '0':
		return []byte{0, 0, 0, 0}
	case '1':
		return []byte{0, 0, 0, 1}
	case '2':
		return []byte{0, 0, 1, 0}
	case '3':
		return []byte{0, 0, 1, 1}
	case '4':
		return []byte{0, 1, 0, 0}
	case '5':
		return []byte{0, 1, 0, 1}
	case '6':
		return []byte{0, 1, 1, 0}
	case '7':
		return []byte{0, 1, 1, 1}
	case '8':
		return []byte{1, 0, 0, 0}
	case '9':
		return []byte{1, 0, 0, 1}
	case 'A':
		return []byte{1, 0, 1, 0}
	case 'B':
		return []byte{1, 0, 1, 1}
	case 'C':
		return []byte{1, 1, 0, 0}
	case 'D':
		return []byte{1, 1, 0, 1}
	case 'E':
		return []byte{1, 1, 1, 0}
	case 'F':
		return []byte{1, 1, 1, 1}
	default:
		return []byte{0, 0, 0, 0}
	}
}

func toDecimal(arr []byte) int {
	sum := 0
	for i := 0; i < len(arr); i++ {
		if arr[i] == 0 {
			continue
		}
		sum += int(math.Pow(2, float64(len(arr)-i-1)))
	}
	return sum
}
