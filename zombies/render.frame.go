package zombies

// Draw the frame on the screen
func (g *Game) RenderFrame() {
	p1 := Point{g.frame.position.x, g.frame.position.y}
	p2 := Point{g.frame.position.x + g.frame.width, g.frame.position.y + g.frame.height}
	g.DrawRect(p1, p2, g.frame.char, g.frame.color, g.frame.bgColor)
}