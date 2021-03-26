package main

import (
	"bufio"
	. "fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

func BOJ8295(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	// single test case

	var n, m, p int
	Fscan(in, &n, &m, &p)

	cnt := 0

	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			if 2*i+2*j >= p {
				cnt += (n - i + 1) * (m - j + 1)
			}
		}
	}
	Fprint(out, cnt)

	// 입력 에러 방지
	_leftData, _ := ioutil.ReadAll(in)
	if _s := strings.TrimSpace(string(_leftData)); _s != "" {
		panic("읽지않은 데이터 발견：\n" + _s)
	}
}

// LETYLETYLETY
func main() { BOJ8295(os.Stdin, os.Stdout) }

// YTELYTELYTEL
