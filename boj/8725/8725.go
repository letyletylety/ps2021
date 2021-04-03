package main

import (
	"bufio"
	. "fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

func BOJ8725(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	// single test case

	var n int
	Fscan(in, &n)

	var board [][]int

	board = make([][]int, n)
	for i := 0; i < n; i++ {
		board[i] = make([]int, n)
	}

	ret := 0
	for i := 0; i < n; i++ {

		max := -1000001
		for j := 0; j < n; j++ {
			Fscan(in, &board[i][j])

			if max < board[i][j] {
				max = board[i][j]
			}
		}
		if max > 0 {
			ret += max
		}
	}

	Fprint(out, ret)

	// 입력 에러 방지
	_leftData, _ := ioutil.ReadAll(in)
	if _s := strings.TrimSpace(string(_leftData)); _s != "" {
		panic("읽지않은 데이터 발견：\n" + _s)
	}
}

// LETYLETYLETY
func main() { BOJ8725(os.Stdin, os.Stdout) }

// YTELYTELYTEL
