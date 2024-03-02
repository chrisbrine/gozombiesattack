package zombies

import (
	"fmt"
	// "math/rand"

	"github.com/gdamore/tcell"
)

type Start struct {
	// Starting name
	name 										string
	// Starting health of the hero
	health 									int
	// Starting max health of the hero
	healthMax 							int
	// Starting weapon of the hero (1 or 2)
	weapon 									int
	// Starting herbs of the hero
	herbs 									int
	// Max herbs of the hero
	herbsMax 								int
	// Starting damage herbs of the hero
	damageHerbs 						int
	// Max damage herbs of the hero
	damageHerbsMax 					int
	// Starting zombies on screen
	zombies 								int
	// Max health of the zombie
	zombieMaxHealth 				int
	// Min health of the zombie
	zombieMinHealth 				int
	// Max damage of zombie
	zombieMaxDamage 				int
	// Min damage of zombie
	zombieMinDamage 				int
	// Max speed of zombie
	zombieMaxSpeed					float64
	// Min speed of zombie
	zombieMinSpeed 					float64
}

func (s *Start) String() string {
	return fmt.Sprintf("Start{name: %v, health: %v, healthMax: %v, weapon: %v, herbs: %v, herbsMax: %v, damageHerbs: %v, damageHerbsMax: %v, zombies: %v, zombieMaxHealth: %v, zombieMinHealth: %v, zombieMaxDamage: %v, zombieMinDamage: %v, zombieMaxSpeed: %v, zombieMinSpeed: %v}", s.name, s.health, s.healthMax, s.weapon, s.herbs, s.herbsMax, s.damageHerbs, s.damageHerbsMax, s.zombies, s.zombieMaxHealth, s.zombieMinHealth, s.zombieMaxDamage, s.zombieMinDamage, s.zombieMaxSpeed, s.zombieMinSpeed)
}

type Options struct {
	// Starting herbs on screen
	herbs 									int
	// Max herbs on screen
	herbsMax 								int
	// Herb heals this much if healing herb max
	healingHerbMax 					int
	// Herb heals this much if healing herb min
	healingHerbMin 					int
	// Minimum timer for herb to last for
	herbTimerMin 						int
	// Maximum timer for herb to last for
	herbTimerMax 						int
	// Herb damages this much if damage herb max
	damageHerbMax 					int
	// Herb damages this much if damage herb min
	damageHerbMin 					int
	// Increase herbs every x points
	herbsIncRate 						int
	// Change herb is healing herb %
	healingHerbPercent 			int
	// Change herb is damage herb %
	damageHerbPercent 			int
	// Change herb is zombie damage herb %
	zombieDamageHerbPercent int
	// Change herb is zombie heal herb %
	zombieHealHerbPercent 	int
	// Increase zombies every x points
	zombiesIncRate 					int
	// Max zombies on screen
	zombiesMax 							int
	// Increase zombie health every x points
	zombieHealthIncRate 		int
	// Max health inc rate for zombies
	zombieHealthMaxRate 		int
	// Increase zombie speed every x points
	zombieSpeedIncRate 			float64
	// Max speed inc rate
	zombieSpeedMaxRate 			float64
	// Increase zombie damage every x points
	zombieDamageIncRate 		int
	// Max damage inc rate for zombies
	zombieDamageMaxRate 		int
	// Bite timer min
	zombieBiteTimerMin 			int
	// Bite timer max
	zombieBiteTimerMax 			int
	// Zombie kill points
	zombieKillPoints 				int
	// Zombie shoot points
	zombieShootPoints 			int
	// Survival time points
	survivalTimePoints 			int
	// Time to get survival points in seconds
	survivalTime 						int
	// The counter to count to survivalTime
	survivalTimeCounter 		int
	// Herb get points
	herbGetPoints 					int
	// Gun shoot loss points
	gunShootLossPoints 			int
	// Game Speed
	gameSpeed 							int
	// Frame Character
	frameCharacter 					rune
	// Frame Color
	frameColor  						tcell.Color
	// Frame Background Color
	frameBackgroundColor 		tcell.Color
	// Background Color
	backgroundColor 				tcell.Color
	// Text Color
	textColor 							tcell.Color
	// Zombie Move Counter
	zombieMoveCounter 			int
	// Zombie Move Counter Max
	zombieMoveCounterMax		int
	// Bullet Move Counter
	bulletMoveCounter 			int
	// Bullet Move Counter Max
	bulletMoveCounterMax		int
}

func (o *Options) String() string {
	return fmt.Sprintf("Options{herbs: %v, healingHerbMax: %v, healingHerbMin: %v, herbTimerMin: %v, herbTimerMax: %v, damageHerbMax: %v, damageHerbMin: %v, herbsIncRate: %v, healingHerbPercent: %v, damageHerbPercent: %v, zombieDamageHerbPercent: %v, zombieHealHerbPercent: %v, zombiesIncRate: %v, zombieHealthIncRate: %v, zombieSpeedIncRate: %v, zombieDamageIncRate: %v, zombieBiteTimerMin: %v, zombieBiteTimerMax: %v, zombieKillPoints: %v, zombieShootPoints: %v, survivalTimePoints: %v, survivalTime: %v, herbGetPoints: %v, gunShootLossPoints: %v, gameSpeed: %v, frameCharacter: %v, frameColor: %v, frameBackgroundColor: %v, backgroundColor: %v, textColor: %v, zombieMoveCounter: %v, zombieMoveCounterMax: %v, bulletMoveCounter: %v, bulletMoveCounterMax: %v}", o.herbs, o.healingHerbMax, o.healingHerbMin, o.herbTimerMin, o.herbTimerMax, o.damageHerbMax, o.damageHerbMin, o.herbsIncRate, o.healingHerbPercent, o.damageHerbPercent, o.zombieDamageHerbPercent, o.zombieHealHerbPercent, o.zombiesIncRate, o.zombieHealthIncRate, o.zombieSpeedIncRate, o.zombieDamageIncRate, o.zombieBiteTimerMin, o.zombieBiteTimerMax, o.zombieKillPoints, o.zombieShootPoints, o.survivalTimePoints, o.survivalTime, o.herbGetPoints, o.gunShootLossPoints, o.gameSpeed, o.frameCharacter, o.frameColor, o.frameBackgroundColor, o.backgroundColor, o.textColor, o.zombieMoveCounter, o.zombieMoveCounterMax, o.bulletMoveCounter, o.bulletMoveCounterMax)
}

type Settings struct {
	// Title of the game
	title 									string
	// Author of the game
	author 									string
	// Version of the game
	version 								string
	// Width of the game frame
	width		 								int
	// Height of the game frame
	height 									int
	// Starting health of the hero
	start  									Start
	// Options of the game
	options 								Options
}

func (s *Settings) String() string {
	return fmt.Sprintf("Settings{title: %v, author: %v, version: %v, width: %v, height: %v, start: %v, options: %v}", s.title, s.author, s.version, s.width, s.height, s.start, s.options)
}
