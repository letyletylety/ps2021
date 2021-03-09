package main

import (
	"bufio"
	. "fmt"
	"io"
)

func BOJ6749(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	// problem solving
	var y int
	var m int
	Fscan(in, &y)
	Fscan(in, &m)

	var answer int

	answer = m + m - y
	Fprintln(out, answer)

	// 입력 에러 방지
	// _leftData, _ := ioutil.ReadAll(in)
	// if _s := strings.TrimSpace(string(_leftData)); _s != "" {
	// 	panic("읽지않은 데이터 발견：\n" + _s)
	// }
}

// func main() { BOJ6749(os.Stdin, os.Stdout) }
