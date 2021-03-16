package main

import (
	"bufio"
	. "fmt"
	"io"
	"io/ioutil"
	
	"strconv"
	"strings"
)

func BOJ9524(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	// single test case
	found := "1723"
	var d int
	Fscan(in, &d)

	dd, _ := strconv.Atoi(string(found[d-1]))

	Fprint(out, dd)

	// 입력 에러 방지
	_leftData, _ := ioutil.ReadAll(in)
	if _s := strings.TrimSpace(string(_leftData)); _s != "" {
		panic("읽지않은 데이터 발견：\n" + _s)
	}
}

/*
func main() { BOJ9524(os.Stdin, os.Stdout) }

*/
