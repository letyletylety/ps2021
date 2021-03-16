package main

import (
	"bufio"
	. "fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

func BOJ15051(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	// single test case
	var a, b, c int
	Fscan(in, &a, &b, &c)

	var s1, s2, s3 int
	s1 = b + b + c + c + c + c
	s2 = a + a + c + c
	s3 = a + a + a + a + b + b

	if s1 < s2 {
		s2 = s1
	}

	if s2 < s3 {
		s3 = s2
	}

	Fprint(out, s3)

	// 입력 에러 방지
	_leftData, _ := ioutil.ReadAll(in)
	if _s := strings.TrimSpace(string(_leftData)); _s != "" {
		panic("읽지않은 데이터 발견：\n" + _s)
	}
}

// LETYLETYLETY
func main() { BOJ15051(os.Stdin, os.Stdout) }

// YTELYTELYTEL
