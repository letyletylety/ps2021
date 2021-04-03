package main

import (
	"bufio"
	. "fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

func CF708e(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	// single test case
	primes := make([]int, 0)
	mind := make([]int, 1e7+1)

	for i := 2; i <= 1e7+1; i++ {
		if mind[i] == 0 {
			primes = append(primes, i)
			mind[i] = i
		}

		for k := range primes {
			if primes[k] > mind[i] || primes[k]*i > 1e7 {
				break
			}
			mind[primes[k]*i] = primes[k]
		}
	}

	var tt int
	Fscan(in, &tt)

	for tt > 0 {
		tt--
		var n, k int
		Fscan(in, &n, &k)
		a := make([]int, n, 1)

		for i := 0; i < n; i++ {
			var x int
			Fscan(in, &x)
			cnt := 0
			last := 0
			for x > 1 {
				p := mind[x]
				if last == p {
					cnt++
				} else {
					if cnt%2 == 1 {
						a[i] *= last
					}
					last = p
					cnt = 1
				}
				x /= p
			}
			if cnt%2 == 1 {
				a[i] *= last
			}
		}

		L := 0
		ans := 1
		last := make(map[int]int)
		for i := 0; i < ln; i++ {
			if last[a[i]] {

			}

		}
	}

	// 입력 에러 방지
	_leftData, _ := ioutil.ReadAll(in)
	if _s := strings.TrimSpace(string(_leftData)); _s != "" {
		panic("읽지않은 데이터 발견：\n" + _s)
	}
}

// LETYLETYLETY
func main() { CF708e(os.Stdin, os.Stdout) }

// YTELYTELYTEL
