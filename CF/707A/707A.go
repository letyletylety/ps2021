package main

import (
	"bufio"
	. "fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

func CF707A(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	// single test case
	var n int
	Fscan(in, &n)
	a := make([]int64, n)
	for i := 0; i < n; i++ {
		Fscan(in, &a[i])
	}

	s := [5e6 + 1][2]int{}
	for i, ai := range a {
		for j := i + 1; j < n; j++ {
			sum := ai + a[j]
			if x, y := s[sum][0], s[sum][1]; y > 0 && i != x && i != y && j != x && j != y {
				Fprintln(out, "YES")
				Fprint(out, x+1, y+1, i+1, j+1)
				return
			}
			s[sum] = [2]int{i, j}
		}
	}
	Fprintln(out, "NO")
	// 입력 에러 방지
	_leftData, _ := ioutil.ReadAll(in)
	if _s := strings.TrimSpace(string(_leftData)); _s != "" {
		panic("읽지않은 데이터 발견：\n" + _s)
	}
}

// LETYLETYLETY
func main() { CF707A(os.Stdin, os.Stdout) }

// YTELYTELYTEL
