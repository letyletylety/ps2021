package main

import (
	"bufio"
	. "fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

type uf struct {
	fa []int
}

func UnionFind(n int) uf {
	fa := make([]int, n)
	for i := range fa {
		fa[i] = i
	}
	return uf{fa}
}

func (u uf) find(x int) int {
	if u.fa[x] != x {
		u.fa[x] = u.find(u.fa[x])
	}
	return u.fa[x]
}

func (u uf) merge(from, to int) (isNewMerge bool) {
	x, y := u.find(from), u.find(to)
	if x == y {
		return false
	}
	u.fa[x] = y
	return true
}

func (u uf) same(x, y int) bool { return u.find(x) == u.find(y) }

func (u uf) countRoots(l, r int) (cnt int) {
	for i := l; i <= r; i++ {
		if u.find(i) == i {
			cnt++
		}
	}
	return
}

func CF1506E(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	// single test case

	var tt int
	Fscan(in, &tt)

	for tt > 0 {
		tt--

		var n int
		Fscan(in, &n)

		a := make([]int, n)
		pmin := make([]int, n)
		pmax := make([]int, n)

		check := make(map[int]bool, n)

		u := newUnionFind(n + 1)

		for i := 0; i < n; i++ {
			Fscan(in, &a[i])
			if i == 0 || a[i] != a[i-1] {
				check[a[i]] = true
				pmin[i] = a[i]
				pmax[i] = a[i]
				u.merge(a[i], a[i]-1)
			}
		}

		cur := 1
		for k, v := range pmin {
			if v == 0 {
				for ; check[cur]; cur++ {
				}
				pmin[k] = cur
				cur++
			}
		}

		pre := 0
		for k, v := range pmax {
			if v > 0 {
				pre = v
			} else {
				cur := u.find(pre)
				pmax[k] = cur
				u.merge(pre, cur-1)
			}
		}

		for k := range pmin {
			Fprint(out, pmin[k])
			Fprint(out, " ")
		}
		Fprint(out, "\n")
		for k := range pmax {
			Fprint(out, pmax[k])
			Fprint(out, " ")
		}
		Fprint(out, "\n")
	}

	// 입력 에러 방지
	_leftData, _ := ioutil.ReadAll(in)
	if _s := strings.TrimSpace(string(_leftData)); _s != "" {
		panic("읽지않은 데이터 발견：\n" + _s)
	}
}

// LETYLETYLETY
func main() { CF1506E(os.Stdin, os.Stdout) }

// YTELYTELYTEL
