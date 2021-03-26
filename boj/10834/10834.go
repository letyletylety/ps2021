package main

import (
	"bufio"
	. "fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

func BOJ10834(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	// single test case

	var m int
	Fscan(in, &m)

	var rot int
	rot = 1

	clockwise := 0

	var a, b, c int

	for i := 0; i < m; i++ {
		Fscan(in, &a, &b, &c)

		rot = rot * b / a

		if c == 1 {
			clockwise = 1 - clockwise
		}
	}

	Fprint(out, clockwise, " ", rot)

	// 입력 에러 방지
	_leftData, _ := ioutil.ReadAll(in)
	if _s := strings.TrimSpace(string(_leftData)); _s != "" {
		panic("읽지않은 데이터 발견：\n" + _s)
	}
}

// LETYLETYLETY
func main() { BOJ10834(os.Stdin, os.Stdout) }

// YTELYTELYTEL
