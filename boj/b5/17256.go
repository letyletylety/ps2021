package main

import (
	"bufio"
	. "fmt"
	"io"
	"io/ioutil"
	"strings"
)

func BOJ17256(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	// single test case
	var ax, ay, az int
	var cx, cy, cz int
	var bx, by, bz int
	Fscan(in, &ax, &ay, &az)
	Fscan(in, &cx, &cy, &cz)

	bx = cx - az
	by = cy / ay
	bz = cz - ax

	Fprint(out, bx, by, bz)

	// 입력 에러 방지
	_leftData, _ := ioutil.ReadAll(in)
	if _s := strings.TrimSpace(string(_leftData)); _s != "" {
		panic("읽지않은 데이터 발견：\n" + _s)
	}
}

// TODO : 끝나고 주석 처리
// func main() { BOJ17256(os.Stdin, os.Stdout) }
