package main

import (
	"bufio"
	. "fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func mcost(c, d int) int {
	ret := c*30 + d*40

	ret = min(ret, c*35+d*30)
	ret = min(ret, c*40+d*20)

	return ret
}

func BOJ15340(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	// single test case

	var c, d int
	for {
		Fscan(in, &c, &d)
		if c+d == 0 {
			break
		}

		Fprint(out, mcost(c, d), "\n")
	}

	// 입력 에러 방지
	_leftData, _ := ioutil.ReadAll(in)
	if _s := strings.TrimSpace(string(_leftData)); _s != "" {
		panic("읽지않은 데이터 발견：\n" + _s)
	}
}

// LETYLETYLETY
func main() { BOJ15340(os.Stdin, os.Stdout) }

// YTELYTELYTEL
