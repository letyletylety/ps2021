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

func BOJ16693(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	// single test case
	var a, p1, r, p2 float64
	Fscan(in, &a, &p1, &r, &p2)

	a = a / p1
	r = r * r * math.Pi
	r = r / p2

	if a < r {
		Fprint(out, "Whole pizza")
	} else {
		Fprint(out, "Slice of pizza")
	}

	// 입력 에러 방지
	_leftData, _ := ioutil.ReadAll(in)
	if _s := strings.TrimSpace(string(_leftData)); _s != "" {
		panic("읽지않은 데이터 발견：\n" + _s)
	}
}

// LETYLETYLETY
func main() { BOJ16693(os.Stdin, os.Stdout) }

// YTELYTELYTEL
