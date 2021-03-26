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

func equality(aa, bb string) bool {
	if aa == bb {
		return true
	}
	return false
}

func next(aa string) string {
	aaa := strings.Split(aa, ":")

	h := aaa[0]
	m := aaa[1]
	s := aaa[2]

	hh, _ := strconv.Atoi(h)
	mm, _ := strconv.Atoi(m)
	ss, _ := strconv.Atoi(s)

	ss++
	if ss == 60 {
		mm++
		ss = 0
	}
	if mm == 60 {
		hh++
		mm = 0
	}
	if hh == 24 {
		hh = 0
		mm = 0
		ss = 0
	}

	return Sprintf("%02d:%02d:%02d", hh, mm, ss)
}

func mod3(aa string) bool {
	aaa := strings.Split(aa, ":")
	h := aaa[0]
	m := aaa[1]
	s := aaa[2]

	hh, _ := strconv.Atoi(h)
	mm, _ := strconv.Atoi(m)
	ss, _ := strconv.Atoi(s)

	hh = hh*10000 + mm*100 + ss
	// log.Print("hh: ", hh)
	if hh%3 == 0 {
		return true
	}
	return false
}

func BOJ1942(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	// single test case

	for i := 0; i < 3; i++ {
		var a, b string
		Fscan(in, &a, &b)
		answer := 0

		for equality(a, b) == false {
			if mod3(a) {
				answer++
			}

			a = next(a)
		}
		if mod3(b) {
			answer++
		}

		Fprint(out, answer, "\n")
	}

	// 입력 에러 방지
	_leftData, _ := ioutil.ReadAll(in)
	if _s := strings.TrimSpace(string(_leftData)); _s != "" {
		panic("읽지않은 데이터 발견：\n" + _s)
	}
}

// LETYLETYLETY
func main() { BOJ1942(os.Stdin, os.Stdout) }

// YTELYTELYTEL
