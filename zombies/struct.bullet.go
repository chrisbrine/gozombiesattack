package zombies

import "fmt"

type Bullet struct {
	// The position of the bullet
	position 			Point
	// The drawing of the bullet
	drawing 			Drawing
	// The damage of the bullet
	damage 				int
	// The speed of the bullet
	speed 				float64
}

func (b *Bullet) String() string {
	return fmt.Sprintf("Bullet{position: %v, drawing: %v, damage: %v, speed: %v}", b.position, b.drawing.String(), b.damage, b.speed)
}
