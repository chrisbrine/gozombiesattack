package zombies

import (
	"fmt"

	"github.com/gdamore/tcell"
)

type Status struct {
	// Game paused?
	paused 			bool
	// Game over?
	gameOver 		bool
	// Game started?
	started 		bool
	// Exit game?
	exitGame 		bool
	// Screen
	screen   		tcell.Screen
}

func (s *Status) String() string {
	return fmt.Sprintf("Status{paused: %v, gameOver: %v, started: %v, screen: %v}", s.paused, s.gameOver, s.started, s.screen)
}
