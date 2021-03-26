package main

import (
	"bufio"
	. "fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

func canRead(time, s, t, r, n int) bool {

	d := time / (t + r)
	mod := time % (t + r)
	p := 0

	if mod <= t {
		p = s*t*d + s*mod
	} else {
		p = s * t * (d + 1)
	}

	return p >= n
}

func BOJ6139(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	// single test case

	var n, k int
	Fscan(in, &n, &k)

	var s, t, r int

	for i := 0; i < k; i++ {
		Fscan(in, &s, &t, &r)

		start := 1
		end := 300000
		mid := start

		for j := 0; j < 100; j++ {
			mid = (start + end) >> 1

			if canRead(mid, s, t, r, n) {
				end = mid
			} else {
				start = mid + 1
			}
		}
		Fprint(out, start, "\n")
	}

	// 입력 에러 방지
	_leftData, _ := ioutil.ReadAll(in)
	if _s := strings.TrimSpace(string(_leftData)); _s != "" {
		panic("읽지않은 데이터 발견：\n" + _s)
	}
}

// LETYLETYLETY
func main() { BOJ6139(os.Stdin, os.Stdout) }

// YTELYTELYTEL
