package main

import (
	"bufio"
	. "fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

func BOJ16199(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	// single test case
	var y, m, d int
	var y1, m1, d1 int

	Fscan(in, &y, &m, &d)
	Fscan(in, &y1, &m1, &d1)

	answer := y1 - y

	isOver := false

	if m < m1 {
		isOver = true
	} else if m == m1 {
		if d <= d1 {
			isOver = true
		}
	}

	if isOver == false {
		answer--
	}

	Fprintln(out, answer)
	Fprintln(out, y1-y+1)
	Fprintln(out, y1-y)

	// 입력 에러 방지
	_leftData, _ := ioutil.ReadAll(in)
	if _s := strings.TrimSpace(string(_leftData)); _s != "" {
		panic("읽지않은 데이터 발견：\n" + _s)
	}
}

// LETYLETYLETY
func main() { BOJ16199(os.Stdin, os.Stdout) }

// YTELYTELYTEL
