package zombies

import (
	"fmt"

	"github.com/gdamore/tcell"
)

type Mark struct {
	// The character to draw
	char 						rune
	// The color of the character
	color 					tcell.Color
	// does this point have the item?
	item 						bool
	// direction of the item
	direction 			Point
}

func (m *Mark) String() string {
	return fmt.Sprintf("Mark{char: %v, color: %v, item: %v, direction: %v}", m.char, m.color, m.item, m.direction)
}

type Drawing struct {
	// the drawing of the object
	drawing 				[][]Mark
	width						int
	height					int
}

func (d *Drawing) String() string {
	drawingString := d.DrawingString()
	return fmt.Sprintf("Drawing{drawing: %v, width: %v, height: %v}", drawingString, d.width, d.height)
}

func (d *Drawing) DrawingString() string {
	str := ""
	for _, row := range d.drawing {
		for _, mark := range row {
			str += string(mark.char)
		}
		str += "\n"
	}
	return str
}
