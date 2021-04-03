package main

import (
	"bufio"
	. "fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

func CFe(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	// single test case

	var a, b int
	Fscan(in, &a, &b)

	var board [][]int
	input := make([]string, a)
	board = make([][]int, a)
	for i := 0; i < a; i++ {
		board[i] = make([]int, b)
	}

	for i := 0; i < a; i++ {
		Fscan(in, &input[i])
		// Fprint(out, input[i], "\n")
		for j := 0; j < b; j++ {
			if input[i][j] == '*' {
				board[i][j] = 1
			} else {
				board[i][j] = 0
			}
		}
	}

	// Fprint(out, board)
	cx := 0
	cy := 0

	answer := 0

	for {
		if board[cy][cx] == 1 {
			answer += 1
		}

		// Fprint(out, cy, cx, "\n")
		// print(cy, cx)

		// dist_right := b - 1 - cx

		dist_right := 100
		if cx == b-1 {
			// right wall
		} else {
			// go right (cy)
			for i := cx + 1; i < b; i++ {
				if board[cy][i] == 1 {
					dist_right = i - cx
					break
				}
			}
		}

		// dist_bottom := a - 1 - cy

		dist_bottom := 100
		if cy == a-1 {
			// bottom wall
		} else {
			// see bottom (cx)
			for j := cy + 1; j < a; j++ {
				if board[j][cx] == 1 {
					dist_bottom = j - cy
					break
				}
			}
		}

		if dist_bottom == 100 && dist_right == 100 {
			// no berry
			// move 1 right
			if cx == b-1 {
				cy++
			} else {
				cx++
			}
		} else if dist_right <= dist_bottom {
			// go right (move cx)
			cx += dist_right
		} else {
			cy += dist_bottom
		}

		if cy >= a-1 && cx >= b-1 {
			break
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
func main() { CFe(os.Stdin, os.Stdout) }

// YTELYTELYTEL
