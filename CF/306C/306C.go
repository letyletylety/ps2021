package main

import (
	"bufio"
	. "fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

func CF306C(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	// single test case
	var s string
	Fscan(in, &s)

	slen := len(s)

	var a, b, c int
	for i := 0; i < slen; i++ {
		a = int(s[i] - '0')

		if a%8 == 0 {
			Fprint(out, "YES\n")
			Fprint(out, a)
			return
		}

		for j := i + 1; j < slen; j++ {
			b = int(s[j] - '0')

			if (10*a+b)%8 == 0 {
				Fprint(out, "YES\n")
				Fprint(out, 10*a+b)
				return
			}
			for k := j + 1; k < slen; k++ {
				c = int(s[k] - '0')

				if (100*a+10*b+c)%8 == 0 {
					Fprint(out, "YES\n")
					Fprint(out, (100*a + 10*b + c))
					return
				}

			}
		}
	}

	Fprint(out, "NO")

	// 입력 에러 방지
	_leftData, _ := ioutil.ReadAll(in)
	if _s := strings.TrimSpace(string(_leftData)); _s != "" {
		panic("읽지않은 데이터 발견：\n" + _s)
	}
}

// LETYLETYLETY
func main() { CF306C(os.Stdin, os.Stdout) }

// YTELYTELYTEL
