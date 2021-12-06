package day_5

import (
	"github.com/sebastiangeiger/advent-of-code-2021/common"
)

type Line struct {
	start Point
	end   Point
}

type Vector struct {
	x int
	y int
}

func (line *Line) IsHorizontal() bool {
	return line.start.y == line.end.y
}

func (line *Line) IsVertical() bool {
	return line.start.x == line.end.x
}

func (line *Line) IsDiagonal() bool {
	xStart, xEnd := common.StartEnd(line.start.x, line.end.x)
	yStart, yEnd := common.StartEnd(line.start.y, line.end.y)
	return xEnd-xStart == yEnd-yStart
}

func (line *Line) ToPoints() []Point {
	if !line.IsHorizontal() && !line.IsVertical() && !line.IsDiagonal() {
		panic("Only works on horizontal/vertical/diagnoal lines!")
	}
	result := []Point{}
	vector := line.start.DirectionTo(line.end)
	currentPoint := line.start
	for line.end.ManhattanDistance(currentPoint) > 0 {
		result = append(result, currentPoint)
		currentPoint = currentPoint.AddVector(vector)
	}
	result = append(result, line.end)
	return result
}
