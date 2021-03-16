package main

import (
	"bufio"
	. "fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

func minmax(vals ...int) (int, int) {
	if len(vals) == 0 {
		return 0, 0
	}
	retmin := vals[0]
	retmax := vals[0]
	for _, val := range vals {
		if retmin > val {
			retmin = val
		}
		if retmax < val {
			retmax = val
		}
	}
	return retmin, retmax
}

func BOJ14173(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	// single test case
	var x1, y1, x2, y2 int
	var x3, y3, x4, y4 int
	Fscan(in, &x1, &y1, &x2, &y2)
	Fscan(in, &x3, &y3, &x4, &y4)

	var xl, xh, yl, yh int
	xl, _ = minmax(x1, x3)
	_, xh = minmax(x2, x4)
	yl, _ = minmax(y1, y3)
	_, yh = minmax(y2, y4)

	_, x := minmax((xh - xl), (yh - yl))

	Fprint(out, x*x)
	// 입력 에러 방지
	_leftData, _ := ioutil.ReadAll(in)
	if _s := strings.TrimSpace(string(_leftData)); _s != "" {
		panic("읽지않은 데이터 발견：\n" + _s)
	}
}

// LETYLETYLETY
func main() { BOJ14173(os.Stdin, os.Stdout) }

// YTELYTELYTEL
