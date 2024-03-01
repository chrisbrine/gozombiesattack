package zombies

func (g *Game) RenderZombie(index int) {
	if index < 0 || index >= len(g.zombies) {
		return
	}
	z := g.zombies[index]
	g.DrawDrawing(z.position, z.drawing)
}

func (g *Game) RenderZombies() {
	for _, z := range g.zombies {
		g.DrawDrawing(z.position, z.drawing)
	}
}

func (g *Game) UnRenderZombie(index int) {
	if index < 0 || index >= len(g.zombies) {
		return
	}
	z := g.zombies[index]
	g.EraseDrawing(z.position, z.drawing)
}

func (g *Game) UnRenderZombies() {
	for _, z := range g.zombies {
		g.EraseDrawing(z.position, z.drawing)
	}
}
