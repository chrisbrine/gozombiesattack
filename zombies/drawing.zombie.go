package zombies

func (z *Zombie) Draw() Drawing {
	marks := [][]Mark{
		{
			CreateMark(' ', ColorBlack),
			CreateMark('Z', ColorGreen),
		},
		{
			CreateMark('/', ColorGreen),
			CreateMark('|', ColorGreen),
			CreateMark('\\', ColorGreen),
		},
		{
			CreateMark('/', ColorGreen),
			CreateMark(' ', ColorBlack),
			CreateMark('\\', ColorGreen),
		},
	}
	z.drawing = NewDrawing(marks)
	return z.drawing
}