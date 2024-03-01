package zombies

func NewGame() *Game {
	game := &Game{}
	game.Init()
	return game
}