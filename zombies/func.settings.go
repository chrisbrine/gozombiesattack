package zombies

import (
	"math/rand/v2"
)

func (g *Game) Rand(num int) int {
	return rand.IntN(num)
}

func (g *Game) RandFloat(num float64) float64 {
	result := rand.Float64() * num
	return result
}

// Initialize Settings
func (g *Game) InitSettings() {
	g.settings = Settings{
		width: 60,
		height: 23,
		title: "Zombies Attack!",
		author: "Chris Brine",
		version: "1.0.0",
		start: Start{
			name: "Hero",
			health: 3,
			healthMax: 6,
			weapon: 1,
			herbs: 1,
			herbsMax: 10,
			damageHerbs: 0,
			damageHerbsMax: 5,
			zombies: 1,
			zombieMaxHealth: 1,
			zombieMinHealth: 1,
			zombieMaxDamage: 1, // this will always be 1
			zombieMinDamage: 1, // this will always be 1
			zombieMaxSpeed: 1,
			zombieMinSpeed: 1,
		},
		options: Options{
			herbs: 1,
			herbsIncRate: 2000,
			herbsMax: 10,
			healingHerbMax: 2,
			healingHerbMin: 1,
			damageHerbMax: 2,
			damageHerbMin: 1,
			herbTimerMin: 40,
			herbTimerMax: 100,
			healingHerbPercent: 70,
			damageHerbPercent: 0,
			zombieDamageHerbPercent: 30,
			zombieHealHerbPercent: 0,
			zombiesIncRate: 1000,
			zombiesMax: 20,
			zombieHealthIncRate: 3000, // Set to 0 to prevent health increase
			zombieHealthMaxRate: 2,
			zombieSpeedIncRate: 7500,
			zombieSpeedMaxRate: 3,
			zombieDamageIncRate: 4000,
			zombieDamageMaxRate: 1,
			zombieKillPoints: 100,
			zombieShootPoints: 10,
			zombieBiteTimerMin: 10,
			zombieBiteTimerMax: 25,
			survivalTimePoints: 1,
			survivalTime: 1,
			survivalTimeCounter: 0, // always set to 0
			herbGetPoints: 20,
			gunShootLossPoints: 1,
			gameSpeed: 50,
			frameCharacter: '@',
			frameColor: ColorWhite,
			frameBackgroundColor: ColorBlack,
			backgroundColor: ColorLightGrey,
			textColor: ColorBlack,
			zombieMoveCounter: 0, // do not change this
			zombieMoveCounterMax: 2, // allows slowing down zombie movements globally, higher is slower
			bulletMoveCounter: 0, // do not change this
			bulletMoveCounterMax: 1, // allows slowing down bullet movements globally, higher is slower
		},
	}
}
