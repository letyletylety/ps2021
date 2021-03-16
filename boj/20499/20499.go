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

func BOJ20499(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var str string
	// single test case
	Fscan(in, &str)
	bb := strings.Split(str, "/")

	// Fscan(in, &k, &d, &a)
	k, _ := strconv.Atoi(bb[0])
	d, _ := strconv.Atoi(bb[1])
	a, _ := strconv.Atoi(bb[2])

	// Fprint(out, k, d, a)

	if k+a < d || d == 0 {
		Fprint(out, "hasu")
	} else {
		Fprint(out, "gosu")
	}

	// 입력 에러 방지
	_leftData, _ := ioutil.ReadAll(in)
	if _s := strings.TrimSpace(string(_leftData)); _s != "" {
		// panic("읽지않은 데이터 발견：\n" + _s)
	}
}

// LETYLETYLETY
func main() { BOJ20499(os.Stdin, os.Stdout) }

// YTELYTELYTEL
