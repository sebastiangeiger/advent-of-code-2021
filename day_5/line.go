package day_5

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
	// This is not right yet
	result := []Point{}
	result = append(result, line.start)
	result = append(result, line.end)
	return result
}
