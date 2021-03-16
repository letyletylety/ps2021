package main

import (
	"bufio"
	. "fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

func BOJ15128(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	// single test case
	var p1, q1, p2, q2 int64
	Fscan(in, &p1, &q1, &p2, &q2)

	p1 = p1 * p2
	q1 = q1 * q2 * 2

	if p1%q1 == 0 {
		Fprint(out, 1)
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
func main() { BOJ15128(os.Stdin, os.Stdout) }

// YTELYTELYTEL
