package main

import (
	"bufio"
	. "fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

func BOJ17009(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	// single test case
	var a, b, c, d, e, f int
	Fscan(in, &a, &b, &c, &d, &e, &f)

	a = 3*a + 2*b + c
	d = 3*d + 2*e + f

	if a < d {
		Fprint(out, "B")
	} else if a > d {
		Fprint(out, "A")
	} else {
		Fprint(out, "T")
	}

	// 입력 에러 방지
	_leftData, _ := ioutil.ReadAll(in)
	if _s := strings.TrimSpace(string(_leftData)); _s != "" {
		panic("읽지않은 데이터 발견：\n" + _s)
	}
}

// LETYLETYLETY
func main() { BOJ17009(os.Stdin, os.Stdout) }

// YTELYTELYTEL
