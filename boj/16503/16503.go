package main

import (
	"bufio"
	. "fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

func BOJ16503(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	// single test case

	var k1, k2, k3 int
	var o1, o2 string

	Fscan(in, &k1, &o1, &k2, &o2, &k3)

	exp := func(k1, k2 int, op1 string) int {
		t := 0
		k1minus := false
		k2minus := false

		switch op1 {
		case "+":
			t = k1 + k2
			break
		case "-":
			t = k1 - k2
			break
		case "*":
			t = k1 * k2
			break
		case "/":
			if k1 < 0 {
				k1minus = true
				k1 = -k1
			}
			if k2 < 0 {
				k2minus = true
				k2 = -k2
			}
			t = k1 / k2
			if k1minus {
				k1 = -k1
				t = -t
			}

			if k2minus {
				k2 = -k2
				t = -t
			}
			break
		}

		// if k1minus && k2minus {
		// 	t = -t
		// }
		return t
	}

	sol := func(k1, k2, k3 int, o1, o2 string) int {
		k12 := exp(k1, k2, o1)
		return exp(k12, k3, o2)
	}

	sol2 := func(k1, k2, k3 int, o1, o2 string) int {
		k23 := exp(k2, k3, o2)
		return exp(k1, k23, o1)
	}
	ans := sol(k1, k2, k3, o1, o2)
	ans2 := sol2(k1, k2, k3, o1, o2)

	if ans < ans2 {
		Fprintln(out, ans)
		Fprintln(out, ans2)
	} else {
		Fprintln(out, ans2)
		Fprintln(out, ans)
	}

	// 입력 에러 방지
	_leftData, _ := ioutil.ReadAll(in)
	if _s := strings.TrimSpace(string(_leftData)); _s != "" {
		panic("읽지않은 데이터 발견：\n" + _s)
	}
}

// LETYLETYLETY
func main() { BOJ16503(os.Stdin, os.Stdout) }

// YTELYTELYTEL
