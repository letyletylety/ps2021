package main

import (
	"bufio"
	. "fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

func BOJ4655(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	// single test case
	for {
		var tt float64
		Fscan(in, &tt)
		if tt == 0.00 {
			break
		}

		sum := float64(0.00)
		answer := 1

		for {
			answer++
			sum += 1.00 / float64(answer)

			if sum >= tt {
				break
			}
		}

		Fprint(out, answer-1, " card(s)\n")
	}

	// 입력 에러 방지
	_leftData, _ := ioutil.ReadAll(in)
	if _s := strings.TrimSpace(string(_leftData)); _s != "" {
		panic("읽지않은 데이터 발견：\n" + _s)
	}
}

// LETYLETYLETY
func main() { BOJ4655(os.Stdin, os.Stdout) }

// YTELYTELYTEL
