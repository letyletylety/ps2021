package main

import (
	"bufio"
	. "fmt"
	"io"
	"io/ioutil"
	
	"strings"
)

func BOJ13136(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	// single test case

	var r, c, n int64
	Fscan(in, &r, &c, &n)

	Fprint(out, ((r-1)/n+1)*((c-1)/n+1))

	// 입력 에러 방지
	_leftData, _ := ioutil.ReadAll(in)
	if _s := strings.TrimSpace(string(_leftData)); _s != "" {
		panic("읽지않은 데이터 발견：\n" + _s)
	}
}

/*
func main() { BOJ13136(os.Stdin, os.Stdout) }

*/
