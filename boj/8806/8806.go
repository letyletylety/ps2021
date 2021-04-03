package main

import (
	"bufio"
	. "fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

func BOJ8806(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	// single test case
	var z int
	Fscan(in, &z)

	for z > 0 {
		z--

		var a, b, c float64
		var a1, b1, c1 float64

		Fscan(in, &a, &b, &c)
		Fscan(in, &a1, &b1, &c1)

		adamwin := 0.0
		gosiawin := 0.0

		adamwin = a*b1 + b*c1 + c*a1
		gosiawin = a1*b + b1*c + c1*a

		if adamwin < gosiawin {
			Fprint(out, "GOSIA")
		} else if adamwin > gosiawin {
			Fprint(out, "ADAM")
		} else {
			Fprint(out, "=")
		}
		Fprint(out, "\n")
	}

	// 입력 에러 방지
	_leftData, _ := ioutil.ReadAll(in)
	if _s := strings.TrimSpace(string(_leftData)); _s != "" {
		panic("읽지않은 데이터 발견：\n" + _s)
	}
}

// LETYLETYLETY
func main() { BOJ8806(os.Stdin, os.Stdout) }

// YTELYTELYTEL
