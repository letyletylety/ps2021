package main

import (
	"bufio"
	. "fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

func BOJ20673(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()
	// single test case

	var p, q float64

	Fscan(in, &p, &q)

	if p <= 50 && q <= 10 {
		Fprint(out, "White")
	} else if q >= 30 {
		Fprint(out, "Red")
	} else {
		Fprint(out, "Yellow")
	}

	// 입력 에러 방지
	_leftData, _ := ioutil.ReadAll(in)
	if _s := strings.TrimSpace(string(_leftData)); _s != "" {
		panic("읽지않은 데이터 발견：\n" + _s)
	}
}

// LETYLETYLETY
func main() { BOJ20673(os.Stdin, os.Stdout) }

// YTELYTELYTEL
