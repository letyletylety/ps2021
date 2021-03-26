package main

import (
	"bufio"
	. "fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

func BOJ21185(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	// single test case
	var n int
	Fscan(in, &n)

	switch n {
	case 1:
		Fprint(out, "Either")
		break
	case 2:
		Fprint(out, "Odd")
		break
	default:
		if n%2 == 0 {
			if n%4 == 0 {

				Fprint(out, "Even")
			} else {
				Fprint(out, "Odd")
			}
		} else {
			Fprint(out, "Either")
		}
	}
	// 입력 에러 방지
	_leftData, _ := ioutil.ReadAll(in)
	if _s := strings.TrimSpace(string(_leftData)); _s != "" {
		panic("읽지않은 데이터 발견：\n" + _s)
	}
}

// LETYLETYLETY
func main() { BOJ21185(os.Stdin, os.Stdout) }

// YTELYTELYTEL
