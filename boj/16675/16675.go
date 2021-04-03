package main

import (
	"bufio"
	. "fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func BOJ16675(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	// single test case

	var a, b, c, d string
	Fscan(in, &a, &b, &c, &d)

	rsp := map[string]int{
		"S": 0,
		"R": 1,
		"P": 2,
	}

	// -1 0 1
	rwin := func(l, r string) int {
		rl := rsp[l]
		rr := rsp[r]

		// 2 0

		if ((rl < rr) && (rl-rr != -2)) || (rr-rl == -2) {
			return 1
		} else if rl == rr {
			return 0
		} else {
			return -1
		}
	}

	awin := func(a, b, c, d string) bool {
		if (rwin(c, a) == 1) && (rwin(d, a) == 1) {
			log.Print("a: ", a)
			return true
		}
		if (rwin(c, b) == 1) && (rwin(d, b) == 1) {
			log.Print("b: ", b)
			return true
		}
		return false
	}

	cwin := func(a, b, c, d string) bool {
		if (1 == rwin(a, c)) && (1 == rwin(b, c)) {
			log.Print("c: ", c)
			return true
		}
		if (1 == rwin(a, d)) && (1 == rwin(b, d)) {
			log.Print("d: ", d)
			return true
		}
		return false
	}

	if awin(a, b, c, d) {
		Fprint(out, "MS")
	} else if cwin(a, b, c, d) {
		Fprint(out, "TK")
	} else {
		Fprint(out, "?")
	}

	// 입력 에러 방지
	_leftData, _ := ioutil.ReadAll(in)
	if _s := strings.TrimSpace(string(_leftData)); _s != "" {
		panic("읽지않은 데이터 발견：\n" + _s)
	}
}

// LETYLETYLETY
func main() { BOJ16675(os.Stdin, os.Stdout) }

// YTELYTELYTEL
