package main

import (
	"bufio"
	. "fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

func BOJ15792(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	// single test case
	var a, b int

	Fscan(in, &a, &b)
	qu := a / b

	Fprint(out, qu, ".")

	a %= b
	// assert(
	for i := 0; i < 1500; i++ {
		a *= 10
		Fprint(out, a/b)
		a %= b
	}

	// 입력 에러 방지
	_leftData, _ := ioutil.ReadAll(in)
	if _s := strings.TrimSpace(string(_leftData)); _s != "" {
		panic("읽지않은 데이터 발견：\n" + _s)
	}
}

// LETYLETYLETY
func main() { BOJ15792(os.Stdin, os.Stdout) }

// YTELYTELYTEL
