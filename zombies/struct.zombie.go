package zombies

import "fmt"

type Zombie struct {
	// The position of the zombie
	position Point
	// The drawing of the zombie
	drawing Drawing
	// The health of the zombie
	health int
	// The damage of the zombie
	damage int
	// The speed of the zombie
	speed float64
	// Bite Timer starts at maxBiteTimer
	biteTimer int
	// The max bite timer
	maxBiteTimer int
}

func (z *Zombie) String() string {
	return fmt.Sprintf("Zombie{position: %v, drawing: %v, health: %v, damage: %v, speed: %v, biteTimer: %v, maxBiteTimer: %v}", z.position, z.drawing.String(), z.health, z.damage, z.speed, z.biteTimer, z.maxBiteTimer)
}
