package zombies

// Render the hero drawing, include the current gun

func (g *Game) RenderHero() {
	// Draw the hero
	heroPoint := Point{g.hero.position.x, g.hero.position.y}
	g.DrawDrawing(heroPoint, g.hero.drawing)

	// Draw the gun
	weaponPoint := g.hero.drawing.ItemPosition(g.hero.position)
	g.DrawDrawing(weaponPoint, g.hero.weapon.drawing)
}

func (g *Game) UnRenderHero() {
	// UnDraw the hero
	heroPoint := Point{g.hero.position.x, g.hero.position.y}
	g.EraseDrawing(heroPoint, g.hero.drawing)

	// UnDraw the gun
	weaponPoint := g.hero.drawing.ItemPosition(g.hero.position)
	g.EraseDrawing(weaponPoint, g.hero.weapon.drawing)
}
