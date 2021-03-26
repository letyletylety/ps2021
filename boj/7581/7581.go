package main

import (
	"bufio"
	. "fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

func BOJ7581(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	// single test case

	for {
		var a, b, c, d int
		Fscan(in, &a, &b, &c, &d)

		if a+b+c+d == 0 {
			break
		}

		if a == 0 {
			a = d / b / c
		}
		if b == 0 {
			b = d / a / c
		}
		if c == 0 {
			c = d / a / b
		}
		if d == 0 {
			d = a * b * c
		}

		Fprint(out, a, " ", b, " ", c, " ", d, "\n")
	}

	// 입력 에러 방지
	_leftData, _ := ioutil.ReadAll(in)
	if _s := strings.TrimSpace(string(_leftData)); _s != "" {
		panic("읽지않은 데이터 발견：\n" + _s)
	}
}

// LETYLETYLETY
func main() { BOJ7581(os.Stdin, os.Stdout) }

// YTELYTELYTEL
