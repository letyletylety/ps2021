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

func Reverse(b string) string {
	if len(b) == 0 {
		return ""
	}

	return Reverse(b[1:]) + string(b[0])
}

func BOJ3486(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	// single test case
	var tt int
	Fscan(in, &tt)

	for {
		if tt == 0 {
			break
		}

		var a, b string
		Fscan(in, &a, &b)

		a = Reverse(a)
		b = Reverse(b)

		aa, _ := strconv.Atoi(a)
		bb, _ := strconv.Atoi(b)

		aa += bb

		a = strconv.Itoa(aa)
		a = Reverse(a)

		aa, _ = strconv.Atoi(a)
		Fprint(out, aa, "\n")

		tt--
	}

	// 입력 에러 방지
	_leftData, _ := ioutil.ReadAll(in)
	if _s := strings.TrimSpace(string(_leftData)); _s != "" {
		panic("읽지않은 데이터 발견：\n" + _s)
	}
}

// LETYLETYLETY
func main() { BOJ3486(os.Stdin, os.Stdout) }

// YTELYTELYTEL
