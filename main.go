package main

import (
	"bufio"
	"fmt"
	"os"
)

/**
 * Auto-generated code below aims at helping you parse
 * the standard input according to the problem statement.
 **/

type Map [][]int
type Pos struct {
	x int
	y int
}
type V []Pos

func (m Map) set(p Pos, v int) {
	m[p.y][p.x] = v
}

func (m Map) get(p Pos) int {
	return m[p.y][p.x]
}

func isInLine(v int, p Pos, m Map) bool {
	for i := 0; i < 9; i++ {
		if (m[p.y][i] == v && p.x != i) || (m[i][p.x] == v && p.y != i) {
			return true
		}
	}
	return false
}

func findSquare(i int) (int, int) {
	var min int
	var max int

	if i <= 2 {
		max = 3
		min = 0
	} else if i >= 6 {
		max = 9
		min = 6
	} else {
		min = 3
		max = 6
	}
	return min, max
}

func isInSquare(v int, p Pos, m Map) bool {
	var min Pos
	var max Pos

	min.x, max.x = findSquare(p.x)
	min.y, max.y = findSquare(p.y)
	for y := min.y; y != max.y; y++ {
		for x := min.x; x != max.x; x++ {
			if m[y][x] == v {
				return true
			}
		}
	}

	return false
}

func check(e error) {
	if e != nil {
			panic(e)
	}
}

func getMap() (m Map) {
	f, err := os.Open(os.Args[1])
	check(err)
	scanner := bufio.NewScanner(f)
	scanner.Buffer(make([]byte, 10), 10)
	m = make([][]int, 9)
	for i := 0; i < 9; i++ {
		m[i] = make([]int, 9)
		scanner.Scan()
		line := scanner.Text()
		_ = line // to avoid unused error

		for j := 0; j < 9; j++ {
			m[i][j] = int(line[j]) - 48
		}
	}

	return
}

func resolv(m Map, allZero V, i int) bool {
	if i == len(allZero) {
		return true
	}

	var v int = m.get(allZero[i])

	for isInSquare(v, allZero[i], m) || isInLine(v, allZero[i], m) {
		v++
		if v > 9 {
			m.set(allZero[i], 0)
			return false
		}
	}

	m.set(allZero[i], v)

	var res bool = false

	for !res {
		res = resolv(m, allZero, i+1)
		if !res {
			return resolv(m, allZero, i)
		}
	}

	return true
}

func getAllZero(m Map) (res V) {
	for y := 0; y != 9; y++ {
		for x := 0; x != 9; x++ {
			if m[y][x] == 0 {
				res = append(res, Pos{y: y, x: x})
			}
		}
	}
	return
}

func main() {
	m := getMap()
	allZero := getAllZero(m)

	resolv(m, allZero, 0)
	for y := 0; y < 9; y++ {
		for x := 0; x < 9; x++ {
			fmt.Printf("%d", m.get(Pos{x, y}))
		}
		fmt.Printf("\n")
	}
}
