package main

import (
	"bufio"
	. "fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

func BOJ9297(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	// single test case

	var n int
	Fscan(in, &n)

	for i := 1; i <= n; i++ {
		Fprint(out, "Case ", i, ": ")

		var a int
		Fscan(in, &a)
		var b int
		Fscan(in, &b)

		c := a / b
		d := a % b

		if d == 0 {

			Fprintf(out, "%d\n", c)
		} else if c == 0 {
			Fprintf(out, "%d/%d\n", d, b)
		} else {
			Fprintf(out, "%d %d/%d\n", c, d, b)
		}
	}

	// 입력 에러 방지
	_leftData, _ := ioutil.ReadAll(in)
	if _s := strings.TrimSpace(string(_leftData)); _s != "" {
		panic("읽지않은 데이터 발견：\n" + _s)
	}
}

// LETYLETYLETY
func main() { BOJ9297(os.Stdin, os.Stdout) }

// YTELYTELYTEL
