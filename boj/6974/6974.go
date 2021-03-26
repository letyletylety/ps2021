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

func BOJ6974(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	// single test case

	var ttt int
	Fscan(in, &ttt)

	for ttt > 0 {
		ttt--

		var a string
		var b string

		Fscan(in, &a, &b)

		aa := big.NewInt(0)
		bb := big.NewInt(0)
		mm := big.NewInt(0)
		qq := big.NewInt(0)

		aa.SetString(a, 10)
		bb.SetString(b, 10)

		qq.DivMod(aa, bb, mm)

		Fprint(out, qq, "\n")
		Fprint(out, mm, "\n")
		Fprint(out, "\n")
	}

	// 입력 에러 방지
	_leftData, _ := ioutil.ReadAll(in)
	if _s := strings.TrimSpace(string(_leftData)); _s != "" {
		panic("읽지않은 데이터 발견：\n" + _s)
	}
}

// LETYLETYLETY
func main() { BOJ6974(os.Stdin, os.Stdout) }

// YTELYTELYTEL
