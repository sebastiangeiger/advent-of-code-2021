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

func (line *Line) ToPoints() []Point {
	if !line.IsHorizontal() && !line.IsVertical() {
		panic("Only works on horizontal/vertical lines!")
	}
	result := []Point{}
	if line.IsHorizontal() {
		xStart := common.Min(line.start.x, line.end.x)
		xEnd := common.Max(line.start.x, line.end.x)
		y := line.start.y
		for x := xStart; x <= xEnd; x++ {
			result = append(result, Point{x, y})
		}
	}
	if line.IsVertical() {
		yStart := common.Min(line.start.y, line.end.y)
		yEnd := common.Max(line.start.y, line.end.y)
		x := line.start.x
		for y := yStart; y <= yEnd; y++ {
			result = append(result, Point{x, y})
		}
	}
	return result
}
