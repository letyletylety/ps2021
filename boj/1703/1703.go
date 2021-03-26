package main

import (
	"bufio"
	. "fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

func BOJ1703(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	// single test case
	for {
		var a int
		Fscan(in, &a)
		if a == 0 {
			break
		}

		// log.Print("a: ", a)
		var b, c int
		answer := 1
		for i := 0; i < a; i++ {
			Fscan(in, &b)
			answer *= b
			Fscan(in, &c)
			answer -= c
		}
		// log.Print("answer: ", answer)
		Fprint(out, answer)
		Fprint(out, "\n")
	}

	// 입력 에러 방지
	_leftData, _ := ioutil.ReadAll(in)
	if _s := strings.TrimSpace(string(_leftData)); _s != "" {
		panic("읽지않은 데이터 발견：\n" + _s)
	}
}

// LETYLETYLETY
func main() { BOJ1703(os.Stdin, os.Stdout) }

// YTELYTELYTEL
