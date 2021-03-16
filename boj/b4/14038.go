package main

import (
	"bufio"
	. "fmt"
	"io"
	"io/ioutil"
	
	"strings"
)

func win(s string) int {
	if s == "W" {
		return 1
	} else {
		return 0
	}
}

func BOJ14038(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	// single test case
	var a, b, c, d, e, f string
	Fscan(in, &a, &b, &c, &d, &e, &f)

	answer := win(a) + win(b) + win(c) + win(d) + win(f) + win(e)

	if answer > 4 {
		Fprint(out, 1)
	} else if answer > 2 {

		Fprint(out, 2)
	} else if answer > 0 {

		Fprint(out, 3)
	} else {
		Fprint(out, -1)
	}

	// 입력 에러 방지
	_leftData, _ := ioutil.ReadAll(in)
	if _s := strings.TrimSpace(string(_leftData)); _s != "" {
		panic("읽지않은 데이터 발견：\n" + _s)
	}
}

/*
func main() { BOJ14038(os.Stdin, os.Stdout) }

*/
