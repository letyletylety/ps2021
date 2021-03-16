package main

import (
	"bufio"
	. "fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

func BOJ16785(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var a, b, c int
	Fscan(in, &a, &b, &c)

	sum := 0
	ans := 0

	for sum < c {
		ans++
		sum += a
		if ans%7 == 0 {
			sum += b
		}
	}

	Fprint(out, ans)

	// single test case

	// 입력 에러 방지
	_leftData, _ := ioutil.ReadAll(in)
	if _s := strings.TrimSpace(string(_leftData)); _s != "" {
		panic("읽지않은 데이터 발견：\n" + _s)
	}
}

// LETYLETYLETY
func main() { BOJ16785(os.Stdin, os.Stdout) }

// YTELYTELYTEL
