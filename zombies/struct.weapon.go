package zombies

import "fmt"

type Weapon struct {
	// The name of the weapon
	name 						string
	// The damage of the weapon
	damage 					int
	// The range of the weapon
	weaponRange 		int
	// The drawing of the weapon
	drawing 				Drawing
	// The speed of the weapon
	speed 					float64
	// The drawing of the bullet
	bulletDrawing 	Drawing
	// bullet position relative to drawing
	bulletPosition 	Point
}

func (w *Weapon) String() string {
	return fmt.Sprintf("Weapon{name: %v, damage: %v, weaponRange: %v, drawing: %v, speed: %v, bulletDrawing: %v, bulletPosition: %v}", w.name, w.damage, w.weaponRange, w.drawing.String(), w.speed, w.bulletDrawing.String(), w.bulletPosition)
}
