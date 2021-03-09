package main

import (
	"bufio"
	. "fmt"
	"io"
	"io/ioutil"
	"strings"
)

func BOJ17496(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	// single test case
	var n, t, c, p int
	Fscan(in, &n, &t, &c, &p)

	answer := ((n - 1) / t) * c * p
	Fprint(out, answer)

	// 입력 에러 방지
	_leftData, _ := ioutil.ReadAll(in)
	if _s := strings.TrimSpace(string(_leftData)); _s != "" {
		panic("읽지않은 데이터 발견：\n" + _s)
	}
}

// TODO : 끝나고 주석 처리
// func main() { BOJ17496(os.Stdin, os.Stdout) }
