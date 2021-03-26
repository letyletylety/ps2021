package main

import (
	"bufio"
	. "fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

func BOJ14720(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	// single test case
	var n int
	Fscan(in, &n)

	need := 0

	t := 0
	ans := 0

	for i := 0; i < n; i++ {
		Fscan(in, &t)

		if t == need {
			ans++

			need++
			need %= 3
		}

	}

	Fprint(out, ans)

	// 입력 에러 방지
	_leftData, _ := ioutil.ReadAll(in)
	if _s := strings.TrimSpace(string(_leftData)); _s != "" {
		panic("읽지않은 데이터 발견：\n" + _s)
	}
}

// LETYLETYLETY
func main() { BOJ14720(os.Stdin, os.Stdout) }

// YTELYTELYTEL
