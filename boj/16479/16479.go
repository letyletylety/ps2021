package main

import (
	"bufio"
	. "fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

func BOJ16479(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	// single test case
	var k, d1, d2 float64
	Fscan(in, &k, &d1, &d2)

	d1 -= d2
	d1 /= 2
	ans := k*k - d1*d1

	Fprint(out, ans)
	// 입력 에러 방지
	_leftData, _ := ioutil.ReadAll(in)
	if _s := strings.TrimSpace(string(_leftData)); _s != "" {
		panic("읽지않은 데이터 발견：\n" + _s)
	}
}

// LETYLETYLETY
func main() { BOJ16479(os.Stdin, os.Stdout) }

// YTELYTELYTEL
