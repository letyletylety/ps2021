package main

import (
	"bufio"
	. "fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

func BOJ5101(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	// single test case

	var a, b, c int

	for {
		Fscan(in, &a, &b, &c)

		if a == 0 && b == 0 && c == 0 {
			break
		}

		c -= a
		c += b

		if c%b != 0 {
			Fprint(out, "X\n")
		} else {
			c /= b

			if c >= 1 {
				Fprint(out, c, "\n")
			} else {
				Fprint(out, "X\n")
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
func main() { BOJ5101(os.Stdin, os.Stdout) }

// YTELYTELYTEL
