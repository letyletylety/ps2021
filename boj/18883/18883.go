package main

import (
	"bufio"
	. "fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

func BOJ18883(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	// single test case

	var n int
	Fscan(in, &n)
	var m int
	Fscan(in, &m)

	cnt := 1
	for i := 0; i < n; i++ {
		Fprint(out, cnt)
		cnt++
		for j := 1; j < m; j++ {
			Fprint(out, " ", cnt)
			cnt++
		}
		Fprintln(out)
	}

	// 입력 에러 방지
	_leftData, _ := ioutil.ReadAll(in)
	if _s := strings.TrimSpace(string(_leftData)); _s != "" {
		panic("읽지않은 데이터 발견：\n" + _s)
	}
}

// LETYLETYLETY
func main() { BOJ18883(os.Stdin, os.Stdout) }

// YTELYTELYTEL
