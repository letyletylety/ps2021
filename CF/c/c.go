package main

import (
	"bufio"
	. "fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

func CFc(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	// single test case

	var s string
	Fscan(in, &s)

	// a := strings.Split(s, "")

	b := []byte(s)

	// for k := range b {
	// 	Fprint(out, b[k]-'A', " ")
	// }

	if len(b) < 2 {
		Fprint(out, "YES")
	} else {
		result := "YES"
		pre := b[0] - 'A'
		cur := b[1] - 'A'
		for i := 2; i < len(b); i++ {
			if (pre+cur)%26 == b[i]-'A' {
				pre = cur
				cur = b[i] - 'A'
			} else {
				// Fprint(out, cur, "===")
				result = "NO"
				break
			}
		}

		Fprint(out, result)
	}

	// 입력 에러 방지
	_leftData, _ := ioutil.ReadAll(in)
	if _s := strings.TrimSpace(string(_leftData)); _s != "" {
		panic("읽지않은 데이터 발견：\n" + _s)
	}
}

// LETYLETYLETY
func main() { CFc(os.Stdin, os.Stdout) }

// YTELYTELYTEL
