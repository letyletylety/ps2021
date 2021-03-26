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

func BOJ5220(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	// single test case
	var tt int
	Fscan(in, &tt)

	for tt > 0 {
		tt--

		var a, b uint64
		Fscan(in, &a, &b)

		aa := strconv.FormatUint(a, 2)

		count := uint64(0)
		for _, v := range aa {
			if v == '1' {
				count++
			}
		}

		if count&(uint64(1)) == b {
			Fprint(out, "Valid\n")
		} else {
			Fprint(out, "Corrupt\n")
		}

	}

	// 입력 에러 방지
	_leftData, _ := ioutil.ReadAll(in)
	if _s := strings.TrimSpace(string(_leftData)); _s != "" {
		panic("읽지않은 데이터 발견：\n" + _s)
	}
}

// LETYLETYLETY
func main() { BOJ5220(os.Stdin, os.Stdout) }

// YTELYTELYTEL
