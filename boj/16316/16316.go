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

func BOJ16316(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	// single test case

	var n int
	Fscan(in, &n)

	var t string

	var answer = "makes sense"

	for i := 1; i <= n; i++ {
		Fscan(in, &t)

		if t == "mumble" {
			continue
		} else {
			tt, _ := strconv.Atoi(t)
			if tt != i {
				answer = "something is fishy"
			}
		}
	}

	Fprint(out, answer)

	// 입력 에러 방지
	_leftData, _ := ioutil.ReadAll(in)
	if _s := strings.TrimSpace(string(_leftData)); _s != "" {
		panic("읽지않은 데이터 발견：\n" + _s)
	}
}

// LETYLETYLETY
func main() { BOJ16316(os.Stdin, os.Stdout) }

// YTELYTELYTEL
