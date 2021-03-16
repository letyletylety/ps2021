package main

import (
	"bufio"
	. "fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

func max(a, b int64) int64 {
	if a > b {
		return a
	}
	return b
}

func CFA(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n int
	var temp int64
	var limit int64
	var result int64
	Fscan(in, &n)

	// single test case
	cnt := make([]int64, 100001)
	mem := make([]int64, 100001)

	for i := 0; i < n; i++ {

		Fscan(in, &temp)
		cnt[temp]++

		if limit < temp {
			limit = temp
		}
	}

	mem[0] = 0
	mem[1] = cnt[1]
	result = mem[1]

	for i := int64(2); i <= limit; i++ {
		mem[i] = mem[i-1]
		mem[i] = max(mem[i-2]+cnt[i]*i, mem[i-1])
		if result < mem[i] {
			result = mem[i]
		}
	}

	Fprint(out, result)

	// 입력 에러 방지
	_leftData, _ := ioutil.ReadAll(in)
	if _s := strings.TrimSpace(string(_leftData)); _s != "" {
		panic("읽지않은 데이터 발견：\n" + _s)
	}
}

// LETYLETYLETY
func main() { CFA(os.Stdin, os.Stdout) }

// YTELYTELYTEL
