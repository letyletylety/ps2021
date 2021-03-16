package main

import (
	"bufio"
	. "fmt"
	"io"
	"io/ioutil"
	
	"strings"
)

func BOJ8718(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	// single test case
	var x, k, answer int

	Fscan(in, &x, &k)
	x *= 1000
	k *= 1000

	answer = 7 * k
	if answer <= x {
	} else if answer/2 <= x {
		answer = answer / 2
	} else if answer/4 <= x {
		answer = answer / 4
	} else {
		answer = 0
	}

	Fprint(out, answer)

	// 입력 에러 방지
	_leftData, _ := ioutil.ReadAll(in)
	if _s := strings.TrimSpace(string(_leftData)); _s != "" {
		panic("읽지않은 데이터 발견：\n" + _s)
	}
}

/*
func main() { BOJ8718(os.Stdin, os.Stdout) }

*/
