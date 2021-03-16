package main

import (
	"bufio"
	. "fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

func BOJ19602(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	// single test case
	var s, m, l int
	Fscan(in, &s, &m, &l)

	happy := s + m + m + l + l + l

	if happy >= 10 {
		Fprint(out, "happy")
	} else {
		Fprint(out, "sad")
	}

	// 입력 에러 방지
	_leftData, _ := ioutil.ReadAll(in)
	if _s := strings.TrimSpace(string(_leftData)); _s != "" {
		panic("읽지않은 데이터 발견：\n" + _s)
	}
}

// LETYLETYLETY
func main() { BOJ19602(os.Stdin, os.Stdout) }

// YTELYTELYTEL
