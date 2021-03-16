package main

import (
	"bufio"
	. "fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

func CF704C(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	// single test case
	var n, m int
	Fscan(in, &n, &m)

	var s, t []byte
	Fscan(in, &s, &t)

	l := make([]int, m)
	r := make([]int, m)

	pos := 0
	for i, b := range t {
		for s[pos] != b {
			pos++
		}
		l[i] = pos
		pos++
	}

	pos = n - 1
	for i := m - 1; i >= 0; i-- {
		b := t[i]
		for s[pos] != b {
			pos--
		}
		r[i] = pos
		pos--
	}

	answer := 0
	for i := 1; i < m; i++ {
		answer = max(answer, max(max(l[i]-l[i-1], r[i]-r[i-1]), r[i]-l[i-1]))

	}
	Fprint(out, answer)

	// 입력 에러 방지
	_leftData, _ := ioutil.ReadAll(in)
	if _s := strings.TrimSpace(string(_leftData)); _s != "" {
		panic("읽지않은 데이터 발견：\n" + _s)
	}
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

// LETYLETYLETY
func main() { CF704C(os.Stdin, os.Stdout) }

// YTELYTELYTEL
