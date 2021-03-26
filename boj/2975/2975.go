package main

import (
	"bufio"
	. "fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

func BOJ2975(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()
	// single test case

	for {
		start := 0
		task := "W"
		amount := 0

		Fscan(in, &start, &task, &amount)

		if start == 0 && task == "W" && amount == 0 {
			break
		}

		if task == "W" {
			start -= amount
		} else {
			start += amount
		}

		if start < -200 {
			Fprint(out, "Not allowed\n")
		} else {
			Fprint(out, start, "\n")
		}
	}

	// 입력 에러 방지
	_leftData, _ := ioutil.ReadAll(in)
	if _s := strings.TrimSpace(string(_leftData)); _s != "" {
		panic("읽지않은 데이터 발견：\n" + _s)
	}
}

// LETYLETYLETY
func main() { BOJ2975(os.Stdin, os.Stdout) }

// YTELYTELYTEL
