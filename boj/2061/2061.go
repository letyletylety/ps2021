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

func BOJ2061(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	// single test case
	var k string
	var l int
	Fscan(in, &k, &l)

	a := big.NewInt(0)
	a.SetString(k, 10)

	mm := big.NewInt(0)

	for i := 2; i < l; i++ {
		mm.Mod(a, big.NewInt(int64(i)))

		if mm.Cmp(big.NewInt(0)) == 0 {
			Fprint(out, "BAD ", i)
			return
		}
	}

	Fprint(out, "GOOD")
	// 입력 에러 방지
	_leftData, _ := ioutil.ReadAll(in)
	if _s := strings.TrimSpace(string(_leftData)); _s != "" {
		panic("읽지않은 데이터 발견：\n" + _s)
	}
}

// LETYLETYLETY
func main() { BOJ2061(os.Stdin, os.Stdout) }

// YTELYTELYTEL
