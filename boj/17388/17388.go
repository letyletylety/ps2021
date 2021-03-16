package main

import (
	"bufio"
	. "fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func BOJ17388(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	// single test case
	var s, k, h int
	Fscan(in, &s, &k, &h)

	if s+k+h >= 100 {
		Fprint(out, "OK")
	} else {

		answer := "Soongsil"
		if s > k {
			answer = "Korea"
			s = k
		}
		if s > h {
			answer = "Hanyang"
			s = h
		}
		Fprint(out, answer)
	}

	// 입력 에러 방지
	_leftData, _ := ioutil.ReadAll(in)
	if _s := strings.TrimSpace(string(_leftData)); _s != "" {
		panic("읽지않은 데이터 발견：\n" + _s)
	}
}

// LETYLETYLETY
func main() { BOJ17388(os.Stdin, os.Stdout) }

// YTELYTELYTEL
