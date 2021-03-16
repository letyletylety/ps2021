package main

import (
	"bufio"
	"fmt"
	. "fmt"
	"io"
	"io/ioutil"
	
	"strings"
)

func BOJ11282(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	// single test case
	var d int
	Fscan(in, &d)

	s := fmt.Sprintf("%c", '가'+d-1)

	Fprintf(out, s)

	// 입력 에러 방지
	_leftData, _ := ioutil.ReadAll(in)
	if _s := strings.TrimSpace(string(_leftData)); _s != "" {
		panic("읽지않은 데이터 발견：\n" + _s)
	}
}

/*
func main() { BOJ11282(os.Stdin, os.Stdout) }

*/
