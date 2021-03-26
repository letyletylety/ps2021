package main

import (
	"bufio"
	. "fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

func fit(h, w, h1, w1 float64, per int) bool {

	h = h / 100.0 * float64(per)
	w = w / 100.0 * float64(per)

	if h <= h1 && w <= w1 {
		return true
	}

	if h <= w1 && w <= h1 {
		return true
	}
	return false
}

func BOJ4623(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	// single test case

	for {
		var a, b, c, d int

		Fscan(in, &a)
		Fscan(in, &b)
		Fscan(in, &c)
		Fscan(in, &d)

		if a == 0 && b == 0 && c == 0 && d == 0 {
			break
		}

		for i := 100; i >= 1; i-- {
			if fit(float64(a), float64(b), float64(c), float64(d), i) {
				Fprint(out, i, "%\n")
				break
			}
		}
	}

	// 입력 에러 방지
	_leftData, _ := ioutil.ReadAll(in)
	if _s := strings.TrimSpace(string(_leftData)); _s != "" {
		panic("읽지않은 데이터 발견：\n" + _s)
	}
}

// LETYLETYLETY
func main() { BOJ4623(os.Stdin, os.Stdout) }

// YTELYTELYTEL
