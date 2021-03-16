package main

import (
	"bufio"
	. "fmt"
	"io"
	"io/ioutil"
	
	"strings"
)

func BOJ14065(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	// single test case
	onegal := 3.785411784 // l
	mile := 1609.344      // m

	// mile /gal -> l / 100km
	var x float64
	Fscan(in, &x)

	x = onegal * 100000 / x / mile

	Fprint(out, x)
	// Fprintf(out, "%.6f", x)
	// 입력 에러 방지
	_leftData, _ := ioutil.ReadAll(in)
	if _s := strings.TrimSpace(string(_leftData)); _s != "" {
		panic("읽지않은 데이터 발견：\n" + _s)
	}
}

/*
func main() { BOJ14065(os.Stdin, os.Stdout) }

*/
