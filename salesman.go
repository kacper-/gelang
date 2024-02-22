package main

import (
	"fmt"
	"math"
)

type Points struct {
	x int16
	y int16
}

func p2p(from Points, to Points) (d float64) {
	xd := from.x - to.x
	yd := from.y - to.y
	d = float64(xd*xd) + float64(yd*yd)
	d = math.Sqrt(d)
	return
}

func dist(idx []int, points []Points) (sum float64) {
	sum = 0
	for i := 1; i < len(idx); i++ {
		sum += p2p(points[idx[i-1]], points[idx[i]])
	}
	sum += p2p(points[idx[0]], points[idx[len(idx)-1]])
	return
}

func main() {
	idx := [4]int{0, 1, 2, 3}

	points := [4]Points{
		Points{1, 1},
		Points{2, 1},
		Points{1, 2},
		Points{2, 2},
	}

	min := dist(idx[:], points[:])

	size := len(idx)
	p := make([]int, size+1)
	for i := 0; i < size+1; i++ {
		p[i] = i
	}
	for i := 1; i < size; {
		p[i]--
		j := 0
		if i%2 == 1 {
			j = p[i]
		}
		tmp := idx[j]
		idx[j] = idx[i]
		idx[i] = tmp
		d := dist(idx[:], points[:])
		if d < min {
			min = d
		}
		for i = 1; p[i] == 0; i++ {
			p[i] = i
		}
	}

	fmt.Println(min)
}
