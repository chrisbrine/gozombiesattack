package zombies

import "github.com/gdamore/tcell"

func BlendMarkRow(row []Mark, color tcell.Color) []Mark {
	for i := range row {
		row[i].color = color
	}
	return row
}

func CreateMark(char rune, color tcell.Color) Mark {
	return Mark{
		char:  char,
		color: color,
		item: false,
		direction: Point{},
	}
}

func (m *Mark) AddItem(item bool, direction Point) {
	m.item = item
	m.direction = direction
}

func NewDrawing(marks [][]Mark) Drawing {
	height := len(marks)
	width := len(marks[0])
	for i := 1; i < height; i++ {
		if len(marks[i]) > width {
			width = len(marks[i])
		}
	}
	return Drawing{
		drawing: marks,
		height:  height,
		width:   width,
	}
}