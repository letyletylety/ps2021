package main

import (
	"bufio"
	. "fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

func CF671C(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	// single test case
	var t int
	Fscan(in, &t)

	for t > 0 {
		t--

		var n, x int
		Fscan(in, &n, &x)

		cnt := 0
		sum := 0
		a := 0
		for i := 0; i < n; i++ {
			Fscan(in, &a)
			if a == x {
				cnt++
			}
			sum += a
		}

		if cnt == n {
			Fprint(out, 0)
		} else if cnt > 0 {
			Fprint(out, 1)
		} else if sum == n*x {
			Fprint(out, 1)
		} else {
			Fprint(out, 2)
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
func main() { CF671C(os.Stdin, os.Stdout) }

// YTELYTELYTEL
