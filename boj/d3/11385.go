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

type fft11385 struct {
	n               int
	omega, omegaInv []complex128
}

func newFFT11385(n int) *fft11385 {
	omega := make([]complex128, n)
	omegaInv := make([]complex128, n)
	for i := range omega {
		sin, cos := math.Sincos(2 * math.Pi * float64(i) / float64(n))
		omega[i] = complex(cos, sin)
		omegaInv[i] = complex(cos, -sin)
	}
	return &fft11385{n, omega, omegaInv}
}

func (t *fft11385) transform11385(a, omega []complex128) {
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

func (t *fft11385) dft11385(a []complex128) {
	t.transform11385(a, t.omega)
}

func (t *fft11385) idft11385(a []complex128) {
	t.transform11385(a, t.omegaInv)
	cn := complex(float64(t.n), 0)
	for i := range a {
		a[i] /= cn
	}
}

// convolution of A(x) B(x)
// 전체적으로 int64
func polyConvFFT11385(a, b []int64) []int64 {
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
	t := newFFT11385(limit)
	t.dft11385(A)
	t.dft11385(B)
	for i := range A {
		A[i] *= B[i]
	}
	t.idft11385(A)
	conv := make([]int64, n+m-1)
	for i := range conv {
		conv[i] = int64(math.Round(real(A[i]))) // % mod
	}
	return conv
}

func BOJ11385(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	// single test case
	var n, m int
	Fscan(in, &n, &m)
	n++
	m++

	a := make([]int64, n)
	b := make([]int64, m)

	for i := 0; i < n; i++ {
		Fscan(in, &a[i])
	}
	for i := 0; i < m; i++ {
		Fscan(in, &b[i])
	}

	c := polyConvFFT11385(a, b)

	answer := int64(0)

	for _, cc := range c {
		answer ^= cc
		// Fprintf(out, "%d\n", cc)
	}

	Fprint(out, answer)

	// 입력 에러 방지
	_leftData, _ := ioutil.ReadAll(in)
	if _s := strings.TrimSpace(string(_leftData)); _s != "" {
		panic("읽지않은 데이터 발견：\n" + _s)
	}
}

// LETYLETYLETY
func main() { BOJ11385(os.Stdin, os.Stdout) }

// YTELYTELYTEL
