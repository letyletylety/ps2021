package main

import (
	"bufio"
	. "fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

func BOJ15059(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	// single test case
	a := make([]int64, 3)
	for i := 0; i < 3; i++ {
		Fscan(in, &a[i])
	}
	ii := int64(0)
	for i := 0; i < 3; i++ {
		Fscan(in, &ii)
		a[i] = ii - a[i]
	}

	sum := int64(0)
	for i := 0; i < 3; i++ {
		if a[i] > 0 {
			sum += a[i]
		}
	}

	Fprint(out, sum)

	// 입력 에러 방지
	_leftData, _ := ioutil.ReadAll(in)
	if _s := strings.TrimSpace(string(_leftData)); _s != "" {
		panic("읽지않은 데이터 발견：\n" + _s)
	}
}

// LETYLETYLETY
func main() { BOJ15059(os.Stdin, os.Stdout) }

// YTELYTELYTEL
