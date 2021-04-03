package main

import (
	"bufio"
	. "fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

func BOJ14761(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	// single test case

	var x, y, n int
	Fscan(in, &x, &y, &n)
	chk := false
	for i := 1; i <= n; i++ {
		chk = false
		if i%x == 0 {
			Fprint(out, "Fizz")
			chk = true
		}
		if i%y == 0 {
			Fprint(out, "Buzz")
			chk = true
		}

		if !chk {
			Fprint(out, i)
		}

		Fprintln(out)
	}

	// 입력 에러 방지
	_leftData, _ := ioutil.ReadAll(in)
	if _s := strings.TrimSpace(string(_leftData)); _s != "" {
		panic("읽지않은 데이터 발견：\n" + _s)
	}
}

// LETYLETYLETY
func main() { BOJ14761(os.Stdin, os.Stdout) }

// YTELYTELYTEL
