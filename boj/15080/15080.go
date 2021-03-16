package main

import (
	"bufio"
	. "fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
	"time"
)

func BOJ15080(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	input, _ := ioutil.ReadAll(in)

	a := strings.Split(string(input), "\n")

	layout := "15 : 4 : 5"
	t1, _ := time.Parse(layout, a[0])
	t2, _ := time.Parse(layout, a[1])

	// Fprint(out, t1, t2)
	ret := t2.Sub(t1).Seconds()

	if ret < 0 {
		Fprint(out, 86400+ret)
	} else {
		Fprint(out, ret)
	}

	// 입력 에러 방지
	_leftData, _ := ioutil.ReadAll(in)
	if _s := strings.TrimSpace(string(_leftData)); _s != "" {
		panic("읽지않은 데이터 발견：\n" + _s)
	}
}

// LETYLETYLETY
func main() { BOJ15080(os.Stdin, os.Stdout) }

// YTELYTELYTEL
