package main

import (
	"bufio"
	. "fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

func BOJ9286(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	// single test case

	var n int
	Fscan(in, &n)

	for i := 1; i <= n; i++ {
		var m int
		Fscan(in, &m)

		Fprint(out, "Case ", i, ":\n")

		for j := 0; j < m; j++ {

			var nn int
			Fscan(in, &nn)
			if nn < 6 {
				Fprintln(out, nn+1)
			}
		}

	}

	// 입력 에러 방지
	_leftData, _ := ioutil.ReadAll(in)
	if _s := strings.TrimSpace(string(_leftData)); _s != "" {
		panic("읽지않은 데이터 발견：\n" + _s)
	}
}

// LETYLETYLETY
func main() { BOJ9286(os.Stdin, os.Stdout) }

// YTELYTELYTEL
