package main

import (
	"bufio"
	. "fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

func BOJ15474(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	// single test case
	var n, a, b, c, d int
	Fscan(in, &n, &a, &b, &c, &d)

	aa := 0
	cc := 0

	aa = ((n - 1) / a) + 1
	cc = ((n - 1) / c) + 1

	aa *= b
	cc *= d

	if aa < cc {
		cc = aa
	}

	Fprint(out, cc)

	// 입력 에러 방지
	_leftData, _ := ioutil.ReadAll(in)
	if _s := strings.TrimSpace(string(_leftData)); _s != "" {
		panic("읽지않은 데이터 발견：\n" + _s)
	}
}

// LETYLETYLETY
func main() { BOJ15474(os.Stdin, os.Stdout) }

// YTELYTELYTEL
