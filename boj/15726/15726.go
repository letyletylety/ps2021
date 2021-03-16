package main

import (
	"bufio"
	. "fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

func BOJ15726(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	// single test case
	var a, b, c float64

	Fscan(in, &a, &b, &c)

	aa := a * b / c
	bb := a * c / b

	if aa < bb {
		aa = bb
	}

	Fprint(out, int(aa))

	// 입력 에러 방지
	_leftData, _ := ioutil.ReadAll(in)
	if _s := strings.TrimSpace(string(_leftData)); _s != "" {
		panic("읽지않은 데이터 발견：\n" + _s)
	}
}

// LETYLETYLETY
func main() { BOJ15726(os.Stdin, os.Stdout) }

// YTELYTELYTEL
