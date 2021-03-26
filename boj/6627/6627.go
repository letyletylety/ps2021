package main

import (
	"bufio"
	. "fmt"
	"io"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func s(n int) int {
	ss := strconv.Itoa(n)
	ret := 0
	for k := range ss {
		ret += int(ss[k] - '0')
	}
	return ret
}

func BOJ6627(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	// single test case

	for {
		var n int
		Fscan(in, &n)

		if n == 0 {
			break
		}

		sn := s(n)
		// log.Print("sn: ", sn)
		p := 11
		for {
			if s(p*n) == sn {
				Fprint(out, p, "\n")
				break
			}
			p++
		}
	}

	// 입력 에러 방지
	_leftData, _ := ioutil.ReadAll(in)
	if _s := strings.TrimSpace(string(_leftData)); _s != "" {
		panic("읽지않은 데이터 발견：\n" + _s)
	}
}

// LETYLETYLETY
func main() { BOJ6627(os.Stdin, os.Stdout) }

// YTELYTELYTEL
