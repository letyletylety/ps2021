package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

func BOJ4909(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	// single test case

	for {
		sum := 0
		max := -1
		min := 11

		a := make([]int, 6)
		for i := 0; i < 6; i++ {
			Fscan(in, &a[i])
		}

		for i := 0; i < 6; i++ {
			sum += a[i]

			if max < a[i] {
				max = a[i]
			}

			if min > a[i] {
				min = a[i]
			}
		}

		if sum == 0 {
			break
		}
		sum -= max
		sum -= min

		i := sum / 4
		r := sum % 4

		if r == 0 {
			Fprint(out, i, "\n")
		} else {
			Fprint(out, i, ".")

			switch r {
			case 1:
				Fprint(out, 25)
				break
			case 2:
				Fprint(out, 5)
				break
			case 3:
				Fprint(out, 75)
				break
			}

			Fprint(out, "\n")
		}
	}

	// // 입력 에러 방지
	// _leftData, _ := ioutil.ReadAll(in)
	// if _s := strings.TrimSpace(string(_leftData)); _s != "" {
	// 	panic("읽지않은 데이터 발견：\n" + _s)
	// }
}

// LETYLETYLETY
func main() { BOJ4909(os.Stdin, os.Stdout) }

// YTELYTELYTEL
