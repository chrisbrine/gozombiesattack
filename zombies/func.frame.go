package zombies

// Initialize Frame
func (f *Frame) Init(settings Settings) Frame {
	f.width = settings.width
	f.height = settings.height
	f.position = Point{0, 0}
	f.char = settings.options.frameCharacter
	f.color = settings.options.frameColor
	f.bgColor = settings.options.frameBackgroundColor

	return *f
}

func (g *Game) InitFrame() {
	g.frame = Frame{}
	g.frame = g.frame.Init(g.settings)
}

// check if point inside frame

func (g *Game) InFrame(p Point, height int, width int) bool {
	return InDimensions(
		p,
		Point{g.frame.position.x + 1, g.frame.position.y + 1},
		Point{g.frame.position.x + g.frame.width - width + 1, g.frame.position.y + g.frame.height - height + 1},
	)
}
