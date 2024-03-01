package zombies

func (g *Game) RenderBullet(index int) {
	b := g.bullets[index]
	g.DrawDrawing(b.position, b.drawing)
}

func (g *Game) RenderBullets() {
	for _, b := range g.bullets {
		g.DrawDrawing(b.position, b.drawing)
	}
}

func (g *Game) UnRenderBullet(index int) {
	b := g.bullets[index]
	g.EraseDrawing(b.position, b.drawing)
}

func (g *Game) UnRenderBullets() {
	for _, b := range g.bullets {
		g.EraseDrawing(b.position, b.drawing)
	}
}
