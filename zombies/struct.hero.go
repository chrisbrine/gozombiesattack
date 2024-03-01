package zombies

import "fmt"

type Hero struct {
	name						string
	position				Point
	drawing					Drawing
	weapon					Weapon
	weapon1 				Weapon
	weapon2 				Weapon
	currentWeapon	 	int
	herbs						int
	herbsMax  			int
	damageHerbs			int
	damageHerbsMax	int
	health					int
	healthMax				int
	score						int
}

func (h *Hero) String() string {
	return fmt.Sprintf("Hero{name: %v, position: %v, drawing: %v, weapon: %v, weapon1: %v, weapon2: %v, currentWeapon: %v, herbs: %v, herbsMax: %v, damageHerbs: %v, damageHerbsMax: %v, health: %v, healthMax: %v, score: %v}", h.name, h.position, h.drawing.String(), h.weapon.String(), h.weapon1.String(), h.weapon2.String(), h.currentWeapon, h.herbs, h.herbsMax, h.damageHerbs, h.damageHerbsMax, h.health, h.healthMax, h.score)
}
