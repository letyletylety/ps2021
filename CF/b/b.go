package main

import (
	"bufio"
	. "fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

func CFb(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	// single test case

	var n int
	Fscan(in, &n)

	sum := 0
	for i := 0; i < 10000; i++ {
		sum = 0
		for n > 0 {
			sum += n % 10
			n /= 10
		}

		if sum < 10 {
			break
		}
		n = sum
	}

	Fprint(out, sum)
	// nn := math.Sqrt(float64(n))
	// Fprint(out, int(nn))

	// 입력 에러 방지
	_leftData, _ := ioutil.ReadAll(in)
	if _s := strings.TrimSpace(string(_leftData)); _s != "" {
		panic("읽지않은 데이터 발견：\n" + _s)
	}
}

// LETYLETYLETY
func main() { CFb(os.Stdin, os.Stdout) }

// YTELYTELYTEL
