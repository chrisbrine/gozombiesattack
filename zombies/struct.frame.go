package zombies

import (
	"fmt"

	"github.com/gdamore/tcell"
)

type Frame struct {
	// The width of the frame
	width 			int
	// The height of the frame
	height 			int
	// The top left point of the frame
	position 		Point
	// The character to draw the frame
	char 				rune
	// The color of the frame
	color 			tcell.Color
	// The color of the background
	bgColor 		tcell.Color
}

func (f *Frame) String() string {
	return fmt.Sprintf("Frame{width: %v, height: %v, position: %v, char: %v, color: %v, bgColor: %v}", f.width, f.height, f.position, f.char, f.color, f.bgColor)
}
