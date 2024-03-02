package zombies

import (
	"github.com/gdamore/tcell"
)

// Initialize game
func (g *Game) Init() Game {
	g.InitSettings()
	screen, err := tcell.NewScreen()
	if err != nil {
		panic(err)
	}
	if err := screen.Init(); err != nil {
		panic(err)
	}

	g.status = Status{
		paused:   false,
		gameOver: false,
		started: 	false,
		exitGame: false,
		screen:	 	screen,
	}

	g.InitFrame()
	g.InitHero()
	g.InitZombies()
	g.InitWeapons()
	g.InitBullets()
	g.InitHerbs()

	g.timer = 0
	g.timerFloat = 0

	style := tcell.StyleDefault.Background(tcell.Color(g.settings.options.backgroundColor)).Foreground(g.settings.options.textColor)
	g.status.screen.SetStyle(style)
	g.status.screen.Clear()

	g.debugString = ""

	g.RenderInit()

	return *g
}

// see if point is in the given points
func InDimensions(p Point, tl Point, br Point) bool {
	return p.x >= tl.x &&
		p.x < br.x &&
		p.y >= tl.y &&
		p.y < br.y
}

func (g *Game) RandomBetween(min int, max int) int {
	if min >= max {
		return min
	}
	randomNum := g.Rand(max - min)
	return (min + randomNum)
}

func (g *Game) RandomFloatBetween(min float64, max float64) float64 {
	if min >= max {
		return min
	}
	randomNum := g.RandFloat(max - min)
	return (min + randomNum)
}

func CalculateIncRate(option int, incRate int) int {
	if incRate <= 0 || option <= 0 {
		return 0
	}
	return option / incRate
}

func CalculateIncRateFloat(option float64, incRate float64) float64 {
	if incRate <= 0 || option <= 0 {
		return 0
	}
	return option / incRate
}

func (g *Game) CalculateIncRateByScoreFloat(incRate float64) float64 {
	return CalculateIncRateFloat(float64(g.hero.score), incRate)
}

func (g *Game) CalculateIncRateByScore(incRate int) int {
	return CalculateIncRate(g.hero.score, incRate)
}

func (g *Game) CalculateIncRateByTimer(incRate int) int {
	return CalculateIncRate(g.timer, incRate)
}

// Get Item Position from Drawing
func (d *Drawing) ItemPosition(point Point) Point {
	for i, line := range d.drawing {
		for j, mark := range line {
			if mark.item {
				return Point{
					x: point.x + j + mark.direction.x,
					y: point.y + i + mark.direction.y,
				}
			}
		}
	}
	return Point{}
}
