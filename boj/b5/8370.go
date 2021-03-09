package main

import (
	"bufio"
	. "fmt"
	"io"
	"io/ioutil"
	"strings"
)

func BOJ8370(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	// single test case
	var n, k, n2, k2 int
	var answer int
	Fscanf(in, "%d%d%d%d", &n, &k, &n2, &k2)
	answer = n*k + n2*k2
	Fprintf(out, "%d", answer)

	// 입력 에러 방지
	_leftData, _ := ioutil.ReadAll(in)
	if _s := strings.TrimSpace(string(_leftData)); _s != "" {
		panic("읽지않은 데이터 발견：\n" + _s)
	}
}

// func main() { BOJ8370(os.Stdin, os.Stdout) }
