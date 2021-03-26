package main

import (
	"bufio"
	. "fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

func BOJ5751(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	// single test case
	for {
		var n int
		Fscan(in, &n)
		if n == 0 {
			break
		}

		a := 0

		mary := 0
		john := 0
		for i := 0; i < n; i++ {
			Fscan(in, &a)

			if a == 0 {
				mary++
			} else {
				john++
			}
		}

		Fprintf(out, "Mary won %d times and John won %d times\n", mary, john)
	}

	// 입력 에러 방지
	_leftData, _ := ioutil.ReadAll(in)
	if _s := strings.TrimSpace(string(_leftData)); _s != "" {
		panic("읽지않은 데이터 발견：\n" + _s)
	}
}

// LETYLETYLETY
func main() { BOJ5751(os.Stdin, os.Stdout) }

// YTELYTELYTEL
