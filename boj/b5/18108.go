package main

import (
	"bufio"
	. "fmt"
	"io"
	"io/ioutil"
	"strings"
)

func BOJ18108(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	// single test case
	var mod int = 2541 - 1998
	var i int
	Fscan(in, &i)

	Fprint(out, i-mod)

	// 입력 에러 방지
	_leftData, _ := ioutil.ReadAll(in)
	if _s := strings.TrimSpace(string(_leftData)); _s != "" {
		panic("읽지않은 데이터 발견：\n" + _s)
	}
}

// func main() { BOJ18108(os.Stdin, os.Stdout) }
