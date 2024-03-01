package zombies

import "github.com/gdamore/tcell"

// Make the Herb
func DrawHerb(symbol rune, color tcell.Color) Drawing {
	marks := [][]Mark{
		{
			CreateMark(symbol, color),
			CreateMark(symbol, color),
			CreateMark(symbol, color),
		},
		{
			CreateMark(symbol, color),
			CreateMark(symbol, color),
			CreateMark(symbol, color),
		},
	}
	drawing := NewDrawing(marks)
	return drawing
}

// Draw herb that heals
func DrawHerbHealing() Drawing {
	return DrawHerb('♥', ColorRed)
}

// Draw herb that damages zombies
func DrawHerbDamage() Drawing {
	return DrawHerb('X', ColorGreen)
}

// Draw herb that damages hero
func DrawHerbHeroDamage() Drawing {
	return DrawHerb('X', ColorRed)
}

// Draw herb that heals zombies
func DrawHerbZombieHeal() Drawing {
	return DrawHerb('♥', ColorGreen)
}
