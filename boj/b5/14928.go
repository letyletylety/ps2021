package main

import (
	"bufio"
	. "fmt"
	"io"
	"io/ioutil"
	"math/big"
	"strings"
)

func BOJ14928(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	// single test case
	var n big.Int
	var m, result big.Int
	Fscan(in, &n)
	Sscan("20000303", &m)
	result.Mod(&n, &m)
	Fprint(out, result.String())

	// 입력 에러 방지
	_leftData, _ := ioutil.ReadAll(in)
	if _s := strings.TrimSpace(string(_leftData)); _s != "" {
		panic("읽지않은 데이터 발견：\n" + _s)
	}
}

// 끝나고 주석 처리
// func main() { BOJ14928(os.Stdin, os.Stdout) }
