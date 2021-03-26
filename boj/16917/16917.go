package main

import (
	"bufio"
	. "fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

func max(a, b int64) int64 {
	if a < b {
		return b
	}
	return a
}

func min(a, b int64) int64 {
	if a > b {
		return b
	}
	return a
}

func BOJ16917(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	// single test case

	var a, b, c, x, y int64
	Fscan(in, &a, &b, &c, &x, &y)

	answer := int64(0)

	/// 치킨을 버린다는게 말이되나

	xy := max(x, y)

	//fried
	if x < y {
		answer += (y - x) * b
		y = x
	} else if x > y {
		answer += (x - y) * a
		x = y
	}

	// assert(x == y)
	if c+c > a+b {
		answer += a*x + b*x
	} else {
		answer += c*x + c*x
	}

	answer = min(xy*c+xy*c, answer)

	Fprint(out, answer)
	// 입력 에러 방지
	_leftData, _ := ioutil.ReadAll(in)
	if _s := strings.TrimSpace(string(_leftData)); _s != "" {
		panic("읽지않은 데이터 발견：\n" + _s)
	}
}

// LETYLETYLETY
func main() { BOJ16917(os.Stdin, os.Stdout) }

// YTELYTELYTEL
