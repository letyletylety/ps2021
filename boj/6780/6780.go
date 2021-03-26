package main

import (
	"bufio"
	. "fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

func BOJ6780(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	// single test case
	var t1, t2 int

	Fscan(in, &t1, &t2)

	ret := 2

	for {
		if t1 < t2 {
			break
		}
		t1, t2 = t2, t1-t2
		ret++
	}

	Fprint(out, ret)

	// 입력 에러 방지
	_leftData, _ := ioutil.ReadAll(in)
	if _s := strings.TrimSpace(string(_leftData)); _s != "" {
		panic("읽지않은 데이터 발견：\n" + _s)
	}
}

// LETYLETYLETY
func main() { BOJ6780(os.Stdin, os.Stdout) }

// YTELYTELYTEL
