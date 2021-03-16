package main

import (
	"bufio"
	. "fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

func CF266C(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	// single test case
	var n int
	Fscan(in, &n)

	a := make([]int64, n+1)
	s := int64(0)
	cnt := make([]int64, n+10)

	for i := 1; i <= n; i++ {
		Fscan(in, &a[i])
		s += a[i]
	}

	if s%3 != 0 {
		Fprint(out, 0)
	} else {
		s /= 3
		ss := int64(0)

		for i := n; i >= 1; i-- {
			ss += a[i]
			if ss == s {
				cnt[i] = 1
			}
		}

		for i := n - 1; i >= 1; i-- {
			cnt[i] += cnt[i+1]
		}

		ans := int64(0)
		ss = 0

		for i := 1; i < n-1; i++ {
			ss += a[i]
			if ss == s {
				ans += cnt[i+2]
			}
		}

		Fprint(out, ans)
	}

	// 입력 에러 방지
	_leftData, _ := ioutil.ReadAll(in)
	if _s := strings.TrimSpace(string(_leftData)); _s != "" {
		panic("읽지않은 데이터 발견：\n" + _s)
	}
}

// LETYLETYLETY
func main() { CF266C(os.Stdin, os.Stdout) }

// YTELYTELYTEL
