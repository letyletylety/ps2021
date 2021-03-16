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

func BOJ18411(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	// single test case
	a := make([]int, 3)
	Fscan(in, &a[0], &a[1], &a[2])
	sort.Ints(a)

	answer := a[1] + a[2]

	Fprint(out, answer)

	// 입력 에러 방지
	_leftData, _ := ioutil.ReadAll(in)
	if _s := strings.TrimSpace(string(_leftData)); _s != "" {
		panic("읽지않은 데이터 발견：\n" + _s)
	}
}

// LETYLETYLETY
func main() { BOJ18411(os.Stdin, os.Stdout) }

// YTELYTELYTEL
