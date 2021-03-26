package main

import (
	"bufio"
	. "fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

func rjatk(n int64) (string, int64) {
	sum := int64(0)

	i := int64(1)
	for ; i*i < n; i++ {
		if n%i == 0 {
			sum += i
			sum += n / i
		}
	}

	if i*i == n && n%i == 0 {
		sum += i
	}

	sum -= n

	if sum < n {
		return "DEFICIENT", sum
	} else if sum == n {
		return "PERFECT", sum
	} else {
		return "ABUNDANT", sum
	}
}

func BOJ6491(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	// single test case

	for {
		var n int

		Fscan(in, &n)

		if n == 0 {
			break
		}

		Fprint(out, n, " ")

		answer, _ := rjatk(int64(n))

		Fprint(out, answer, "\n")

	}

	// 입력 에러 방지
	_leftData, _ := ioutil.ReadAll(in)
	if _s := strings.TrimSpace(string(_leftData)); _s != "" {
		panic("읽지않은 데이터 발견：\n" + _s)
	}
}

// LETYLETYLETY
func main() { BOJ6491(os.Stdin, os.Stdout) }

// YTELYTELYTEL
