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

func root(s string) string {
	answer := int64(0)

	for _, v := range s {
		answer += int64(v - '0')
	}

	return strconv.FormatInt(answer, 10)
}

func BOJ6378(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	// single test case

	for {
		var n string
		Fscan(in, &n)

		if n == "0" {
			break
		}

		for len(n) > 1 {
			n = root(n)
		}
		Fprint(out, n, "\n")
	}

	// 입력 에러 방지
	_leftData, _ := ioutil.ReadAll(in)
	if _s := strings.TrimSpace(string(_leftData)); _s != "" {
		panic("읽지않은 데이터 발견：\n" + _s)
	}
}

// LETYLETYLETY
func main() { BOJ6378(os.Stdin, os.Stdout) }

// YTELYTELYTEL
