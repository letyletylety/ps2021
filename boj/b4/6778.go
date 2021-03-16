package main

import (
	"bufio"
	. "fmt"
	"io"
	"io/ioutil"
	
	"strings"
)

func BOJ6778(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	// single test case
	var a, e int
	Fscan(in, &a, &e)
	if a >= 3 && e <= 4 {
		Fprint(out, "TroyMartian\n")
	}
	if a <= 6 && e >= 2 {
		Fprint(out, "VladSaturnian\n")
	}
	if a <= 2 && e <= 3 {
		Fprint(out, "GraemeMercurian\n")
	}

	// 입력 에러 방지
	_leftData, _ := ioutil.ReadAll(in)
	if _s := strings.TrimSpace(string(_leftData)); _s != "" {
		panic("읽지않은 데이터 발견：\n" + _s)
	}
}

/*
func main() { BOJ6778(os.Stdin, os.Stdout) }

*/
