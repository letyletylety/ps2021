package main

import (
	"bufio"
	. "fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

func isTele(a, b, c, d int) bool {
	if a != 8 && a != 9 {
		return false
	}

	if b != c {
		return false
	}

	if d != 8 && d != 9 {
		return false
	}

	return true
}

func BOJ16017(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	// single test case

	var a, b, c, d int
	Fscan(in, &a, &b, &c, &d)

	if isTele(a, b, c, d) {
		Fprint(out, "ignore")
	} else {
		Fprint(out, "answer")
	}

	// 입력 에러 방지
	_leftData, _ := ioutil.ReadAll(in)
	if _s := strings.TrimSpace(string(_leftData)); _s != "" {
		panic("읽지않은 데이터 발견：\n" + _s)
	}
}

// LETYLETYLETY
func main() { BOJ16017(os.Stdin, os.Stdout) }

// YTELYTELYTEL
