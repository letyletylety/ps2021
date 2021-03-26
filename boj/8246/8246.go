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

func BOJ8246(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	// single test case

	var a, b, c int
	Fscan(in, &a, &b, &c)

	a /= c
	b /= c

	aa := a * b

	if a > 2 && b > 2 {
		aa -= (a - 2) * (b - 2)
	}

	Fprint(out, aa)

	// aa = (a/c)*2 + ((b-2*c)/c)*2
	// bb = (b/c)*2 + ((a-2*c)/c)*2
	// cc := (a / c) + ((b-c)/c)*2 + ((a - 2*c) / c)
	// dd := (b / c) + ((a-c)/c)*2 + ((b - 2*c) / c)

	// aa = max(aa, bb)
	// cc = max(cc, dd)
	// Fprint(out, max(aa, cc))

	// 입력 에러 방지
	_leftData, _ := ioutil.ReadAll(in)
	if _s := strings.TrimSpace(string(_leftData)); _s != "" {
		panic("읽지않은 데이터 발견：\n" + _s)
	}
}

// LETYLETYLETY
func main() { BOJ8246(os.Stdin, os.Stdout) }

// YTELYTELYTEL
