package main

import (
	"bufio"
	. "fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

func canwidereceiver(s, w, str float64) bool {
	if s <= 4.5 && w >= 150.0 && str >= 200 {
		return true
	}
	return false
}

func canLineman(s, w, str float64) bool {
	if s <= 6.0 && w >= 300.0 && str >= 500 {
		return true
	}
	return false
}

func canQua(s, w, str float64) bool {
	if s <= 5.0 && w >= 200.0 && str >= 300 {
		return true
	}
	return false
}

func BOJ4758(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	// single test case
	for {
		var a, b, c float64
		Fscan(in, &a, &b, &c)

		if a == 0 && b == 0 && c == 0 {
			break
		}

		hasPos := false

		if canwidereceiver(a, b, c) {
			Fprint(out, "Wide Receiver")
			hasPos = true
		}

		if canLineman(a, b, c) {
			if hasPos {
				Fprint(out, " ")
			}
			Fprint(out, "Lineman")
			hasPos = true
		}
		if canQua(a, b, c) {
			if hasPos {
				Fprint(out, " ")
			}
			Fprint(out, "Quarterback")
			hasPos = true
		}

		if !hasPos {
			Fprint(out, "No positions")
		}
		Fprint(out, "\n")
	}

	// 입력 에러 방지
	_leftData, _ := ioutil.ReadAll(in)
	if _s := strings.TrimSpace(string(_leftData)); _s != "" {
		panic("읽지않은 데이터 발견：\n" + _s)
	}
}

// LETYLETYLETY
func main() { BOJ4758(os.Stdin, os.Stdout) }

// YTELYTELYTEL
