package main

import (
	"bufio"
	. "fmt"
	"io"
	"io/ioutil"
	"os"
	"sort"
	"strings"
)

func BOJ20976(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	// single test case
	a := make([]int, 3)
	for i := 0; i < 3; i++ {
		Fscan(in, &a[i])
	}
	sort.Ints(a)

	Fprint(out, a[1])

	// 입력 에러 방지
	_leftData, _ := ioutil.ReadAll(in)
	if _s := strings.TrimSpace(string(_leftData)); _s != "" {
		panic("읽지않은 데이터 발견：\n" + _s)
	}
}

// LETYLETYLETY
func main() { BOJ20976(os.Stdin, os.Stdout) }

// YTELYTELYTEL
