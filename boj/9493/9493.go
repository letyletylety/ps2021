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

func BOJ9493(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	// single test case
	for {
		var m, a, b float64
		Fscan(in, &m, &a, &b)
		if m == 0 && a == 0 && b == 0 {
			break
		}

		bb := m * 3600 / b
		aa := m * 3600 / a
		x := aa - bb

		hh := int(x / 3600)
		m = x - float64(hh*3600)

		mm := int(m / 60)
		ss := int(math.Round(m - float64(mm*60)))

		Fprintf(out, "%d:%02d:%02d\n", hh, mm, ss)
	}

	// 입력 에러 방지
	_leftData, _ := ioutil.ReadAll(in)
	if _s := strings.TrimSpace(string(_leftData)); _s != "" {
		panic("읽지않은 데이터 발견：\n" + _s)
	}
}

// LETYLETYLETY
func main() { BOJ9493(os.Stdin, os.Stdout) }

// YTELYTELYTEL
