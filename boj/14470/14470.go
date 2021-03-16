package main

import (
	"bufio"
	. "fmt"
	"io"
	"io/ioutil"
	"math"
	"os"
	"strings"
)

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func BOJ14470(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	// single test case
	var a int
	Fscan(in, &a)
	var b int
	Fscan(in, &b)
	var c int
	Fscan(in, &c)
	var d int
	Fscan(in, &d)
	var e int
	Fscan(in, &e)

	var answer int

	if a > 0 {
		answer = (b - a) * e
	} else if a < 0 {
		answer = b*e + int(math.Abs(float64(a)))*c + d
	}
	Fprint(out, answer)

	// 입력 에러 방지
	_leftData, _ := ioutil.ReadAll(in)
	if _s := strings.TrimSpace(string(_leftData)); _s != "" {
		panic("읽지않은 데이터 발견：\n" + _s)
	}
}

// LETYLETYLETY
func main() { BOJ14470(os.Stdin, os.Stdout) }

// YTELYTELYTEL
