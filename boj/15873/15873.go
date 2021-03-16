package main

import (
	"bufio"
	. "fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

func BOJ15873(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	// single test case
	var input string

	Fscan(in, &input)

	l := len(input)

	if l == 2 {
		Fprintf(out, "%d", input[0]+input[1]-'0'-'0')
	} else if l == 3 {
		if input[2] == '0' {
			Fprintf(out, "%d", input[0]+10-'0')
		} else {
			Fprintf(out, "%d", 10+input[2]-'0')
		}
	} else {
		Fprint(out, 20)
	}

	// 입력 에러 방지
	_leftData, _ := ioutil.ReadAll(in)
	if _s := strings.TrimSpace(string(_leftData)); _s != "" {
		panic("읽지않은 데이터 발견：\n" + _s)
	}
}

// LETYLETYLETY
func main() { BOJ15873(os.Stdin, os.Stdout) }

// YTELYTELYTEL
