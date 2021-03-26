package main

import (
	"bufio"
	. "fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

func BOJ7598(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	// single test case
	var n int
	Fscan(in, &n)

	for {
		var a string
		var b int
		Fscan(in, &a, &b)

		// log.Print("a: ", a)
		// log.Print("b: ", b)
		if a == "#" && b == 0 {
			break
		}

		for {
			var c string
			var d int
			Fscan(in, &c, &d)

			if c == "X" && d == 0 {
				break
			}

			if c == "B" {
				if b+d > 68 {
					// ignore
				} else {
					b += d
				}
			}

			if c == "C" {
				if d > b {
					// ignore
				} else {
					b -= d
				}
			}
		}

		Fprint(out, a, " ", b, "\n")
	}

	// 입력 에러 방지
	_leftData, _ := ioutil.ReadAll(in)
	if _s := strings.TrimSpace(string(_leftData)); _s != "" {
		panic("읽지않은 데이터 발견：\n" + _s)
	}
}

// LETYLETYLETY
func main() { BOJ7598(os.Stdin, os.Stdout) }

// YTELYTELYTEL
