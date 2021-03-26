package main

import (
	"bufio"
	. "fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

func BOJ4635(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	// single test case

	for {
		var t int
		var sp, h int
		Fscan(in, &t)

		if t == -1 {
			break
		}

		elapsed := 0
		answer := 0

		for i := 0; i < t; i++ {
			Fscan(in, &sp, &h)
			answer += sp * (h - elapsed)
			elapsed = h
		}

		Fprint(out, answer)
		Fprint(out, " miles\n")
	}

	// 입력 에러 방지
	_leftData, _ := ioutil.ReadAll(in)
	if _s := strings.TrimSpace(string(_leftData)); _s != "" {
		panic("읽지않은 데이터 발견：\n" + _s)
	}
}

// LETYLETYLETY
func main() { BOJ4635(os.Stdin, os.Stdout) }

// YTELYTELYTEL
