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

func BOJ1837(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	// single test case

	var p string
	var k int

	Fscan(in, &p)
	Fscan(in, &k)

	pp := big.NewInt(0)
	mm := big.NewInt(0)
	pp.SetString(p, 10)
	// log.Print("pp: ", pp)
	for i := 2; i < k; i++ {
		ii := big.NewInt(int64(i))

		mm.Mod(pp, ii)
		// log.Print("ii: ", ii)
		// log.Print("mm: ", mm)
		if mm.Cmp(big.NewInt(0)) == 0 {
			Fprint(out, "BAD ", ii)
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
func main() { BOJ1837(os.Stdin, os.Stdout) }

// YTELYTELYTEL
