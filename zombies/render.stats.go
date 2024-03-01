package zombies

import "fmt"

func (g *Game) DrawHealth(position Point) {
	if g.hero.health <= 0 {
		g.TextColor(position, "XXXXXX", ColorRed, g.settings.options.backgroundColor)
	} else {
		health := "♥"
		for i := 0; i < g.hero.healthMax; i++ {
			if i < g.hero.health {
				g.TextColor(Point{position.x + i, position.y}, health, ColorRed, g.settings.options.backgroundColor)
			} else {
				g.Text(Point{position.x + i, position.y}, " ")
			}
		}
	}
}

func (g *Game) RenderWeapons(startX int) {
	// render weapon choices
	unselectedWeaponColor := g.settings.options.textColor
	unselectedWeaponBG := g.settings.options.backgroundColor
	selectedWeaponColor := ColorYellow
	selectedWeaponBG := ColorBlue
	posY := g.frame.position.y + g.frame.height - g.hero.weapon1.drawing.height - 1
	if g.hero.currentWeapon == 1 {
		g.DrawDrawingColor(Point{startX, posY}, g.hero.weapon1.drawing, selectedWeaponColor, selectedWeaponBG)
		g.DrawDrawingColor(Point{startX + 4, posY}, g.hero.weapon2.drawing, unselectedWeaponColor, unselectedWeaponBG)
	} else {
		g.DrawDrawingColor(Point{startX, posY}, g.hero.weapon1.drawing, unselectedWeaponColor, unselectedWeaponBG)
		g.DrawDrawingColor(Point{startX + 4, posY}, g.hero.weapon2.drawing, selectedWeaponColor, selectedWeaponBG)
	}
	herbDrawing := DrawHerbHealing()
	g.DrawDrawingColor(Point{startX + 8, posY}, herbDrawing, ColorRed, g.settings.options.backgroundColor)
	herbDrawing = DrawHerbDamage()
	g.DrawDrawingColor(Point{startX + 12, posY}, herbDrawing, ColorRed, g.settings.options.backgroundColor)
}

func (g *Game) GetTimerString() string {
	// returns it as 00:00:00 (hours, minutes, seconds)
	hours := g.timer / 3600
	minutes := (g.timer % 3600) / 60
	seconds := g.timer % 60
	return fmt.Sprintf("%02d:%02d:%02d", hours, minutes, seconds)
}

func (g *Game) RenderStats() {
	startX := g.frame.position.x + g.frame.width + 2
	width, _ := g.status.screen.Size()
	g.Text(Point{startX, 0}, "Score")
	g.Text(Point{startX, 1}, fmt.Sprint(g.hero.score))

	g.Text(Point{width - 9, 0}, "Time")

	g.Text(Point{startX, 3}, "Health")

	posY := g.frame.position.y + g.frame.height - g.hero.weapon1.drawing.height - 2
	g.Text(Point{startX, posY}, "*1*")
	g.Text(Point{startX + 4, posY}, "*2*")
	g.Text(Point{startX + 8, posY}, "*3*")
	g.Text(Point{startX + 12, posY}, "*4*")

	g.RefreshStats()
}

func (g *Game) RefreshGameStatus() {
	startX := g.frame.position.x + g.frame.width + 2
	startY := 4 + ((g.frame.position.y + g.frame.height - g.hero.weapon1.drawing.height - 2) / 2)
	// erase all of the below
	width, _ := g.status.screen.Size()
	for y := startY; y < (startY + 4); y++ {
		for x := startX; x < width; x++ {
			g.SetPoint(Point{x, y}, ' ', g.settings.options.textColor, g.settings.options.backgroundColor)
		}
	}
	if g.status.gameOver {
		g.Text(Point{startX, startY}, "Game Over")
		g.Text(Point{startX, startY + 1}, "'q' to exit")
		g.Text(Point{startX, startY + 2}, "'r' to restart")
	} else if g.status.paused {
		g.Text(Point{startX, startY}, "Paused")
		g.Text(Point{startX, startY + 1}, "'p' to continue")
		g.Text(Point{startX, startY + 2}, "'q' to exit")
		g.Text(Point{startX, startY + 3}, "'r' to restart")
	} else if !g.status.started {
		g.Text(Point{startX, startY}, "Start your game!")
		g.Text(Point{startX, startY + 1}, "'Up' to start")
		g.Text(Point{startX, startY + 2}, "'q' to exit")
	}
}

func (g *Game) RefreshStats() {
	startX := g.frame.position.x + g.frame.width + 2
	width, _ := g.status.screen.Size()
	timer := g.GetTimerString()
	g.Text(Point{width - 9, 1}, timer)

	g.DrawHealth(Point{startX, 4})
	g.Text(Point{startX, 1}, "       ")
	g.Text(Point{startX, 1}, fmt.Sprint(g.hero.score))

	posY := g.frame.position.y + g.frame.height - 1
	g.Text(Point{startX + 9, posY}, "  ")
	g.Text(Point{startX + 9, posY}, fmt.Sprint(g.hero.herbs))
	g.Text(Point{startX + 13, posY}, "  ")
	g.Text(Point{startX + 13, posY}, fmt.Sprint(g.hero.damageHerbs))

	g.RenderWeapons(startX)

	g.RefreshGameStatus()
}