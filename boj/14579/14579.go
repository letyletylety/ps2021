package main

import (
	"bufio"
	. "fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

func BOJ14579(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	// single test case

	var a, b int
	Fscan(in, &a, &b)

	answer := 1
	mod := 14579

	for i := a; i <= b; i++ {
		answer *= i * (i + 1) / 2
		answer %= mod
	}

	Fprint(out, answer)

	// 입력 에러 방지
	_leftData, _ := ioutil.ReadAll(in)
	if _s := strings.TrimSpace(string(_leftData)); _s != "" {
		panic("읽지않은 데이터 발견：\n" + _s)
	}
}

// LETYLETYLETY
func main() { BOJ14579(os.Stdin, os.Stdout) }

// YTELYTELYTEL
