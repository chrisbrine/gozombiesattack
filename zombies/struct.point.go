package zombies

import "fmt"

type Point struct {
	x, y int
}

func (p Point) String() string {
	return fmt.Sprintf("Point{x: %v, y: %v}", p.x, p.y)
}
