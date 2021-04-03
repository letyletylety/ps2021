package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

func CFa(_r io.Reader, _w io.Writer) {
	in := bufio.NewScanner(_r)
	out := bufio.NewWriter(_w)

	// single test case

	for in.Scan() {
		_ = in.Text()
		Fprintln(out, "NO")
		out.Flush()
	}

	// 입력 에러 방지
}

// LETYLETYLETY
func main() { CFa(os.Stdin, os.Stdout) }

// YTELYTELYTEL
