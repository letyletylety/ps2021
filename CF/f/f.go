package main

import (
	"bufio"
	. "fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

func CFf(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	// single test case

	var n int
	Fscan(in, &n)

	for i := 0; i < n; i++ {

		temp := 0
		t := 0
		for i := 0; i < 5; i++ {
			Fscan(in, &t)
			temp *= 2
			temp += t
		}
		Fprint(out, temp)
	}

	// 입력 에러 방지
	_leftData, _ := ioutil.ReadAll(in)
	if _s := strings.TrimSpace(string(_leftData)); _s != "" {
		panic("읽지않은 데이터 발견：\n" + _s)
	}
}

// LETYLETYLETY
func main() { CFf(os.Stdin, os.Stdout) }

// YTELYTELYTEL
