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

func BOJ16428(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	// single test case
	a := new(big.Int)
	b := new(big.Int)
	c := new(big.Int)
	d := new(big.Int)

	Fscan(in, a)
	Fscan(in, b)

	c = c.Div(a, b)
	d = d.Mod(a, b)

	Fprintln(out, c)
	Fprintln(out, d)

	// 입력 에러 방지
	_leftData, _ := ioutil.ReadAll(in)
	if _s := strings.TrimSpace(string(_leftData)); _s != "" {
		panic("읽지않은 데이터 발견：\n" + _s)
	}
}

// LETYLETYLETY
func main() { BOJ16428(os.Stdin, os.Stdout) }

// YTELYTELYTEL
