package main

import (
	"bufio"
	. "fmt"
	"io"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func CFd(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	// single test case

	var n int64
	Fscan(in, &n)
	var m int
	Fscan(in, &m)

	a := strconv.FormatInt(n, m)

	if len(a) == 1 {
		Fprint(out, "YES")
	} else {
		// Fprint(out, a)
		b := []byte(a)

		l := len(b)

		// result := "NO"
		// for i := 0; i < l; i++ {
		// 	if b[i] != b[l-1-i] {
		// 		result = "YES"
		// 		break
		// 	}
		// }

		// for i := 1; i+i < l; i++ {
		// 	if b[i-1] < b[i] {
		// 		result = "YES"
		// 		break
		// 	}
		// }

		result := "YES"
		for i := 1; i < l; i++ {
			if b[i-1] <= b[i] {
				result = "NO"
				break
			}
		}
		Fprint(out, result)
		// cc := make(map[byte]int)

		// for k := range b {
		// 	cc[b[k]] += 1
		// }

	}

	// 입력 에러 방지
	_leftData, _ := ioutil.ReadAll(in)
	if _s := strings.TrimSpace(string(_leftData)); _s != "" {
		panic("읽지않은 데이터 발견：\n" + _s)
	}
}

// LETYLETYLETY
func main() { CFd(os.Stdin, os.Stdout) }

// YTELYTELYTEL
