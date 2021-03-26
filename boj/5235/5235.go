package main

import (
	"bufio"
	. "fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

func BOJ5235(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	// single test case

	var t int
	Fscan(in, &t)
	for t > 0 {
		t--
		var n int
		Fscan(in, &n)

		even := 0
		odd := 0
		a := 0

		for i := 0; i < n; i++ {
			Fscan(in, &a)

			if a&1 > 0 {
				odd += a
			} else {
				even += a
			}
		}

		if even == odd {
			Fprint(out, "TIE")
		} else if even > odd {
			Fprint(out, "EVEN")
		} else {
			Fprint(out, "ODD")
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
func main() { BOJ5235(os.Stdin, os.Stdout) }

// YTELYTELYTEL
