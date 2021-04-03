package main

import (
	"bufio"
	. "fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

func BOJ8678(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	// single test case

	var tt int
	var a, b int
	Fscan(in, &tt)

	for tt > 0 {
		tt--
		Fscan(in, &a, &b)

		if b%a == 0 {
			Fprint(out, "TAK", "\n")
		} else {
			Fprint(out, "NIE", "\n")
		}
	}

	// 입력 에러 방지
	_leftData, _ := ioutil.ReadAll(in)
	if _s := strings.TrimSpace(string(_leftData)); _s != "" {
		panic("읽지않은 데이터 발견：\n" + _s)
	}
}

// LETYLETYLETY
func main() { BOJ8678(os.Stdin, os.Stdout) }

// YTELYTELYTEL
