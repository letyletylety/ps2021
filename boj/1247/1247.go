package main

import (
	"bufio"
	. "fmt"
	"io"
	"io/ioutil"
	"math/big"
	"os"
	"strings"
)

func BOJ1247(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	// single test case

	for tt := 0; tt < 3; tt++ {
		var n int
		Fscan(in, &n)

		var a int64
		sum := big.NewInt(0)
		for i := 0; i < n; i++ {
			Fscan(in, &a)
			sum.Add(sum, big.NewInt(a))
		}

		ret := sum.Cmp(big.NewInt(0))
		if ret == -1 {
			Fprintln(out, "-")
		} else if ret == 0 {
			Fprintln(out, "0")
		} else {
			Fprintln(out, "+")
		}
	}
	// 입력 에러 방지
	_leftData, _ := ioutil.ReadAll(in)
	if _s := strings.TrimSpace(string(_leftData)); _s != "" {
		panic("읽지않은 데이터 발견：\n" + _s)
	}
}

// LETYLETYLETY
func main() { BOJ1247(os.Stdin, os.Stdout) }

// YTELYTELYTEL
