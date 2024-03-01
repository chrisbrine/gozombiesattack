package zombies

// Drawing is the drawing of a gun
func (w *Weapon) DrawGun() Drawing {
	marks := [][]Mark{
		{
			CreateMark('_', ColorRed),
			CreateMark('_', ColorRed),
		},
		{
			CreateMark('/', ColorRed),
		},
	}
	w.drawing = NewDrawing(marks)
	return w.drawing
}

// Drawing bullet for the gun
func (w *Weapon) DrawGunBullet() Drawing {
	marks := [][]Mark{
		{
			CreateMark('.', ColorWhite),
		},
	}
	w.bulletDrawing = NewDrawing(marks)
	return w.bulletDrawing
}

// Drawing is of a strong gun
func (w *Weapon) DrawStrongGun() Drawing {
	marks := [][]Mark{
		{
			CreateMark('=', ColorYellow),
			CreateMark('=', ColorYellow),
		},
		{
			CreateMark('/', ColorYellow),
		},
	}
	w.drawing = NewDrawing(marks)
	return w.drawing
}

// Drawing bullet for the strong gun
func (w *Weapon) DrawStrongGunBullet() Drawing {
	marks := [][]Mark{
		{
			CreateMark('*', ColorYellow),
		},
	}
	w.bulletDrawing = NewDrawing(marks)
	return w.bulletDrawing
}
