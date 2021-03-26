package main

import (
	"bufio"
	. "fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

func BOJ15780(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	// single test case
	var n int64
	Fscan(in, &n)
	var k int
	Fscan(in, &k)

	sum := int64(0)
	a := make([]int64, k)
	for i := 0; i < k; i++ {
		Fscan(in, &a[i])
		a[i] = (a[i] + 1) / 2
		sum += a[i]
	}

	if sum < n {
		Fprint(out, "NO")
	} else {
		Fprint(out, "YES")
	}
	// 3 4 -> 2
	// 5 6 -> 3

	// 입력 에러 방지
	_leftData, _ := ioutil.ReadAll(in)
	if _s := strings.TrimSpace(string(_leftData)); _s != "" {
		panic("읽지않은 데이터 발견：\n" + _s)
	}
}

// LETYLETYLETY
func main() { BOJ15780(os.Stdin, os.Stdout) }

// YTELYTELYTEL
