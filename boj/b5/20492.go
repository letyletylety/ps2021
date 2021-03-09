package main

import (
	"bufio"
	. "fmt"
	"io"
	"io/ioutil"
	"strings"
)

func BOJ20492(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	// single test case
	var n float64
	Fscan(in, &n)
	var n1, n2 float64
	n1 = n * 0.78
	n2 = n*0.8 + n*0.2*0.78

	Fprint(out, int(n1), int(n2))

	// 입력 에러 방지
	_leftData, _ := ioutil.ReadAll(in)
	if _s := strings.TrimSpace(string(_leftData)); _s != "" {
		panic("읽지않은 데이터 발견：\n" + _s)
	}
}

// TODO : 끝나고 주석 처리
// func main() { BOJ20492(os.Stdin, os.Stdout) }
