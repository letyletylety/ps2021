package main

import (
	"bufio"
	. "fmt"
	"io"
	"io/ioutil"
	"strings"
	"time"
)

func BOJ16170(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	// single test case

	t := time.Now().UTC()
	year := t.Year()
	mon := int(t.Month())
	day := t.Day()

	Fprintln(out, year)
	Fprintf(out, "%02d\n", mon)
	Fprintf(out, "%02d\n", day)

	// 입력 에러 방지
	_leftData, _ := ioutil.ReadAll(in)
	if _s := strings.TrimSpace(string(_leftData)); _s != "" {
		panic("읽지않은 데이터 발견：\n" + _s)
	}
}

// TODO : 끝나고 주석 처리
// func main() { BOJ16170(os.Stdin, os.Stdout) }
