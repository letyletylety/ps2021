package main

import (
	"bufio"
	. "fmt"
	"io"
	"io/ioutil"
	"math"
	
	"sort"
	"strings"
)

func BOJ13866(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	// single test case
	var a, b, c, d float64

	Fscan(in, &a, &b, &c, &d)

	e := []float64{a, b, c, d}

	sort.Float64s(e)

	Fprint(out,
		math.Abs((e[0]+e[3])-(e[1]+e[2])))

	// 입력 에러 방지
	_leftData, _ := ioutil.ReadAll(in)
	if _s := strings.TrimSpace(string(_leftData)); _s != "" {
		panic("읽지않은 데이터 발견：\n" + _s)
	}
}

/*
func main() { BOJ13866(os.Stdin, os.Stdout) }

*/
