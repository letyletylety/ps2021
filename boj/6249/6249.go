package main

import (
	"bufio"
	. "fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

func BOJ6249(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	// single test case
	var n, p, h int
	Fscan(in, &n, &p, &h)

	var dollar int
	for i := 0; i < n; i++ {
		Fscan(in, &dollar)

		if p > dollar {
			Fprintf(out, "NTV: Dollar dropped by %d Oshloobs\n", p-dollar)
		}

		if h < dollar {
			h = dollar
			Fprintf(out, "BBTV: Dollar reached %d Oshloobs, A record!\n", h)
		}

		p = dollar
	}

	// 입력 에러 방지
	_leftData, _ := ioutil.ReadAll(in)
	if _s := strings.TrimSpace(string(_leftData)); _s != "" {
		panic("읽지않은 데이터 발견：\n" + _s)
	}
}

// LETYLETYLETY
func main() { BOJ6249(os.Stdin, os.Stdout) }

// YTELYTELYTEL
