package main

import (
	"bufio"
	. "fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

func BOJ15921(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	// single test case
	var n int
	Fscan(in, &n)
	a := make([]int64, n)
	for i := 0; i < n; i++ {
		Fscan(in, &a[i])
	}

	if n == 0 {
		Fprint(out, "divide by zero")
	} else {
		Fprint(out, "1.00")
	}

	// 입력 에러 방지
	_leftData, _ := ioutil.ReadAll(in)
	if _s := strings.TrimSpace(string(_leftData)); _s != "" {
		panic("읽지않은 데이터 발견：\n" + _s)
	}
}

// LETYLETYLETY
func main() { BOJ15921(os.Stdin, os.Stdout) }

// YTELYTELYTEL
