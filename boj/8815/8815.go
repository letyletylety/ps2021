package main

import (
	"bufio"
	. "fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

func BOJ8815(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	// single test case

	var z int
	Fscan(in, &z)

	roots := "BCDAB"

	for z > 0 {
		z--
		var n int
		Fscan(in, &n)

		if n == 1 {
			Fprintln(out, "A")
		} else {
			nn := (n - 2) % 12

			suite := nn / 3
			res := nn % 3
			ch := roots[suite]

			if res == 1 {
				ch = roots[suite+1]
			}

			Fprint(out, string(ch), "\n")
		}
	}

	// 입력 에러 방지
	_leftData, _ := ioutil.ReadAll(in)
	if _s := strings.TrimSpace(string(_leftData)); _s != "" {
		panic("읽지않은 데이터 발견：\n" + _s)
	}
}

// LETYLETYLETY
func main() { BOJ8815(os.Stdin, os.Stdout) }

// YTELYTELYTEL
