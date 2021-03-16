package main

import (
	"bufio"
	. "fmt"
	"io"
	"io/ioutil"
	"strings"
)

func BOJ5928(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	// single test case
	var d, h, m int
	Fscan(in, &d, &h, &m)

	if d == 11 {
	}

	cur := d*24*60 + h*60 + m
	sta := 11*24*60 + 11*60 + 11

	cur = cur - sta

	if cur < 0 {
		Fprint(out, -1)
	} else {
		Fprint(out, cur)
	}
	// 입력 에러 방지
	_leftData, _ := ioutil.ReadAll(in)
	if _s := strings.TrimSpace(string(_leftData)); _s != "" {
		panic("읽지않은 데이터 발견：\n" + _s)
	}
}

/*
func main() { BOJ5928(os.Stdin, os.Stdout) }
*/
