package main

import (
	"bufio"
	. "fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

func BOJ11908(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	// single test case
	var n int
	Fscan(in, &n)

	a := make([]int, n)

	max := 0
	sum := 0
	for i := 0; i < n; i++ {
		Fscan(in, &a[i])

		sum += a[i]
		if max < a[i] {
			max = a[i]
		}
	}

	Fprint(out, sum-max)

	// 입력 에러 방지
	_leftData, _ := ioutil.ReadAll(in)
	if _s := strings.TrimSpace(string(_leftData)); _s != "" {
		panic("읽지않은 데이터 발견：\n" + _s)
	}
}

// LETYLETYLETY
func main() { BOJ11908(os.Stdin, os.Stdout) }

// YTELYTELYTEL
