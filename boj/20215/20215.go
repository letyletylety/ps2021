package main

import (
	"bufio"
	. "fmt"
	"io"
	"io/ioutil"
	"math"
	"os"
	"strings"
)

func BOJ20215(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	// single test case
	var w, h float64
	Fscan(in, &w, &h)

	answer := w + h - math.Hypot(w, h)
	Fprint(out, answer)

	// 입력 에러 방지
	_leftData, _ := ioutil.ReadAll(in)
	if _s := strings.TrimSpace(string(_leftData)); _s != "" {
		panic("읽지않은 데이터 발견：\n" + _s)
	}
}

// LETYLETYLETY
func main() { BOJ20215(os.Stdin, os.Stdout) }

// YTELYTELYTEL
