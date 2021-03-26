package main

import (
	"bufio"
	. "fmt"
	"io"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func BOJ13877(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	// single test case

	var tt int
	Fscan(in, &tt)

	for tt > 0 {
		tt--

		var tc int64
		var n string
		Fscan(in, &tc, &n)

		Fprint(out, tc, " ")

		a, err := strconv.ParseInt(n, 8, 64)
		if err != nil {
			Fprint(out, 0, " ")
		} else {
			Fprint(out, a, " ")
		}
		b, _ := strconv.ParseInt(n, 10, 64)
		Fprint(out, b, " ")
		c, _ := strconv.ParseInt(n, 16, 64)
		Fprint(out, c, "\n")
	}

	// 입력 에러 방지
	_leftData, _ := ioutil.ReadAll(in)
	if _s := strings.TrimSpace(string(_leftData)); _s != "" {
		panic("읽지않은 데이터 발견：\n" + _s)
	}
}

// LETYLETYLETY
func main() { BOJ13877(os.Stdin, os.Stdout) }

// YTELYTELYTEL
