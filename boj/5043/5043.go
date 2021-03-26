package main

import (
	"bufio"
	. "fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

func BOJ5043(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	// single test case
	var n int64
	var b int64

	Fscan(in, &n)
	Fscan(in, &b)

	b = 1<<(b+1) - 1

	if n <= b {
		Fprint(out, "yes")
	} else {
		Fprint(out, "no")
	}

	// 입력 에러 방지
	_leftData, _ := ioutil.ReadAll(in)
	if _s := strings.TrimSpace(string(_leftData)); _s != "" {
		panic("읽지않은 데이터 발견：\n" + _s)
	}
}

// LETYLETYLETY
func main() { BOJ5043(os.Stdin, os.Stdout) }

// YTELYTELYTEL
