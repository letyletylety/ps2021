package main

import (
	"bufio"
	. "fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

func BOJ15784(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	// single test case
	var n int
	Fscan(in, &n)

	var a int
	Fscan(in, &a)

	var b int
	Fscan(in, &b)

	var board [][]int

	board = make([][]int, n)
	for i := 0; i < n; i++ {
		board[i] = make([]int, n)
	}

	found := 0
	answer := []string{"HAPPY", "ANGRY"}

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			Fscan(in, &board[i][j])
		}
	}
	jin := board[a-1][b-1]

	for i := 0; i < n; i++ {
		if jin < board[a-1][i] {
			found = 1
		}

		if jin < board[i][b-1] {
			found = 1
		}
	}

	Fprint(out, answer[found])
	// 입력 에러 방지
	_leftData, _ := ioutil.ReadAll(in)
	if _s := strings.TrimSpace(string(_leftData)); _s != "" {
		panic("읽지않은 데이터 발견：\n" + _s)
	}
}

// LETYLETYLETY
func main() { BOJ15784(os.Stdin, os.Stdout) }

// YTELYTELYTEL
