package zombies

import (
	"github.com/gdamore/tcell"
)

// Draw a char on a point on the screen
func (g *Game) SetPoint(p Point, c rune, fg tcell.Color, bg tcell.Color) {
	g.status.screen.SetContent(p.x, p.y, c, nil, tcell.StyleDefault.Foreground(fg).Background(bg))
}

// Draw an unfilled rectangle on the screen
func (g *Game) DrawRect(p1 Point, p2 Point, c rune, fg tcell.Color, bg tcell.Color) {
	for x := p1.x; x <= p2.x; x++ {
		g.SetPoint(Point{x, p1.y}, c, fg, bg)
	}
	for y := p1.y + 1; y < p2.y; y++ {
		g.SetPoint(Point{p1.x, y}, c, fg, bg)
		for x:= p1.x + 1; x < p2.x; x++ {
			g.SetPoint(Point{x, y}, ' ', fg, bg)
		}
		g.SetPoint(Point{p2.x, y}, c, fg, bg)
	}
	for x := p1.x; x <= p2.x; x++ {
		g.SetPoint(Point{x, p2.y}, c, fg, bg)
	}
}

// Draw a drawing at point on the screen
func (g *Game) DrawDrawing(p Point, d Drawing) {
	for i, markRow := range d.drawing {
		for j, mark := range markRow {
			if mark.char != ' ' {
				g.SetPoint(Point{p.x + j, p.y + i}, mark.char, mark.color, g.frame.bgColor)
			}
		}
	}
}

// Draw a drawing in a new color
func (g *Game) DrawDrawingColor(p Point, d Drawing, fg tcell.Color, bg tcell.Color) {
	for i, markRow := range d.drawing {
		for j, mark := range markRow {
			if mark.char != ' ' {
				g.SetPoint(Point{p.x + j, p.y + i}, mark.char, fg, bg)
			}
		}
	}
}

// Erase a drawing at point on the screen
func (g *Game) EraseDrawing(p Point, d Drawing) {
	for i, markRow := range d.drawing {
		for j, mark := range markRow {
			if mark.char != ' ' {
				g.SetPoint(Point{p.x + j, p.y + i}, ' ', g.frame.color, g.frame.bgColor)
			}
		}
	}
}

// Write text on the screen
func (g *Game) TextColor(p Point, text string, fg tcell.Color, bg tcell.Color) {
	for i, c := range text {
		g.SetPoint(Point{p.x + i, p.y}, c, fg, bg)
	}
}

// Write text on the screen using default colors
func (g *Game) Text(p Point, text string) {
	g.TextColor(p, text, g.settings.options.textColor, g.settings.options.backgroundColor)
}

func (g *Game) RenderInit() {
	g.RenderFrame()
	g.RenderHerbs()
	g.RenderZombies()
	g.RenderHero()
	g.RenderStats()
	g.status.screen.Show()
}

func (g *Game) Render() {
	// Draw the frame on the screen
	// g.Render()
	g.RenderHerbs()
	// g.RenderHero()
	g.RenderZombies()
	// g.RenderBullets()
	g.RefreshStats()
	if g.debugString != "" {
		_, height := g.status.screen.Size()
		g.Text(Point{0, height - 1}, g.debugString)
	}
	g.status.screen.Show()
	g.UnRenderHerbs()
	g.UnRenderZombies()
}