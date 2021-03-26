package main

import (
	"bufio"
	. "fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

func BOJ15820(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	// single test case

	var s1, s2 int
	Fscan(in, &s1, &s2)

	var act, exp int

	answer := []string{"Accepted", "Wrong Answer", "Why Wrong!!!"}

	index := 0

	for i := 0; i < s1; i++ {
		Fscan(in, &act, &exp)

		if act != exp {
			index = 1
		}
	}

	for i := 0; i < s2; i++ {
		Fscan(in, &act, &exp)

		if act != exp && index == 0 {
			index = 2
		}
	}

	Fprint(out, answer[index])
	// 입력 에러 방지
	_leftData, _ := ioutil.ReadAll(in)
	if _s := strings.TrimSpace(string(_leftData)); _s != "" {
		panic("읽지않은 데이터 발견：\n" + _s)
	}
}

// LETYLETYLETY
func main() { BOJ15820(os.Stdin, os.Stdout) }

// YTELYTELYTEL
