package main

import (
	"bufio"
	. "fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

func BOJ17903(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	// single test case
	var m, n int
	Fscan(in, &m, &n)

	var t int
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			Fscan(in, &t)
		}
	}

	if m >= 8 {
		Fprint(out, "satisfactory")
	} else {
		Fprint(out, "unsatisfactory")
	}

	// 입력 에러 방지
	_leftData, _ := ioutil.ReadAll(in)
	if _s := strings.TrimSpace(string(_leftData)); _s != "" {
		panic("읽지않은 데이터 발견：\n" + _s)
	}
}

// LETYLETYLETY
func main() { BOJ17903(os.Stdin, os.Stdout) }

// YTELYTELYTEL
