package main

import (
	"bufio"
	. "fmt"
	"io"
	"io/ioutil"
	"strings"
)

func BOJ6764(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	// single test case
	var a, b, c, d int
	Fscan(in, &a, &b, &c, &d)

	a = b - a
	b = c - b
	c = d - c

	var result string
	if a > 0 && b > 0 && c > 0 {
		result = "Fish Rising"
	} else if a < 0 && b < 0 && c < 0 {
		result = "Fish Diving"
	} else if a == 0 && b == 0 && c == 0 {
		result = "Fish At Constant Depth"
	} else {
		result = "No Fish"
	}
	Fprint(out, result)

	// 입력 에러 방지
	_leftData, _ := ioutil.ReadAll(in)
	if _s := strings.TrimSpace(string(_leftData)); _s != "" {
		panic("읽지않은 데이터 발견：\n" + _s)
	}
}

/*
func main() { BOJ6764(os.Stdin, os.Stdout) }

*/
