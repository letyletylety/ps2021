package main

import (
	"bufio"
	. "fmt"
	"io"
	"io/ioutil"
	"strings"
	"unicode/utf8"
)

func BOJ11283(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	// single test case
	var input string
	Fscan(in, &input)

	r, _ := utf8.DecodeRuneInString(input)

	// Fprint(out, r)

	Fprint(out, r-44032+1)

	// 입력 에러 방지
	_leftData, _ := ioutil.ReadAll(in)
	if _s := strings.TrimSpace(string(_leftData)); _s != "" {
		panic("읽지않은 데이터 발견：\n" + _s)
	}
}

// func main() { BOJ11283(os.Stdin, os.Stdout) }
