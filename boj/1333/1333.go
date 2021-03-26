package main

import (
	"bufio"
	. "fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

func BOJ1333(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	// single test case
	var n, l, d int
	Fscan(in, &n, &l, &d)

	song := make([]bool, 2000000)
	bell := make([]bool, 2000000)
	fulltime := (n-1)*(l+5) + l

	for t := 0; t < 2000000; t++ {
		if fulltime < t {
			break
		}

		if t%(l+5) < l {
			song[t] = true
		}
	}

	for t := 0; t < 2000000; t++ {
		if t%d == 0 {
			bell[t] = true
		}
	}

	for t := 0; t < 2000000; t++ {
		if song[t] == false && bell[t] == true {
			Fprint(out, t)
			break
		}
	}

	// 입력 에러 방지
	_leftData, _ := ioutil.ReadAll(in)
	if _s := strings.TrimSpace(string(_leftData)); _s != "" {
		panic("읽지않은 데이터 발견：\n" + _s)
	}
}

// LETYLETYLETY
func main() { BOJ1333(os.Stdin, os.Stdout) }

// YTELYTELYTEL
