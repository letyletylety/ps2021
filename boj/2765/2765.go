package main

import (
	"bufio"
	. "fmt"
	"io"
	"io/ioutil"
	"math"
	"os"
	"strings"
)

func BOJ2765(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	// single test case

	datanumber := 0
	for {
		datanumber++

		dia := 0.0
		rot := 0
		ti := 0.0

		Fscan(in, &dia, &rot, &ti)

		if rot == 0 {
			break
		}

		Fprint(out, "Trip #", datanumber, ": ")
		wheel := math.Pi * float64(rot) * dia

		dist := wheel / 5280 / 12 // inch -> mile

		titi := dist / ti * 3600

		Fprintf(out, "%.2f ", dist)
		Fprintf(out, "%.2f\n", titi)
	}

	// 입력 에러 방지
	_leftData, _ := ioutil.ReadAll(in)
	if _s := strings.TrimSpace(string(_leftData)); _s != "" {
		panic("읽지않은 데이터 발견：\n" + _s)
	}
}

// LETYLETYLETY
func main() { BOJ2765(os.Stdin, os.Stdout) }

// YTELYTELYTEL
