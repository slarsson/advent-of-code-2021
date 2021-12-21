package main

import "fmt"

type game struct {
	cache map[item]result
}

type item struct {
	Type    int
	PosP1   int
	PosP2   int
	ScoreP1 int64
	ScoreP2 int64
}

type result struct {
	P1 int
	P2 int
}

func (g *game) play(next int, posP1 int, posP2 int, p1Score int64, p2Score int64) result {
	state := item{
		Type:    next,
		PosP1:   posP1,
		PosP2:   posP2,
		ScoreP1: p1Score,
		ScoreP2: p2Score,
	}
	if v, ok := g.cache[state]; ok {
		return v
	}

	if p1Score >= 21 {
		res := result{
			P1: 1,
			P2: 0,
		}
		g.cache[state] = res
		return res
	} else if p2Score >= 21 {
		res := result{
			P1: 0,
			P2: 1,
		}
		g.cache[state] = res
		return res
	}

	res := result{
		P1: 0,
		P2: 0,
	}

	for x := 1; x < 4; x++ {
		for y := 1; y < 4; y++ {
			for z := 1; z < 4; z++ {
				step := x + y + z
				if next == 1 {
					pos := (posP1+step-1)%10 + 1
					v := g.play(2, pos, posP2, p1Score+int64(pos), p2Score)
					res.P1 += v.P1
					res.P2 += v.P2
				} else if next == 2 {
					pos := (posP2+step-1)%10 + 1
					v := g.play(1, posP1, pos, p1Score, p2Score+int64(pos))
					res.P1 += v.P1
					res.P2 += v.P2
				}
			}
		}
	}

	g.cache[state] = res
	return res
}

func minMax(a, b int) (int, int) {
	if a > b {
		return b, a
	}
	return a, b
}

func main() {
	// part a
	p1 := 8
	p1Score := 0
	p2 := 10
	p2Score := 0

	rolls := 0
	offset := 0
	count := 1
	for {
		step := 0
		offset = offset%100 + 1
		step += offset
		offset = offset%100 + 1
		step += offset
		offset = offset%100 + 1
		step += offset

		rolls += 3

		if count%2 == 0 {
			p2 = (p2+step-1)%10 + 1
			p2Score += p2
			if p2Score >= 1000 {
				break
			}
		} else {
			p1 = (p1+step-1)%10 + 1
			p1Score += p1
			if p1Score >= 1000 {
				break
			}
		}

		count++
	}
	min, _ := minMax(p1Score, p2Score)
	fmt.Println("a:", min*rolls)

	// part b
	g := game{
		cache: map[item]result{},
	}
	out := g.play(1, 8, 10, 0, 0)
	//out := g.exec(1, 4, 8, 0, 0)
	_, max := minMax(out.P1, out.P2)
	fmt.Println("b:", max)
}
