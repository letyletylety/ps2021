package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func count(a, b, c int) int {

	count := 0
	for {

		// log.Print("a: ", a)
		// log.Print("b: ", b)
		// log.Print("c: ", c)

		diff1 := b - a
		diff2 := c - b

		if diff1 < 2 && diff2 < 2 {
			break
		}

		count++
		if diff1 < diff2 {
			a = b
			b = c - 1
		} else {
			c = b
			b = a + 1
		}
	}
	return count
}

func BOJ11034(_r io.Reader, _w io.Writer) {
	in := bufio.NewScanner(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	// single test case

	for in.Scan() {
		text := in.Text()

		text = strings.TrimSpace(text)

		if len(text) == 0 {
			continue
		}

		// log.Print("text: ", text)
		a := strings.Split(text, " ")

		aa, _ := strconv.Atoi(a[0])
		bb, _ := strconv.Atoi(a[1])
		cc, _ := strconv.Atoi(a[2])

		Fprint(out, count(aa, bb, cc), "\n")
	}

	// 입력 에러 방지
	// _leftData, _ := ioutil.ReadAll(in)
	// if _s := strings.TrimSpace(string(_leftData)); _s != "" {
	// 	panic("읽지않은 데이터 발견：\n" + _s)
	// }
}

// LETYLETYLETY
func main() { BOJ11034(os.Stdin, os.Stdout) }

// YTELYTELYTEL
