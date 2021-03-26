package main

import (
	"bufio"
	. "fmt"
	"io"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func after20(h, m int) (int, int) {
	m += 20
	if m == 60 {
		m = 0
		h++
	}

	if h == 24 {
		h = 0
	}
	return h, m
}

func after10(h, m int) (int, int) {
	m += 10
	if m == 60 {
		m = 0
		h++
	}

	if h == 24 {
		h = 0
	}
	return h, m
}

func BOJ14041(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	// single test case

	var s string
	Fscan(in, &s)

	a := strings.Split(s, ":")

	b := 120

	h, _ := strconv.Atoi(a[0])
	m, _ := strconv.Atoi(a[1])

	for b > 0 {
		if h >= 7 && h <= 9 {
			b -= 10
			h, m = after20(h, m)
		} else if h >= 15 && h <= 18 {
			b -= 10
			h, m = after20(h, m)
		} else {
			if b > 10 {
				b -= 20
				h, m = after20(h, m)
			} else {
				b -= 10
				h, m = after10(h, m)
			}
		}
	}

	if b == 0 {
		Fprintf(out, "%02d:%02d", h, m)
	} else {
		Fprintf(out, "%02d:%02d", h, m)
	}
	// 입력 에러 방지
	_leftData, _ := ioutil.ReadAll(in)
	if _s := strings.TrimSpace(string(_leftData)); _s != "" {
		panic("읽지않은 데이터 발견：\n" + _s)
	}
}

// LETYLETYLETY
func main() { BOJ14041(os.Stdin, os.Stdout) }

// YTELYTELYTEL
