package main

import (
	"bufio"
	. "fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

func star(n int, out *bufio.Writer) {

	if n == 1 {
		Fprint(out, "*\n")
		return
	}

	star(n-1, out)

	for i := 0; i < n; i++ {
		Fprint(out, "*")
	}
	Fprint(out, "\n")
}

func BOJ7595(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	// single test case

	for {
		var n int
		Fscan(in, &n)
		if n == 0 {
			break
		}

		star(n, out)
	}

	// 입력 에러 방지
	_leftData, _ := ioutil.ReadAll(in)
	if _s := strings.TrimSpace(string(_leftData)); _s != "" {
		panic("읽지않은 데이터 발견：\n" + _s)
	}
}

// LETYLETYLETY
func main() { BOJ7595(os.Stdin, os.Stdout) }

// YTELYTELYTEL
