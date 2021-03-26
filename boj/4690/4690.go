package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

func BOJ4690(_r io.Reader, _w io.Writer) {
	// in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	// single test case
	// var n int
	// Fscan(in, &n)

	for a := 6; a <= 100; a++ {
		for b := 2; b < a; b++ {
			for c := b; c < a; c++ {
				for d := c; d < a; d++ {

					if a*a*a == b*b*b+c*c*c+d*d*d {
						Fprintf(out, "Cube = %d, Triple = (%d,%d,%d)\n", a, b, c, d)
					} else if a*a*a < b*b*b+c*c*c+d*d*d {
						break
					}
				}
			}
		}
	}

	// 입력 에러 방지
	// _leftData, _ := ioutil.ReadAll(in)
	// if _s := strings.TrimSpace(string(_leftData)); _s != "" {
	// 	panic("읽지않은 데이터 발견：\n" + _s)
	// }
}

// LETYLETYLETY
func main() { BOJ4690(os.Stdin, os.Stdout) }

// YTELYTELYTEL
