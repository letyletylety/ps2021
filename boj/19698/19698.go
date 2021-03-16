package main

import (
	"bufio"
	. "fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func BOJ19698(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	// single test case

	var n, w, h, l int
	Fscan(in, &n, &w, &h, &l)

	w = w / l
	h = h / l

	Fprint(out, min(w*h, n))

	// 입력 에러 방지
	_leftData, _ := ioutil.ReadAll(in)
	if _s := strings.TrimSpace(string(_leftData)); _s != "" {
		panic("읽지않은 데이터 발견：\n" + _s)
	}
}

// LETYLETYLETY
func main() { BOJ19698(os.Stdin, os.Stdout) }

// YTELYTELYTEL
