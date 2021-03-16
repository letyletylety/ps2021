package main

import (
	"bufio"
	. "fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

func BOJ17362(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	// single test case
	var n int
	Fscan(in, &n)

	n = n % 8

	var ans int
	if n == 1 {
		ans = 1
	} else if n == 2 || n == 0 {
		ans = 2
	} else if n == 3 || n == 7 {
		ans = 3
	} else if n == 4 || n == 6 {
		ans = 4
	} else {
		ans = 5
	}
	Fprint(out, ans)

	// 입력 에러 방지
	_leftData, _ := ioutil.ReadAll(in)
	if _s := strings.TrimSpace(string(_leftData)); _s != "" {
		panic("읽지않은 데이터 발견：\n" + _s)
	}
}

// LETYLETYLETY
func main() { BOJ17362(os.Stdin, os.Stdout) }

// YTELYTELYTEL
