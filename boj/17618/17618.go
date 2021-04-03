package main

import (
	"bufio"
	. "fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

func BOJ17618(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	// single test case

	var n int
	Fscan(in, &n)

	amaz := make([]int, n+10)

	answer := 0
	mid := n / 2
	// mid1 := mid / 2
	// mid3 := mid + mid/2
	// midans1 := make(chan int)
	midans2 := make(chan int)
	// midans3 := make(chan int)

	for i := 0; i < 10; i++ {
		amaz[i] = i
	}
	for i := 10; i <= n; i++ {
		amaz[i] = amaz[i/10] + i%10
	}

	// go func() {
	// 	t := 0
	// 	for i := 1; i <= mid1; i++ {
	// 		if i%amaz[i] == 0 {
	// 			t++
	// 		}
	// 	}
	// 	midans1 <- t
	// }()
	go func() {
		t := 0
		for i := mid + 1; i <= n; i++ {
			if i%amaz[i] == 0 {
				t++
			}
		}
		midans2 <- t
	}()
	// go func() {
	// 	t := 0
	// 	for i := mid + 1; i <= mid3; i++ {
	// 		if i%amaz[i] == 0 {
	// 			t++
	// 		}
	// 	}
	// 	midans3 <- t
	// }()

	for i := 1; i <= mid; i++ {
		if i%amaz[i] == 0 {
			answer++
		}
	}
	// answer += <-midans1
	answer += <-midans2
	// answer += <-midans3

	Fprint(out, answer)

	// 입력 에러 방지
	_leftData, _ := ioutil.ReadAll(in)
	if _s := strings.TrimSpace(string(_leftData)); _s != "" {
		panic("읽지않은 데이터 발견：\n" + _s)
	}
}

// LETYLETYLETY
func main() { BOJ17618(os.Stdin, os.Stdout) }

// YTELYTELYTEL
