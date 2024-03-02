package zombies

func (h *Hero) Draw() Drawing {
	weaponMark := CreateMark('∏', ColorGrey)
	weaponMark.AddItem(true, Point{x: 1, y: 0})
	marks := [][]Mark{
		{
			CreateMark(' ', ColorBlack),
			CreateMark('ੳ', ColorWhite),
		},
		{
			CreateMark(' ', ColorBlack),
			weaponMark,
		},
		{
			CreateMark(' ', ColorBlack),
			CreateMark('∐', ColorGrey),
		},
		{
			CreateMark('/', ColorWhite),
			CreateMark(' ', ColorBlack),
			CreateMark('\\', ColorWhite),
		},
	}
	h.drawing = NewDrawing(marks)
	return h.drawing
}
