package main

import (
	"bufio"
	. "fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func CF702f(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	// single test case

	var tt int
	Fscan(in, &tt)

	for tt > 0 {
		tt--

		var n int
		Fscan(in, &n)

		t := 0
		occur := map[int]int{}

		for i := 0; i < n; i++ {
			Fscan(in, &t)
			occur[t]++
		}

		cntOccur := make([]int, n+1)

		for _, v := range occur {
			cntOccur[v]++
		}

		save, s := 0, 0
		for i, c := range cntOccur {
			if i*(len(occur)-s) > save {
				save = i * (len(occur) - s)
			}
			s += c
		}
		Fprint(out, n-save, "\n")
	}

	// 입력 에러 방지
	_leftData, _ := ioutil.ReadAll(in)
	if _s := strings.TrimSpace(string(_leftData)); _s != "" {
		panic("읽지않은 데이터 발견：\n" + _s)
	}
}

// LETYLETYLETY
func main() { CF702f(os.Stdin, os.Stdout) }

// YTELYTELYTEL
