package main

import (
	"bufio"
	. "fmt"
	"io"
	"io/ioutil"
	
	"strings"
)

func BOJ8723(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	// single test case
	var a, b, c int
	Fscan(in, &a, &b, &c)

	if a > b {
		a, b = b, a
	}
	if b > c {
		b, c = c, b
	}
	if a > c {
		a, c = c, a
	}

	if a == b && b == c {
		Fprint(out, 2)
	} else if a*a+b*b == c*c {
		Fprint(out, 1)
	} else {
		Fprint(out, 0)
	}
	// 입력 에러 방지
	_leftData, _ := ioutil.ReadAll(in)
	if _s := strings.TrimSpace(string(_leftData)); _s != "" {
		panic("읽지않은 데이터 발견：\n" + _s)
	}
}

/*
func main() { BOJ8723(os.Stdin, os.Stdout) }

*/
