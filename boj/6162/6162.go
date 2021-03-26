package main

import (
	"bufio"
	. "fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

func BOJ6162(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()
	// single test case

	var k int
	Fscan(in, &k)

	for testCase := 1; testCase <= k; testCase++ {
		Fprint(out, "Data Set ", testCase, ":\n")

		var e int
		Fscan(in, &e)
		var a int
		Fscan(in, &a)

		if e <= a {
			// no rkana
			Fprint(out, "no drought")
		} else if e <= 5*a {
			// rkana
			Fprint(out, "drought")
		} else if e <= 25*a {
			// mega
			Fprint(out, "mega drought")
		} else if e <= 125*a {
			// megamega
			Fprint(out, "mega mega drought")
		} else {
			factor := 625
			count := 3

			for factor*a < e {
				count++
				factor *= 5
			}

			for i := 0; i < count; i++ {
				Fprint(out, "mega ")
			}
			Fprint(out, "drought")
		}

		Fprint(out, "\n\n")
	}
	// 입력 에러 방지
	_leftData, _ := ioutil.ReadAll(in)
	if _s := strings.TrimSpace(string(_leftData)); _s != "" {
		panic("읽지않은 데이터 발견：\n" + _s)
	}
}

// LETYLETYLETY
func main() { BOJ6162(os.Stdin, os.Stdout) }

// YTELYTELYTEL
