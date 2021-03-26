package main

import (
	"bufio"
	. "fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

type belt struct {
	end       int
	clockwise int
}

func BOJ6060(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	// single test case
	var n int
	Fscan(in, &n)

	beltsClock := make([]belt, n)

	var s, d, c int

	for i := 1; i < n; i++ {
		Fscan(in, &s, &d, &c)
		beltsClock[s] = belt{end: d, clockwise: c}
	}

	current := 1
	answer := 0

	for current < n {
		// Fprint(out, beltsClock[current])

		c := beltsClock[current].clockwise

		if c == 1 {
			answer = 1 - answer
		}

		current = beltsClock[current].end
	}

	Fprint(out, answer)

	// 입력 에러 방지
	_leftData, _ := ioutil.ReadAll(in)
	if _s := strings.TrimSpace(string(_leftData)); _s != "" {
		panic("읽지않은 데이터 발견：\n" + _s)
	}
}

// LETYLETYLETY
func main() { BOJ6060(os.Stdin, os.Stdout) }

// YTELYTELYTEL
