package main

import (
	"bufio"
	. "fmt"
	"io"
	"io/ioutil"
	"strings"
)

func BOJ6763(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	// single test case

	var limit, speed int

	Fscan(in, &limit, &speed)
	radar := speed - limit
	var fine int

	if radar < 1 {
		Fprint(out, "Congratulations, you are within the speed limit!\n")
		return
	} else if radar < 21 {
		fine = 100
	} else if radar < 31 {
		fine = 270
	} else {
		fine = 500
	}

	Fprintf(out, "You are speeding and your fine is $%d.\n", fine)

	// 입력 에러 방지
	_leftData, _ := ioutil.ReadAll(in)
	if _s := strings.TrimSpace(string(_leftData)); _s != "" {
		panic("읽지않은 데이터 발견：\n" + _s)
	}
}

/*
func main() { BOJ6763(os.Stdin, os.Stdout) }

*/
