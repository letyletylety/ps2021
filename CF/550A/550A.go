package main

import (
	"bufio"
	. "fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

func CF550A(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	// single test case
	var s string
	Fscan(in, &s)

	l := len(s)

	foundAnswer := false
	findAB := false
	findBA := false

	j := l

	for i := 0; i < l-1; i++ {
		if s[i] == 'A' && s[i+1] == 'B' {
			findAB = true
			j = i + 2
			break
		}
	}
	for i := j; i < l-1; i++ {
		if s[i] == 'B' && s[i+1] == 'A' {
			findBA = true
			break
		}
	}

	if findAB && findBA {
		foundAnswer = true
	}

	findAB = false
	findBA = false
	j = l
	for i := 0; i < l-1; i++ {
		if s[i] == 'B' && s[i+1] == 'A' {
			findAB = true
			j = i + 2
			break
		}
	}
	for i := j; i < l-1; i++ {
		if s[i] == 'A' && s[i+1] == 'B' {
			findBA = true
			break
		}
	}
	if findAB && findBA {
		foundAnswer = true
	}

	if foundAnswer {
		Fprint(out, "YES")
	} else {
		Fprint(out, "NO")
	}

	// 입력 에러 방지
	_leftData, _ := ioutil.ReadAll(in)
	if _s := strings.TrimSpace(string(_leftData)); _s != "" {
		panic("읽지않은 데이터 발견：\n" + _s)
	}
}

// LETYLETYLETY
func main() { CF550A(os.Stdin, os.Stdout) }

// YTELYTELYTEL
