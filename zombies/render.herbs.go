package zombies

func (g *Game) RenderHerb(index int) {
	if index < 0 || index >= len(g.herbs) {
		return
	}
	h := g.herbs[index]
	g.DrawDrawing(h.position, h.drawing)
}

func (g *Game) RenderHerbs() {
	for _, h := range g.herbs {
		g.DrawDrawing(h.position, h.drawing)
	}
}

func (g *Game) UnRenderHerb(index int) {
	if index < 0 || index >= len(g.herbs) {
		return
	}
	h := g.herbs[index]
	g.EraseDrawing(h.position, h.drawing)
}

func (g *Game) UnRenderHerbs() {
	for _, h := range g.herbs {
		g.EraseDrawing(h.position, h.drawing)
	}
}
