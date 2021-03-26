package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

func BOJ6609(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	for {
		var mogi, pun, larva, eggs, r, s, n int

		result, _ := in.ReadString(10)
		sline := result[:len(result)-1]

		Sscan(sline, &mogi, &pun, &larva, &eggs, &r, &s, &n)

		// eggs
		// e := 0

		var newmogi, newpun, newlarva int

		for i := 0; i < n; i++ {
			newmogi = (pun-1)/s + 1
			newpun = (larva-1)/r + 1
			newlarva = mogi * eggs

			pun = newpun
			mogi = newmogi
			larva = newlarva
		}

		Fprint(out, mogi, "\n")
	}

	// single test case
	// 입력 에러 방지
	// _leftData, _ := ioutil.ReadAll(in)
	// if _s := strings.TrimSpace(string(_leftData)); _s != "" {
	// 	panic("읽지않은 데이터 발견：\n" + _s)
	// }
}

// LETYLETYLETY
func main() { BOJ6609(os.Stdin, os.Stdout) }

// YTELYTELYTEL
