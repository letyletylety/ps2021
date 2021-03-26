package main

import (
	"bufio"
	. "fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

var graph [][]int
var cat []int
var n, m int
var visited []bool

func move(start int, cats int) int {

	edges := graph[start]
	ret := 0

	consecutiveCats := 0

	// log.Print("start: ", start)
	// log.Print("edges: ", edges)

	// 만약 고양이가 있으면
	if cat[start] == 1 {
		consecutiveCats = cats + cat[start]
	}

	// 방문확인
	visited[start] = true

	child := 0

	// if consecutiveCats >= m no way
	if consecutiveCats > m {
		return 0
	}

	for _, v := range edges {
		if visited[v] {
			continue
		}
		child++
		ret += move(v, consecutiveCats)
	}

	if child == 0 {
		//leaf
		return 1
	}

	return ret
}

func CF580C(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	// single test case
	Fscan(in, &n, &m)

	cat = make([]int, n)
	for i := 0; i < n; i++ {
		Fscan(in, &cat[i])
	}

	graph = make([][]int, n+1)
	visited = make([]bool, n+1)

	var u, v int
	for i := 1; i < n; i++ {
		Fscan(in, &u, &v)
		u--
		v--
		graph[u] = append(graph[u], v)
		graph[v] = append(graph[v], u)
	}

	// for k := range graph {
	// 	for j := range graph[k] {
	// 		log.Print("j: ", j)
	// 	}
	// 	log.Print("\n")
	// }

	Fprint(out, move(0, 0))

	// 입력 에러 방지
	_leftData, _ := ioutil.ReadAll(in)
	if _s := strings.TrimSpace(string(_leftData)); _s != "" {
		panic("읽지않은 데이터 발견：\n" + _s)
	}
}

// LETYLETYLETY
func main() { CF580C(os.Stdin, os.Stdout) }

// YTELYTELYTEL
