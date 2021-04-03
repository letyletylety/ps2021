package main

import (
	"bufio"
	. "fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

func BOJ9288(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	// single test case

	var t int
	Fscan(in, &t)

	cas := 0
	for cas < t {
		cas++
		Fprint(out, "Case ", cas, ":\n")

		var n int
		Fscan(in, &n)

		for i := 1; i <= 6; i++ {
			j := n - i

			if i > j {
				break
			}
			if j >= 1 && j <= 6 {
				Fprintf(out, "(%d,%d)\n", i, j)
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
func main() { BOJ9288(os.Stdin, os.Stdout) }

// YTELYTELYTEL
