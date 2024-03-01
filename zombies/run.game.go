package zombies

import (
	"time"
)

// Game cycle

func (g *Game) IsRunning() bool {
	return !g.status.gameOver && !g.status.paused && g.status.started
}

func (g *Game) GameSpeed() int {
	speed := 200 - (g.settings.options.gameSpeed * 2)
	if speed < 0 {
		speed = 0
	} else if speed > 200 {
		speed = 200
	}
	return speed
}

func (g *Game) secondCycle() {
	g.HerbTimer()
	g.settings.options.survivalTimeCounter++
	if g.settings.options.survivalTimeCounter >= g.settings.options.survivalTime {
		g.hero.score += g.settings.options.survivalTimePoints
	}
}

func (g *Game) handleTimer() {
	g.timerFloat += float64(g.GameSpeed()) / 12.25
	newTimer := int(g.timerFloat)
	if newTimer > g.timer {
		g.secondCycle()
	}
	g.timer = newTimer
}

func (g *Game) handleCounters() {
	g.settings.options.zombieMoveCounter++
	g.settings.options.bulletMoveCounter++
}

func (g *Game) Cycle() {
	if g.IsRunning() {
		g.handleCounters()
		g.handleTimer()
		if g.settings.options.zombieMoveCounter >= g.settings.options.zombieMoveCounterMax {
			g.MoveZombies()
		}
		if g.settings.options.bulletMoveCounter >= g.settings.options.bulletMoveCounterMax {
			g.MoveBullets()
		}

		// refresh items
		g.RefreshHerbs()
		g.CheckZombies()
	}
}

func (g*Game) Restart() {
	g.Init()
}

func (g *Game) Run() {
	inputChan := g.GetInput()
	for !g.status.exitGame {
		time.Sleep(time.Duration(g.GameSpeed()) * time.Millisecond)
		g.Input(inputChan)
		g.Cycle()
		g.Render()
	}
	g.status.screen.Fini()
}