package main

import (
	"bufio"
	. "fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

func BOJ11109(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	// single test case

	var tt int
	Fscan(in, &tt)

	for tt > 0 {
		tt--

		var d, n, s, p int
		Fscan(in, &d, &n, &s, &p)

		ss := d + n*p
		pp := n * s

		if ss < pp {
			Fprint(out, "parallelize")
		} else if ss == pp {
			Fprint(out, "does not matter")
		} else {
			Fprint(out, "do not parallelize")
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
func main() { BOJ11109(os.Stdin, os.Stdout) }

// YTELYTELYTEL
