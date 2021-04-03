package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

type pair struct {
	x, y int
}

func BOJ9161(_r io.Reader, _w io.Writer) {
	// in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	// single test case

	for i := 10; i < 100; i++ {
		for j := 1; j < 10; j++ {
			for k := 10; k < 100; k++ {
				if i == k {
					continue
				}

				x := i*10 + j
				y := j*100 + k

				if x == y {
					continue
				}

				if x*k == y*i {
					Fprintf(out, "%d / %d = %d / %d\n", x, y, i, k)
				}
			}
		}
	}
}

// LETYLETYLETY
func main() { BOJ9161(os.Stdin, os.Stdout) }

// YTELYTELYTEL
