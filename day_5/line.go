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
