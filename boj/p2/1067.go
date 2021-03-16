package main

import (
	"bufio"
	. "fmt"
	"io"
	"io/ioutil"
	"math"
	"math/bits"
	"os"
	"strings"
)

type fft struct {
	n               int
	omega, omegaInv []complex128
}

func newFFT(n int) *fft {
	omega := make([]complex128, n)
	omegaInv := make([]complex128, n)
	for i := range omega {
		sin, cos := math.Sincos(2 * math.Pi * float64(i) / float64(n))
		omega[i] = complex(cos, sin)
		omegaInv[i] = complex(cos, -sin)
	}
	return &fft{n, omega, omegaInv}
}

func (t *fft) transform(a, omega []complex128) {
	for i, j := 0, 0; i < t.n; i++ {
		if i > j {
			a[i], a[j] = a[j], a[i]
		}
		for l := t.n >> 1; ; l >>= 1 {
			j ^= l
			if j >= l {
				break
			}
		}
	}
	for l := 2; l <= t.n; l <<= 1 {
		m := l >> 1
		for st := 0; st < t.n; st += l {
			b := a[st:]
			for i := 0; i < m; i++ {
				d := omega[t.n/l*i] * b[m+i]
				b[m+i] = b[i] - d
				b[i] += d
			}
		}
	}
}

func (t *fft) dft(a []complex128) {
	t.transform(a, t.omega)
}

func (t *fft) idft(a []complex128) {
	t.transform(a, t.omegaInv)
	cn := complex(float64(t.n), 0)
	for i := range a {
		a[i] /= cn
	}
}

// convolution of A(x) B(x)
// 전체적으로 int64
func polyConvFFT(a, b []int64) []int64 {
	n, m := len(a), len(b)
	limit := 1 << bits.Len(uint(n+m-1))
	A := make([]complex128, limit)
	for i, v := range a {
		A[i] = complex(float64(v), 0)
	}
	B := make([]complex128, limit)
	for i, v := range b {
		B[i] = complex(float64(v), 0)
	}
	t := newFFT(limit)
	t.dft(A)
	t.dft(B)
	for i := range A {
		A[i] *= B[i]
	}
	t.idft(A)
	conv := make([]int64, n+m-1)
	for i := range conv {
		conv[i] = int64(math.Round(real(A[i]))) // % mod
	}
	return conv
}

func BOJ1067(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	// single test case
	var n int
	Fscan(in, &n)
	a := make([]int64, n+n)
	b := make([]int64, n)

	for i := 0; i < n; i++ {
		Fscan(in, &a[i])
		a[i+n] = a[i]
	}

	for i := n - 1; i >= 0; i-- {
		Fscan(in, &b[i])
	}

	c := polyConvFFT(a, b)

	answer := c[0]
	for _, cc := range c {
		if answer < cc {
			answer = cc
		}
	}
	Fprint(out, answer)

	// 입력 에러 방지
	_leftData, _ := ioutil.ReadAll(in)
	if _s := strings.TrimSpace(string(_leftData)); _s != "" {
		panic("읽지않은 데이터 발견：\n" + _s)
	}
}

// LETYLETYLETY
func main() { BOJ1067(os.Stdin, os.Stdout) }

// YTELYTELYTEL
