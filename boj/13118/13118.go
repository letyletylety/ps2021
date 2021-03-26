package main

import (
	"bufio"
	. "fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

func BOJ13118(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	// single test case

	a := make([]int, 4)
	for i := 0; i < 4; i++ {
		Fscan(in, &a[i])
	}

	var x, y, r int
	Fscan(in, &x, &y, &r)

	man := 0

	for i := 0; i < 4; i++ {
		if a[i] == x {
			man = i + 1
			break
		}
	}

	Fprint(out, man)

	// 입력 에러 방지
	_leftData, _ := ioutil.ReadAll(in)
	if _s := strings.TrimSpace(string(_leftData)); _s != "" {
		panic("읽지않은 데이터 발견：\n" + _s)
	}
}

// LETYLETYLETY
func main() { BOJ13118(os.Stdin, os.Stdout) }

// YTELYTELYTEL
