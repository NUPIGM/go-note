package main

import (
	"fmt"
	"os"
)

/**
 * Auto-generated code below aims at helping you parse
 * the standard input according to the problem statement.
 **/

func main() {
	// W: width of the building.
	// H: height of the building.
	var W, H int
	fmt.Scan(&W, &H)

	// N: maximum number of turns before game over.
	var N int
	fmt.Scan(&N)

	var X0, Y0 int
	fmt.Scan(&X0, &Y0)
	var x int = 0
	var y int = 0
	for {
		// bombDir: the direction of the bombs from batman's current location (U, UR, R, DR, D, DL, L or UL)
		var bombDir string
		fmt.Scan(&bombDir)

		// 第一象限
		if bombDir == "UR" || bombDir == "R" {
			if bombDir == "R" {
				r := W - X0
				x = count(r) + X0
			}
			r := W - X0
			x = count(r) + X0
			u := Y0
			y = count(u)
		}
		// 第二象限
		if bombDir == "UL" || bombDir == "U" {
			if bombDir == "U" {
				u := Y0
				y = count(u)
			}
			u := Y0
			y = count(u)
			r := X0
			x = count(r)
		}
		// 第三象限
		if bombDir == "DL" || bombDir == "L" {
			if bombDir == "L" {
				r := X0
				x = count(r)
			}
			r := X0
			x = count(r)
			u := H - Y0
			y = count(u) + Y0
		}
		// 第四象限
		if bombDir == "DR" || bombDir == "D" {
			if bombDir == "D" {
				u := H - Y0
				y = count(u) + Y0
			}
			u := H - Y0
			y = count(u) + Y0
			r := W - X0
			x = count(r) + X0
		}

		// fmt.Fprintln(os.Stderr, "Debug messages...")
		fmt.Fprintln(os.Stderr, N)
		fmt.Fprintln(os.Stderr, W, H)
		fmt.Fprintln(os.Stderr, X0, Y0)
		fmt.Fprintln(os.Stderr, bombDir)
		// the location of the next window Batman should jump to.
		fmt.Println(x, y)
		X0 = x
		Y0 = y
	}
}

func count(i int) (v int) {
	if i > 1 {
		v = i / 2
	} else if i == 1 {
		v = 0
	} else {
		v = i
	}

	return
}
