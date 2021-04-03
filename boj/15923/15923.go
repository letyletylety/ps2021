package main

import (
	"bufio"
	. "fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

func BOJ15923(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	// single test case
	var n int
	Fscan(in, &n)

	var basea, baseb int
	var a, b, c, d int
	Fscan(in, &a, &b)

	basea = a
	baseb = b

	answer := 0
	for i := 1; i < n; i++ {
		Fscan(in, &c, &d)
		if a == c {
			dist := d - b
			if dist < 0 {
				answer -= dist
			} else {
				answer += dist
			}
		} else {
			dist := c - a
			if dist < 0 {
				answer -= dist
			} else {
				answer += dist
			}
		}
		a = c
		b = d
	}

	c = basea
	d = baseb
	if a == c {
		dist := d - b
		if dist < 0 {
			answer -= dist
		} else {
			answer += dist
		}
	} else {
		dist := c - a
		if dist < 0 {
			answer -= dist
		} else {
			answer += dist
		}
	}

	Fprint(out, answer)
	// 입력 에러 방지
	_leftData, _ := ioutil.ReadAll(in)
	if _s := strings.TrimSpace(string(_leftData)); _s != "" {
		panic("읽지않은 데이터 발견：\n" + _s)
	}
}

// LETYLETYLETY
func main() { BOJ15923(os.Stdin, os.Stdout) }

// YTELYTELYTEL
