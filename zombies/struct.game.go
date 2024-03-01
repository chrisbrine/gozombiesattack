package zombies

import "fmt"

type Game struct {
	// The hero of the game
	hero Hero
	// The zombies of the game
	zombies []Zombie
	// The weapons of the game
	weapons []Weapon
	// The bullets
	bullets []Bullet
	// The herbs of the game
	herbs []Herb
	// The frame of the game
	frame Frame
	// The settings
	settings Settings
	// Game timer for how long game is running in seconds
	timer int
	timerFloat float64
	// The status of the game
	status Status
	// Debug String
	debugString string
}

func (g *Game) String() string {
	return fmt.Sprintf("Game{hero: %v, zombies: %v, weapons: %v, bullets: %v, herbs: %v, frame: %v, settings: %v, timer: %v, status: %v, debugString: %v}", g.hero, g.zombies, g.weapons, g.bullets, g.herbs, g.frame, g.settings, g.timer, g.status, g.debugString)
}
