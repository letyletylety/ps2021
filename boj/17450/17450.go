package main

import (
	"bufio"
	. "fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

func BOJ17450(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	// single test case

	var a, b int
	var index int
	var ppa, ppb int

	index = 0
	ppa = 10000
	ppb = 0
	for i := 0; i < 3; i++ {
		Fscan(in, &a, &b)

		if a >= 500 {
			a -= 50
		}

		if ppb*a < b*ppa {
			ppa = a
			ppb = b
			index = i
		}
	}

	answer := "SNU"

	Fprint(out, string(answer[index]))

	// 입력 에러 방지
	_leftData, _ := ioutil.ReadAll(in)
	if _s := strings.TrimSpace(string(_leftData)); _s != "" {
		panic("읽지않은 데이터 발견：\n" + _s)
	}
}

// LETYLETYLETY
func main() { BOJ17450(os.Stdin, os.Stdout) }

// YTELYTELYTEL
