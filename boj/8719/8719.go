package main

import (
	"bufio"
	. "fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

func BOJ8719(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	// single test case
	var t int
	Fscan(in, &t)

	var a, b int64
	for t > 0 {
		t--

		Fscan(in, &a, &b)

		cnt := 0

		for a < b {
			a *= 2
			cnt++
		}

		Fprint(out, cnt, "\n")

	}

	// 입력 에러 방지
	_leftData, _ := ioutil.ReadAll(in)
	if _s := strings.TrimSpace(string(_leftData)); _s != "" {
		panic("읽지않은 데이터 발견：\n" + _s)
	}
}

// LETYLETYLETY
func main() { BOJ8719(os.Stdin, os.Stdout) }

// YTELYTELYTEL
