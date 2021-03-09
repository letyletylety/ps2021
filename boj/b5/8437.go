package main

import (
	"bufio"
	. "fmt"
	"io"
	"io/ioutil"
	"math/big"
	"strings"
)

func BOJ8437(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	// single test case
	var sum, more, a, b big.Int
	Fscan(in, &sum)
	Fscan(in, &more)

	sum.Add(&sum, &more)

	a.Div(&sum, big.NewInt(2))
	b.Sub(&a, &more)

	// sum + more = a + a
	// sum - more = b + b

	Fprintln(out, a.String())
	Fprintln(out, b.String())

	// 입력 에러 방지
	_leftData, _ := ioutil.ReadAll(in)
	if _s := strings.TrimSpace(string(_leftData)); _s != "" {
		panic("읽지않은 데이터 발견：\n" + _s)
	}
}

// func main() { BOJ8437(os.Stdin, os.Stdout) }
