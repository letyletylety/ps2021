package main

import (
	"bufio"
	. "fmt"
	"io"
	"io/ioutil"
	"math"
	
	"strings"
)

func BOJ13597(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	// single test case
	var a, b int
	Fscan(in, &a, &b)

	if a == b {
		Fprint(out, a)
	} else {
		Fprint(out, math.Max(float64(a), float64(b)))
	}

	// 입력 에러 방지
	_leftData, _ := ioutil.ReadAll(in)
	if _s := strings.TrimSpace(string(_leftData)); _s != "" {
		panic("읽지않은 데이터 발견：\n" + _s)
	}
}

/*
func main() { BOJ13597(os.Stdin, os.Stdout) }

*/
