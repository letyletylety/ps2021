package main

import (
	"bufio"
	. "fmt"
	"io"
	"io/ioutil"
	"math"
	
	"strings"
)

func BOJ11466(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	// single test case
	var t1, t2 float64
	var a, b float64
	Fscan(in, &a, &b)

	if a > b {
		a, b = b, a
	}

	t1 = a / 2

	if a < b/3 {
		t2 = a
	} else {
		t2 = b / 3
	}

	Fprint(out,
		math.Max(t1, t2))

	// 입력 에러 방지
	_leftData, _ := ioutil.ReadAll(in)
	if _s := strings.TrimSpace(string(_leftData)); _s != "" {
		panic("읽지않은 데이터 발견：\n" + _s)
	}
}

/*
func main() { BOJ11466(os.Stdin, os.Stdout) }

*/
