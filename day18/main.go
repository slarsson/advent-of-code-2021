package main

import (
	"bufio"
	"fmt"
	"math"
	"math/rand"
	"os"
	"time"
)

type Node struct {
	Parent *Node
	First  *Node
	Second *Node
	Value  int
	Hash   string
}

func main() {
	rand.Seed(time.Now().UnixNano())

	nodes := []*Node{}

	captureLines("input.txt", func(v string) {
		root := Node{
			Parent: nil,
			First:  nil,
			Second: nil,
			Value:  -1,
		}
		nodes = append(nodes, &root)
		cur := &root
		for _, r := range v[1 : len(v)-1] {
			switch r {
			case '[':
				n := &Node{
					Parent: cur,
					First:  nil,
					Second: nil,
					Value:  -1,
				}
				if cur.First == nil {
					cur.First = n
				} else {
					cur.Second = n
				}
				cur = n
			case ']':
				cur = cur.Parent
			case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
				valueNode := &Node{
					Parent: cur,
					First:  nil,
					Second: nil,
					Value:  int(r - '0'),
				}
				if cur.First == nil {
					cur.First = valueNode
				} else {
					cur.Second = valueNode
				}
			}
		}
	})

	prevSum := clone(nodes[0])
	for _, cur := range nodes[1:] {
		node := &Node{First: prevSum, Second: clone(cur), Value: -1}
		node.First.Parent = node
		node.Second.Parent = node
		reduce(node)
		prevSum = node
	}
	fmt.Println("a:", magnitude(prevSum))

	max := math.MinInt32
	for i := 0; i < len(nodes); i++ {
		for j := 0; j < len(nodes); j++ {
			if i == j {
				continue
			}
			node := &Node{First: clone(nodes[i]), Second: clone(nodes[j]), Value: -1}
			node.First.Parent = node
			node.Second.Parent = node
			reduce(node)
			if sum := magnitude(node); sum > max {
				max = sum
			}
		}
	}
	fmt.Println("b:", max)
}

func clone(node *Node) *Node {
	if node == nil {
		return nil
	}
	cur := &Node{
		Parent: nil,
		First:  nil,
		Second: nil,
		Value:  node.Value,
		Hash:   node.Hash,
	}
	first := clone(node.First)
	if first != nil {
		cur.First = first
		first.Parent = cur
	}
	second := clone(node.Second)
	if second != nil {
		cur.Second = second
		second.Parent = cur
	}
	return cur
}

func magnitude(node *Node) int {
	if node.Value != -1 {
		return node.Value
	}
	sum := 0
	sum += 3 * magnitude(node.First)
	sum += 2 * magnitude(node.Second)
	return sum
}

func reduce(node *Node) {
	for {
		v1, v2 := getNodeList(node)
		opz := op{
			done:   false,
			lookup: v1,
			order:  v2,
		}
		opz.explode(node, 0)
		if opz.done {
			continue
		}
		opz.split(node)
		if opz.done {
			continue
		}
		break
	}
}

type op struct {
	done   bool
	order  []string
	lookup map[string]*Node
}

func (o *op) before(node *Node) *Node {
	for i, v := range o.order {
		if v == node.Hash {
			if i-1 < 0 {
				return nil
			}
			return o.lookup[o.order[i-1]]
		}
	}
	return nil
}

func (o *op) after(node *Node) *Node {
	for i, v := range o.order {
		if v == node.Hash {
			if i+1 >= len(o.order) {
				return nil
			}
			return o.lookup[o.order[i+1]]
		}
	}
	return nil
}

func (o *op) explode(node *Node, depthSoFar int) {
	if node == nil || o.done {
		return
	}

	depthSoFar++
	if depthSoFar == 5 {
		if node.First == nil || node.Second == nil {
			return
		}

		if left := o.before(node.First); left != nil {
			left.Value += node.First.Value
		}

		if right := o.after(node.Second); right != nil {
			right.Value += node.Second.Value
		}

		node.First = nil
		node.Second = nil
		node.Value = 0

		o.done = true

		return
	}
	o.explode(node.First, depthSoFar)
	o.explode(node.Second, depthSoFar)
}

func (o *op) split(node *Node) {
	if o.done || node == nil {
		return
	}

	if node.First == nil && node.Second == nil && node.Value >= 10 {
		left := node.Value / 2
		right := node.Value/2 + node.Value%2

		first := &Node{Parent: node, First: nil, Second: nil, Value: left}
		second := &Node{Parent: node, First: nil, Second: nil, Value: right}

		node.Value = -1
		node.First = first
		node.Second = second

		o.done = true

		return
	}
	o.split(node.First)
	o.split(node.Second)
}

func order(node *Node, list map[string]*Node) []string {
	if node.Hash == "" {
		node.Hash = fmt.Sprintf("%d", rand.Intn(1000000000))
	}
	list[node.Hash] = node
	if node.Value != -1 {
		return []string{node.Hash}
	}
	out := []string{}
	out = append(out, order(node.First, list)...)
	out = append(out, order(node.Second, list)...)
	return out
}

func getNodeList(node *Node) (map[string]*Node, []string) {
	lookup := map[string]*Node{}
	list := order(node, lookup)
	return lookup, list
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
