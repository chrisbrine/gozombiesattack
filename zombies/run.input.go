package zombies

import "github.com/gdamore/tcell"

func (g *Game) GetInput() chan string {
	inputChan := make(chan string)
	go func() {
		for {
			switch ev := g.status.screen.PollEvent().(type) {
				case *tcell.EventKey:
					switch ev.Key() {
						case tcell.KeyEscape:
							inputChan <- "exit"
						case tcell.KeyUp:
							inputChan <- "up"
						case tcell.KeyDown:
							inputChan <- "down"
						case tcell.KeyRight:
							inputChan <- "right"
						case tcell.KeyLeft:
							inputChan <- "left"
						case tcell.KeyCtrlC:
							inputChan <- "exit"
						case tcell.KeyRune:
							if ev.Rune() == 'p' || ev.Rune() == 'P' {
								inputChan <- "pause"
							}
							if ev.Rune() == 'r' || ev.Rune() == 'R' {
								inputChan <- "restart"
							}
							if ev.Rune() == 'q' || ev.Rune() == 'Q' {
								inputChan <- "exit"
							}
							if ev.Rune() == 'w' || ev.Rune() == 'W' {
								inputChan <- "up"
							}
							if ev.Rune() == 's' || ev.Rune() == 'S' {
								inputChan <- "down"
							}
							if ev.Rune() == ' ' {
								inputChan <- "shoot"
							}
							if ev.Rune() == '1' {
								inputChan <- "weapon1"
							}
							if ev.Rune() == '2' {
								inputChan <- "weapon2"
							}
							if ev.Rune() == '3' {
								inputChan <- "herb"
							}
							if ev.Rune() == '4' {
								inputChan <- "damageHerb"
							}
					}
			}
		}
	}()

	return inputChan
}

func (g *Game) Input(inputChan chan string) {
	var input string
	select {
		case input = <-inputChan:
		default:
			input = ""
	}
	if input != "" {
		g.status.paused = false
		g.status.started = true
	}
	switch input {
		case "exit":
			g.status.exitGame = true
		case "up":
			if g.IsRunning() {
				g.MoveHero(Point{0, -1})
			}
		case "down":
			if g.IsRunning() {
				g.MoveHero(Point{0, 1})
			}
		case "right":
			if g.IsRunning() {
				g.MoveHero(Point{1, 0})
			}
		case "left":
			if g.IsRunning() {
				g.MoveHero(Point{-1, 0})
			}
		case "pause":
			g.status.paused = true
		case "shoot":
			if g.IsRunning() {
				g.ShootGun()
			}
		case "weapon1":
			if g.IsRunning() {
				g.HeroWeapon1()
			}
		case "weapon2":
			if g.IsRunning() {
				g.HeroWeapon2()
			}
		case "herb":
			if g.IsRunning() {
				g.HeroUseHerb()
			}
		case "damageHerb":
			if g.IsRunning() {
				g.HeroUseDamageHerb()
			}
		case "restart":
			g.Restart()
	}
}
