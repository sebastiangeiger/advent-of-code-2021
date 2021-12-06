package day_5

import (
	"math"
)

type Point struct {
	x int
	y int
}

func (p Point) AddVector(v Vector) Point {
	return Point{p.x + v.x, p.y + v.y}
}

func (p1 Point) ManhattanDistance(p2 Point) int {
	return int(math.Abs(float64(p1.x-p2.x)) + math.Abs(float64(p1.y-p2.y)))
}

func (p1 Point) DirectionTo(p2 Point) Vector {
	return Vector{normalize(p2.x - p1.x), normalize(p2.y - p1.y)}
}

func normalize(i int) int {
	if i == 0 {
		return 0
	} else if i > 0 {
		return 1
	} else {
		return -1
	}
}
