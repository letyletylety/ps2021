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

func BOJ17874(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	// single test case

	var n, a, b int
	Fscan(in, &n, &a, &b)

	a = max(a, n-a)
	b = max(b, n-b)

	Fprint(out, a*b*4)

	// 입력 에러 방지
	_leftData, _ := ioutil.ReadAll(in)
	if _s := strings.TrimSpace(string(_leftData)); _s != "" {
		panic("읽지않은 데이터 발견：\n" + _s)
	}
}

// LETYLETYLETY
func main() { BOJ17874(os.Stdin, os.Stdout) }

// YTELYTELYTEL
