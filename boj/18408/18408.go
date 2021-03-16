package main

import (
	"bufio"
	. "fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

func one(a int) {

}

func BOJ18408(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	// single test case
	var a, b, c int
	Fscan(in, &a, &b, &c)

	one := 0
	if a == 1 {
		one++
	}
	if b == 1 {
		one++
	}
	if c == 1 {
		one++
	}

	if one >= 2 {
		Fprint(out, 1)
	} else {
		Fprint(out, 2)
	}

	// 입력 에러 방지
	_leftData, _ := ioutil.ReadAll(in)
	if _s := strings.TrimSpace(string(_leftData)); _s != "" {
		panic("읽지않은 데이터 발견：\n" + _s)
	}
}

// LETYLETYLETY
func main() { BOJ18408(os.Stdin, os.Stdout) }

// YTELYTELYTEL
