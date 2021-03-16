package main

import (
	"bufio"
	. "fmt"
	"io"
	"io/ioutil"
	
	"strings"
)

func BOJ11549(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	// single test case
	var t int
	var a, b, c, d, e int
	var answer int = 0

	Fscan(in, &t)
	Fscan(in, &a, &b, &c, &d, &e)

	if a == t {
		answer++
	}
	if b == t {
		answer++
	}
	if c == t {
		answer++
	}

	if d == t {
		answer++
	}
	if e == t {
		answer++
	}

	Fprint(out, answer)
	// 입력 에러 방지
	_leftData, _ := ioutil.ReadAll(in)
	if _s := strings.TrimSpace(string(_leftData)); _s != "" {
		panic("읽지않은 데이터 발견：\n" + _s)
	}
}

/*
func main() { BOJ11549(os.Stdin, os.Stdout) }

*/
