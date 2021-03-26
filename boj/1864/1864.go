package main

import (
	"bufio"
	. "fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

func BOJ1864(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	// single test case
	m := map[byte]int{
		'-':  0,
		'\\': 1,
		'(':  2,
		'@':  3,
		'?':  4,
		'>':  5,
		'&':  6,
		'%':  7,
		'/':  -1,
	}

	input := ""

	for {
		Fscan(in, &input)
		if input == "#" {
			break
		}

		answer := 0

		for _, v := range input {
			answer *= 8
			answer += m[byte(v)]
		}
		Fprint(out, answer)
		Fprint(out, "\n")
	}

	// 입력 에러 방지
	_leftData, _ := ioutil.ReadAll(in)
	if _s := strings.TrimSpace(string(_leftData)); _s != "" {
		panic("읽지않은 데이터 발견：\n" + _s)
	}
}

// LETYLETYLETY
func main() { BOJ1864(os.Stdin, os.Stdout) }

// YTELYTELYTEL
