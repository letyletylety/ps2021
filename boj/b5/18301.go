package main

import (
	"bufio"
	. "fmt"
	"io"
	"io/ioutil"
	"math"
	"strings"
)

func BOJ18301(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	// single test case
	var n1, n2, n12 float64
	Fscan(in, &n1, &n2, &n12)

	answer := (n1+1)*(n2+1)/(n12+1) - 1

	Fprint(out, math.Floor(answer))
	// 입력 에러 방지
	_leftData, _ := ioutil.ReadAll(in)
	if _s := strings.TrimSpace(string(_leftData)); _s != "" {
		panic("읽지않은 데이터 발견：\n" + _s)
	}
}

// TODO : 끝나고 주석 처리
// func main() { BOJ18301(os.Stdin, os.Stdout) }
