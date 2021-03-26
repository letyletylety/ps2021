package main

import (
	"bufio"
	. "fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

func v(n int) string {

	sum := 0
	for i := 1; i <= n; i++ {
		if n%i == 0 {
			sum += i
		}
	}

	sum -= n

	ret := Sprintf("%d is a", n)

	if sum > n {
		ret += "n abundant number."
	} else if sum < n {
		ret += " deficient number."
	} else {
		ret += " perfect number."
	}
	return ret
}

func BOJ6975(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	// single test case

	var tt int
	Fscan(in, &tt)
	for tt > 0 {
		tt--

		var n int
		Fscan(in, &n)

		Fprint(out, v(n))
		Fprint(out, "\n")

		if tt > 0 {
			Fprint(out, "\n")
		}
	}

	// 입력 에러 방지
	_leftData, _ := ioutil.ReadAll(in)
	if _s := strings.TrimSpace(string(_leftData)); _s != "" {
		panic("읽지않은 데이터 발견：\n" + _s)
	}
}

// LETYLETYLETY
func main() { BOJ6975(os.Stdin, os.Stdout) }

// YTELYTELYTEL
