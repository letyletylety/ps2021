package main

import (
	"bufio"
	. "fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

func BOJ20232(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	// single test case

	var y int
	Fscan(in, &y)

	if y == 1996 || y == 1997 || y == 2000 || y == 2007 || y == 2008 || y == 2013 || y == 2018 {
		Fprint(out, "SPbSU")
	} else if y == 2006 {
		Fprint(out, "PetrSU, ITMO")
	} else {
		Fprint(out, "ITMO")
	}

	// 입력 에러 방지
	_leftData, _ := ioutil.ReadAll(in)
	if _s := strings.TrimSpace(string(_leftData)); _s != "" {
		panic("읽지않은 데이터 발견：\n" + _s)
	}
}

// LETYLETYLETY
func main() { BOJ20232(os.Stdin, os.Stdout) }

// YTELYTELYTEL
