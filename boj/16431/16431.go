package main

import (
	"bufio"
	. "fmt"
	"io"
	"io/ioutil"
	"math"
	"os"
	"strings"
)

func BOJ16431(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	// single test case
	var br, bc, dr, dc, jr, jc float64

	Fscan(in, &br, &bc, &dr, &dc, &jr, &jc)

	var result string

	answers := []string{"bessie", "daisy", "tie"}

	dr = jr - dr
	dc = jc - dc

	br = jr - br
	bc = jc - bc

	dr = math.Abs(float64(dr))
	dc = math.Abs(float64(dc))

	br = math.Abs(float64(br))
	bc = math.Abs(float64(bc))

	dr = dr + dc
	br = math.Max(br, bc)

	if dr == br {
		result = answers[2]
	} else if dr < br {
		result = answers[1]
	} else {
		result = answers[0]
	}

	Fprint(out, result)

	// 입력 에러 방지
	_leftData, _ := ioutil.ReadAll(in)
	if _s := strings.TrimSpace(string(_leftData)); _s != "" {
		panic("읽지않은 데이터 발견：\n" + _s)
	}
}

// LETYLETYLETY
func main() { BOJ16431(os.Stdin, os.Stdout) }

// YTELYTELYTEL
