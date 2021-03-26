package main

import (
	"bufio"
	. "fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

func BOJ4714(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	// single test case

	for {
		var n float64
		Fscan(in, &n)

		if n == -1.0 {
			break
		}

		Fprintf(out, "Objects weighing %.2f on Earth will weigh %.2f on the moon.\n", n, 0.167*n)
	}

	// 입력 에러 방지
	_leftData, _ := ioutil.ReadAll(in)
	if _s := strings.TrimSpace(string(_leftData)); _s != "" {
		panic("읽지않은 데이터 발견：\n" + _s)
	}
}

// LETYLETYLETY
func main() { BOJ4714(os.Stdin, os.Stdout) }

// YTELYTELYTEL
