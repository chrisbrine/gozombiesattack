package zombies

import "fmt"

type Herb struct {
	// The position of the herb
	position 			Point
	// The drawing of the herb
	drawing 			Drawing
	// The health of the herb
	health 				int
	// The damage of the herb
	damage				int
	// the timer on the herb
	timer 				int
	// the herb damage for zombies
	zombieDamage 	int
	// the herb heal for zombies
	zombieHeal 		int
}

func (h *Herb) String() string {
	return fmt.Sprintf("Herb{position: %v, drawing: %v, health: %v, damage: %v, timer: %v, zombieDamage: %v, zombieHeal: %v}", h.position, h.drawing.String(), h.health, h.damage, h.timer, h.zombieDamage, h.zombieHeal)
}
