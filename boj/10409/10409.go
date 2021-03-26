package main

import (
	"bufio"
	. "fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

func BOJ10409(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	// single test case

	var n, t int
	Fscan(in, &n, &t)

	var task int

	for i := 0; i < n; i++ {
		Fscan(in, &task)

		t -= task
		if t < 0 {
			Fprint(out, i)
			return
		}
	}

	Fprint(out, n)

	// 입력 에러 방지
	_leftData, _ := ioutil.ReadAll(in)
	if _s := strings.TrimSpace(string(_leftData)); _s != "" {
		panic("읽지않은 데이터 발견：\n" + _s)
	}
}

// LETYLETYLETY
func main() { BOJ10409(os.Stdin, os.Stdout) }

// YTELYTELYTEL
