package day_5

import "github.com/sebastiangeiger/advent-of-code-2021/common"

type Line struct {
	start Point
	end   Point
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
	if !line.IsHorizontal() && !line.IsVertical() {
		panic("Only works on horizontal/vertical lines!")
	}
	result := []Point{}
	if line.IsHorizontal() {
		xStart, xEnd := common.StartEnd(line.start.x, line.end.x)
		y := line.start.y
		for x := xStart; x <= xEnd; x++ {
			result = append(result, Point{x, y})
		}
	}
	if line.IsVertical() {
		yStart, yEnd := common.StartEnd(line.start.y, line.end.y)
		x := line.start.x
		for y := yStart; y <= yEnd; y++ {
			result = append(result, Point{x, y})
		}
	}
	return result
}
