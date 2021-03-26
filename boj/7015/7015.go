package main

import (
	"bufio"
	. "fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

func yesterday(y, m, d int) (int, int) {
	mm := m
	dd := d

	dd--
	if dd == 0 {
		mm--
		if y%3 > 0 && mm%2 == 0 {
			dd = 19
		} else {
			dd = 20
		}
	}

	return mm, dd
}

func BOJ7015(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	// single test case

	var ttt int
	Fscan(in, &ttt)

	for ttt > 0 {
		ttt--

		var y, m, d int
		Fscan(in, &y, &m, &d)

		answer := 0

		yy := 1000
		// mm := 1
		// dd := 1

		for yy > y+1 {
			yy--

			// div 3
			if yy%3 == 0 {
				answer += 200 // 200 days
			} else {
				answer += 195 // 2 4 6 8 10
			}
		}

		/// now (1 + y , 1, 1)
		// log.Print("answer: ", answer)

		answer++
		yy = y
		mm := 10

		dd := 19
		if yy%3 == 0 {
			dd++
		}

		// log.Print("yy: ", yy)
		for {
			if mm == m && dd == d {
				break
			}
			mm, dd = yesterday(yy, mm, dd)
			answer++
		}

		Fprint(out, answer, "\n")
	}

	// 입력 에러 방지
	_leftData, _ := ioutil.ReadAll(in)
	if _s := strings.TrimSpace(string(_leftData)); _s != "" {
		panic("읽지않은 데이터 발견：\n" + _s)
	}
}

// LETYLETYLETY
func main() { BOJ7015(os.Stdin, os.Stdout) }

// YTELYTELYTEL
