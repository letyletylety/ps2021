package main

import (
	"bufio"
	. "fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

func BOJ19944(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	// single test case
	var n int
	Fscan(in, &n)
	var m int
	Fscan(in, &m)

	if m <= 2 {
		Fprint(out, "NEWBIE!")
	} else if m <= n {
		Fprint(out, "OLDBIE!")
	} else {
		Fprint(out, "TLE!")
	}

	// 입력 에러 방지
	_leftData, _ := ioutil.ReadAll(in)
	if _s := strings.TrimSpace(string(_leftData)); _s != "" {
		panic("읽지않은 데이터 발견：\n" + _s)
	}
}

// LETYLETYLETY
func main() { BOJ19944(os.Stdin, os.Stdout) }

// YTELYTELYTEL
