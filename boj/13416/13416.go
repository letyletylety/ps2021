package main

import (
	"bufio"
	. "fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
func BOJ13416(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	// single test case

	var tt int
	Fscan(in, &tt)

	for tt > 0 {
		tt--

		var n int
		Fscan(in, &n)

		var m, t int

		answer := 0
		for i := 0; i < n; i++ {
			m = 0
			for j := 0; j < 3; j++ {
				Fscan(in, &t)
				m = max(m, t)
			}
			answer += m
		}

		Fprint(out, answer, "\n")
	}

	// 입력 에러 방지
	_leftData, _ := ioutil.ReadAll(in)
	if _s := strings.TrimSpace(string(_leftData)); _s != "" {
		panic("읽지않은 데이터 발견：\n" + _s)
	}
}

// LETYLETYLETY
func main() { BOJ13416(os.Stdin, os.Stdout) }

// YTELYTELYTEL
