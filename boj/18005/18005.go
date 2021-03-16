package main

import (
	"bufio"
	. "fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

func BOJ18005(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	// single test case
	var n int
	Fscan(in, &n)

	if n == 1 {
		Fprint(out, 0)
	} else if n%2 == 0 {
		// n is even
		nn := n / 2
		if nn%2 == 0 {
			Fprint(out, 2)
		} else {
			Fprint(out, 1)
		}
	} else {
		Fprint(out, 0)
	}

	// 입력 에러 방지
	_leftData, _ := ioutil.ReadAll(in)
	if _s := strings.TrimSpace(string(_leftData)); _s != "" {
		panic("읽지않은 데이터 발견：\n" + _s)
	}
}

// LETYLETYLETY
func main() { BOJ18005(os.Stdin, os.Stdout) }

// YTELYTELYTEL
