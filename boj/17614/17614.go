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

func BOJ17614(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	// single test case

	var n int
	Fscan(in, &n)

	sum := 0

	var s string
	var nn int

	for i := 1; i <= n; i++ {
		s = strconv.Itoa(i)

		for _, v := range s {
			nn = int(v - '0')
			if nn == 3 || nn == 6 || nn == 9 {
				sum++
			}
		}
	}

	Fprint(out, sum)
	// 입력 에러 방지
	_leftData, _ := ioutil.ReadAll(in)
	if _s := strings.TrimSpace(string(_leftData)); _s != "" {
		panic("읽지않은 데이터 발견：\n" + _s)
	}
}

// LETYLETYLETY
func main() { BOJ17614(os.Stdin, os.Stdout) }

// YTELYTELYTEL
