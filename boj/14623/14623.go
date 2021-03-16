package main

import (
	"bufio"
	. "fmt"
	"io"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func BOJ14623(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	// single test case
	var b1 string
	Fscan(in, &b1)

	var b2 string
	Fscan(in, &b2)

	a, _ := strconv.ParseInt(b1, 2, 64)
	b, _ := strconv.ParseInt(b2, 2, 64)

	n := a * b

	answer := strconv.FormatInt(n, 2)

	Fprint(out, answer)

	// 입력 에러 방지
	_leftData, _ := ioutil.ReadAll(in)
	if _s := strings.TrimSpace(string(_leftData)); _s != "" {
		panic("읽지않은 데이터 발견：\n" + _s)
	}
}

// LETYLETYLETY
func main() { BOJ14623(os.Stdin, os.Stdout) }

// YTELYTELYTEL
