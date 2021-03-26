package main

import (
	"bufio"
	. "fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

func isPrime(n int) bool {

	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}

	return true
}

func BOJ7789(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	// single test case

	var a, b int
	Fscan(in, &a, &b)

	b = b*1000000 + a

	if isPrime(a) && isPrime(b) {
		Fprint(out, "Yes")
	} else {
		Fprint(out, "No")
	}

	// 입력 에러 방지
	_leftData, _ := ioutil.ReadAll(in)
	if _s := strings.TrimSpace(string(_leftData)); _s != "" {
		panic("읽지않은 데이터 발견：\n" + _s)
	}
}

// LETYLETYLETY
func main() { BOJ7789(os.Stdin, os.Stdout) }

// YTELYTELYTEL
