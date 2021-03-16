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

func BOJ15048(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	// single test case
	var n string
	Fscan(in, &n)

	sum := 0
	for _, v := range n {
		ii, _ := strconv.Atoi(string(v))
		sum += ii
	}
	Fprint(out, sum%3)

	// 입력 에러 방지
	_leftData, _ := ioutil.ReadAll(in)
	if _s := strings.TrimSpace(string(_leftData)); _s != "" {
		panic("읽지않은 데이터 발견：\n" + _s)
	}
}

// LETYLETYLETY
func main() { BOJ15048(os.Stdin, os.Stdout) }

// YTELYTELYTEL
