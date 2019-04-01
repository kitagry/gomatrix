package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"time"

	"github.com/buger/goterm"
	colorable "github.com/mattn/go-colorable"
)

var (
	LINE int = goterm.Height()
	COLS int = goterm.Width()

	matrix      [][]string
	spaces      []int
	length      []int
	stdOut      = bufio.NewWriter(colorable.NewColorableStdout())
	letterRunes = []rune("  abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ!@#$%^&*()1234567890")
)

func init() {
	matrix = make([][]string, COLS)
	for i := 0; i < COLS; i++ {
		matrix[i] = make([]string, LINE)
	}

	spaces = make([]int, COLS)
	length = make([]int, COLS)
	rand.Seed(time.Now().UnixNano())
}

func clear() {
	fmt.Fprint(stdOut, "\x1b[2")
	fmt.Fprint(stdOut, "\x1b[1;1H")
	fmt.Fprint(stdOut, "\x1b[32m")
	fmt.Fprint(stdOut, "\x1b[40m")
	stdOut.Flush()
}

func randString() string {
	b := make([]rune, 1)
	for k := range b {
		b[k] = letterRunes[rand.Intn(len(letterRunes))]
	}
	if string(b) == " " {
		return ""
	}
	return string(b)
}

func main() {
	for {
		clear()

		for i := 0; i < len(matrix); i += 2 {
			if matrix[i][0] == "" && spaces[i] > 0 {
				spaces[i]--
			} else if matrix[i][0] == "" {
				spaces[i] = rand.Intn(LINE-3) + 3
				length[i] = rand.Intn(LINE) + 1
				matrix[i][0] = randString()
			}

			j := 0
			y := 0
			firstColDone := false
			for j <= LINE {
				for j < LINE && matrix[i][j] == "" {
					j++
				}

				if j >= LINE {
					break
				}

				z := j
				y = 0
				for j < LINE && matrix[i][j] != "" {
					j++
					y++
				}

				if j >= LINE {
					matrix[i][z] = ""
					continue
				}

				matrix[i][j] = randString()

				if y > length[i] || firstColDone {
					matrix[i][z] = ""
					matrix[i][0] = ""
				}

				firstColDone = true
				j++
			}
		}

		matrixString := ""
		for j := 0; j < len(matrix[0]); j++ {
			for i := 0; i < len(matrix); i++ {
				if matrix[i][j] == "" {
					matrixString += " "
				} else {
					matrixString += matrix[i][j]
				}
			}

			if j != len(matrix[0])-1 {
				matrixString += "\n"
			}
		}

		fmt.Fprint(stdOut, matrixString)
		stdOut.Flush()

		time.Sleep(50 * time.Millisecond)
	}
}
