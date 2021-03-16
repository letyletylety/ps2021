package main

import (
	"bufio"
	. "fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

func BOJ15025(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	// single test case
	var l int
	Fscan(in, &l)
	var r int
	Fscan(in, &r)

	answers := []string{"Odd ", "Even ", "Not a moose"}

	if l == 0 && r == 0 {
		Fprint(out, answers[2])
	} else if l == r {
		Fprint(out, answers[1], l+r)
	} else {
		if l < r {
			l = r
		}
		Fprint(out, answers[0], l+l)
	}

	// 입력 에러 방지
	_leftData, _ := ioutil.ReadAll(in)
	if _s := strings.TrimSpace(string(_leftData)); _s != "" {
		panic("읽지않은 데이터 발견：\n" + _s)
	}
}

// LETYLETYLETY
func main() { BOJ15025(os.Stdin, os.Stdout) }

// YTELYTELYTEL
