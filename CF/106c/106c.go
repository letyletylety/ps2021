package main

import (
	"bufio"
	. "fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

func min(a, b int64) int64 {
	if a > b {
		return b
	}
	return a
}

func CF106c(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	// single test case
	var tt int
	Fscan(in, &tt)

	for tt > 0 {
		tt--

		var nn int64
		Fscan(in, &nn)

		mn := []int64{1e9, 1e9}
		rem := []int64{int64(nn), int64(nn)}
		sum := int64(0)
		ans := int64(1e18)

		a := make([]int64, nn)
		for i := int64(0); i < nn; i++ {
			Fscan(in, &a[i])
		}

		for k := range a {
			mn[k%2] = min(mn[k%2], a[k])
			rem[k%2]--
			sum += a[k]
			if k > 0 {
				cur := sum + rem[0]*mn[0] + rem[1]*mn[1]
				ans = min(ans, cur)
			}
		}

		Fprint(out, ans, "\n")
	}

	// 입력 에러 방지
	_leftData, _ := ioutil.ReadAll(in)
	if _s := strings.TrimSpace(string(_leftData)); _s != "" {
		panic("읽지않은 데이터 발견：\n" + _s)
	}
}

// LETYLETYLETY
func main() { CF106c(os.Stdin, os.Stdout) }

// YTELYTELYTEL
