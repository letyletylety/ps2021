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

func BOJ5074(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	// single test case

	for {
		var a, b string
		Fscan(in, &a, &b)

		aa := strings.Split(a, ":")
		bb := strings.Split(b, ":")

		h, _ := strconv.Atoi(aa[0])
		m, _ := strconv.Atoi(aa[1])
		day := 0

		dh, _ := strconv.Atoi(bb[0])
		dm, _ := strconv.Atoi(bb[1])

		if h == 0 && m == 0 && dh == 0 && dm == 0 {
			break
		}

		m += dm
		if m >= 60 {
			m -= 60
			dh += 1
		}

		h += dh

		day = h / 24
		h %= 24

		if day == 0 {
			Fprintf(out, "%02d:%02d\n", h, m)
		} else {
			Fprintf(out, "%02d:%02d +%d\n", h, m, day)
		}
	}

	// 입력 에러 방지
	_leftData, _ := ioutil.ReadAll(in)
	if _s := strings.TrimSpace(string(_leftData)); _s != "" {
		panic("읽지않은 데이터 발견：\n" + _s)
	}
}

// LETYLETYLETY
func main() { BOJ5074(os.Stdin, os.Stdout) }

// YTELYTELYTEL
