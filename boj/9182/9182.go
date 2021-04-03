package main

import (
	"bufio"
	. "fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

func BOJ9182(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	// single test case

	cas := 0
	for {
		cas++
		var p, e, i, d int
		Fscan(in, &p, &e, &i, &d)

		if p+e+i+d == -4 {
			break
		}

		for {
			d++
			p++
			e++
			i++
			p %= 23
			e %= 28
			i %= 33
			if p == 0 && e == 0 && i == 0 {
				break
			}
		}

		if d >= 21252 {
			d = 21252 - d
			d += 21252
		} else {
			d = 21252 - d
		}

		Fprintf(out, "Case %d: the next triple peak occurs in %d days.\n", cas, d)
	}

	// 입력 에러 방지
	_leftData, _ := ioutil.ReadAll(in)
	if _s := strings.TrimSpace(string(_leftData)); _s != "" {
		panic("읽지않은 데이터 발견：\n" + _s)
	}
}

// LETYLETYLETY
func main() { BOJ9182(os.Stdin, os.Stdout) }

// YTELYTELYTEL
