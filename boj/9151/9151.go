package main

import (
	"bufio"
	. "fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

func BOJ9151(_r io.Reader, _w io.Writer) {
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

		max := 0
		ret := 0
		tret := 0
		for i := 0; i < 60; i++ {
			ret = i * i * i

			if ret > n {
				break
			}

			for j := 0; j < 100; j++ {
				tret = ret + j*(j+1)*(j+2)/6
				if tret > n {
					break
				} else {
					if max < tret {
						max = tret
					}
				}
			}
		}
		Fprint(out, max, "\n")
	}

	// 입력 에러 방지
	_leftData, _ := ioutil.ReadAll(in)
	if _s := strings.TrimSpace(string(_leftData)); _s != "" {
		panic("읽지않은 데이터 발견：\n" + _s)
	}
}

// LETYLETYLETY
func main() { BOJ9151(os.Stdin, os.Stdout) }

// YTELYTELYTEL
