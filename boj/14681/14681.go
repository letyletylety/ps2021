package main

import (
	"bufio"
	. "fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

func BOJ14681(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	// single test case
	var x int
	Fscan(in, &x)
	var y int
	Fscan(in, &y)

	answer := 0
	if x > 0 {
		if y > 0 {
			answer = 1
		} else {
			answer = 4
		}
	} else {
		if y > 0 {
			answer = 2
		} else {
			answer = 3
		}
	}

	Fprint(out, answer)

	// 입력 에러 방지
	_leftData, _ := ioutil.ReadAll(in)
	if _s := strings.TrimSpace(string(_leftData)); _s != "" {
		panic("읽지않은 데이터 발견：\n" + _s)
	}
}

// LETYLETYLETY
func main() { BOJ14681(os.Stdin, os.Stdout) }

// YTELYTELYTEL
