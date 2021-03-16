package main

import (
	"bufio"
	. "fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

func BOJ18330(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	// single test case
	var n int
	Fscan(in, &n)
	var k int
	Fscan(in, &k)

	k += 60
	answer := 0

	if n > k {
		answer = 3000*(n-k) + k*1500
	} else {
		answer = 1500 * n
	}

	Fprint(out, answer)
	// 입력 에러 방지
	_leftData, _ := ioutil.ReadAll(in)
	if _s := strings.TrimSpace(string(_leftData)); _s != "" {
		panic("읽지않은 데이터 발견：\n" + _s)
	}
}

// LETYLETYLETY
func main() { BOJ18330(os.Stdin, os.Stdout) }

// YTELYTELYTEL
