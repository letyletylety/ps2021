package main

import (
	"bufio"
	. "fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

func BOJ19572(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	// single test case

	var d1, d2, d3 float64
	Fscan(in, &d1, &d2, &d3)

	abc := d1 + d2 + d3
	abc /= 2.0

	a := abc - d3
	b := abc - d2
	c := abc - d1

	if a <= 0 || b <= 0 || c <= 0 {
		Fprint(out, -1)
	} else {
		Fprint(out, 1, "\n")
		Fprintf(out, "%.1f %.1f %.1f", a, b, c)
	}

	// 입력 에러 방지
	_leftData, _ := ioutil.ReadAll(in)
	if _s := strings.TrimSpace(string(_leftData)); _s != "" {
		panic("읽지않은 데이터 발견：\n" + _s)
	}
}

// LETYLETYLETY
func main() { BOJ19572(os.Stdin, os.Stdout) }

// YTELYTELYTEL
