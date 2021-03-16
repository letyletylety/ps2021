package main

import (
	"bufio"
	. "fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

func BOJ16727(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	// single test case
	var p1, s1, p2, s2 int
	Fscan(in, &p1, &s1, &s2, &p2)

	// p
	// s

	var psum, ssum int
	psum = p1 + p2
	ssum = s1 + s2

	if psum == ssum {
		if s1 == p2 {
			Fprint(out, "Penalty")
		} else if s1 > p2 {
			Fprint(out, "Esteghlal")
		} else {
			Fprint(out, "Persepolis")
		}
	} else if psum > ssum {
		Fprint(out, "Persepolis")
	} else {
		Fprint(out, "Esteghlal")
	}

	// 입력 에러 방지
	_leftData, _ := ioutil.ReadAll(in)
	if _s := strings.TrimSpace(string(_leftData)); _s != "" {
		panic("읽지않은 데이터 발견：\n" + _s)
	}
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

// LETYLETYLETY
func main() { BOJ16727(os.Stdin, os.Stdout) }

// YTELYTELYTEL
