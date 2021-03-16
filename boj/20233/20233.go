package main

import (
	"bufio"
	. "fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

func BOJ20233(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	// single test case

	var a, x, b, y, t int

	Fscan(in, &a, &x, &b, &y, &t)

	op1 := a
	op2 := b
	if t >= 30 {
		op1 = a + 21*(t-30)*x
	}

	if t >= 45 {
		op2 = b + 21*(t-45)*y
	}

	Fprint(out, op1, " ", op2)

	// 입력 에러 방지
	_leftData, _ := ioutil.ReadAll(in)
	if _s := strings.TrimSpace(string(_leftData)); _s != "" {
		panic("읽지않은 데이터 발견：\n" + _s)
	}
}

// LETYLETYLETY
func main() { BOJ20233(os.Stdin, os.Stdout) }

// YTELYTELYTEL
