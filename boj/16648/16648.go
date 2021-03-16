package main

import (
	"bufio"
	. "fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

func BOJ16648(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	// single test case
	var t float64
	Fscan(in, &t)
	var p float64
	Fscan(in, &p)

	answer := 100 - p

	if p < 20 {
		answer = t * 2 * p / (120 - 2*p)
	} else {
		answer = t * (p + 20) / (100 - p)
	}

	Fprint(out, answer)

	// 입력 에러 방지
	_leftData, _ := ioutil.ReadAll(in)
	if _s := strings.TrimSpace(string(_leftData)); _s != "" {
		panic("읽지않은 데이터 발견：\n" + _s)
	}
}

// LETYLETYLETY
func main() { BOJ16648(os.Stdin, os.Stdout) }

// YTELYTELYTEL
