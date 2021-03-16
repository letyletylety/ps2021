package main

import (
	"bufio"
	. "fmt"
	"io"
	"io/ioutil"
	
	"strings"
)

func BOJ13985(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	// single test case
	var a, b, c int
	var d, e string
	Fscan(in, &a, &d, &b, &e, &c)

	if a+b == c {
		Fprint(out, "YES")

	} else {
		Fprint(out, "NO")
	}

	// 입력 에러 방지
	_leftData, _ := ioutil.ReadAll(in)
	if _s := strings.TrimSpace(string(_leftData)); _s != "" {
		panic("읽지않은 데이터 발견：\n" + _s)
	}
}

/*
func main() { BOJ13985(os.Stdin, os.Stdout) }

*/
