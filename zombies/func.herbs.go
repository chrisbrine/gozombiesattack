package zombies

// Initialize initial herbs
func (g *Game) InitHerbs() {
	g.herbs = []Herb{}
	for i := 0; i < g.settings.options.herbs; i++ {
		g.CreateRandomHerb()
	}
}

// Create a new herb
func (g *Game) CreateHerb(
	x, y int,
	drawing Drawing,
	heal int,
	damage int,
	zombieDamage int,
	zombieHeal int,
) Herb {
	timer := g.RandomBetween(
		g.settings.options.herbTimerMin,
		g.settings.options.herbTimerMax,
	)
	herb := Herb{
		position: Point{x, y},
		drawing: drawing,
		health: heal,
		damage: damage,
		timer: timer,
		zombieDamage: zombieDamage,
		zombieHeal: zombieHeal,
	}
	return herb
}

// Create a healing herb
func (g *Game) CreateHealingHerb(x, y int) Herb {
	result := g.RandomBetween(
		g.settings.options.healingHerbMin,
		g.settings.options.healingHerbMax,
	)
	drawing := DrawHerbHealing();
	return g.CreateHerb(x, y, drawing, result, 0, 0, 0)
}

// Create a damage herb
func (g *Game) CreateDamageHerb(x, y int) Herb {
	result := g.RandomBetween(
		g.settings.options.damageHerbMin,
		g.settings.options.damageHerbMax,
	)
	drawing := DrawHerbHeroDamage();
	return g.CreateHerb(x, y, drawing, 0, result, 0, 0)
}

// Create a zombie damage herb
func (g *Game) CreateZombieDamageHerb(x, y int) Herb {
	result := g.RandomBetween(
		g.settings.start.zombieMinDamage,
		g.settings.start.zombieMaxDamage,
	)
	drawing := DrawHerbDamage();
	return g.CreateHerb(x, y, drawing, 0, 0, result, 0)
}

// Create a zombie heal herb
func (g *Game) CreateZombieHealHerb(x, y int) Herb {
	result := g.RandomBetween(
		g.settings.start.zombieMinHealth,
		g.settings.start.zombieMaxHealth,
	)
	drawing := DrawHerbZombieHeal();
	return g.CreateHerb(x, y, drawing, 0, 0, 0, result)
}

// Random Change for Herb Type, if all fail it will just create a healing herb
func (g *Game) RandomHerbType(x, y int) Herb {
	percent := g.RandomBetween(0, 100)
	if percent < g.settings.options.healingHerbPercent {
		return g.CreateHealingHerb(x, y)
	}
	percent -= g.settings.options.healingHerbPercent
	if percent < g.settings.options.damageHerbPercent {
		return g.CreateDamageHerb(x, y)
	}
	percent -= g.settings.options.damageHerbPercent
	if percent < g.settings.options.zombieDamageHerbPercent {
		return g.CreateZombieDamageHerb(x, y)
	}
	percent -= g.settings.options.zombieDamageHerbPercent
	if percent < g.settings.options.zombieHealHerbPercent {
		return g.CreateZombieHealHerb(x, y)
	}
	return g.CreateHealingHerb(x, y)
}

// Get Random herb point
func (g *Game) RandomHerbPoint() Point {
	minXPoint := g.hero.position.x + g.hero.drawing.width + 3
	maxXPoint := g.frame.position.x + g.frame.width - 3
	minYPoint := g.frame.position.y + 3
	maxYPoint := g.frame.position.y + g.frame.height - 3
	return Point{
		g.RandomBetween(minXPoint, maxXPoint),
		g.RandomBetween(minYPoint, maxYPoint),
	}
}

// Create Random Herb at random point
func (g *Game) CreateRandomHerb() {
	point := g.RandomHerbPoint()
	herb := g.RandomHerbType(point.x, point.y)
	g.herbs = append(g.herbs, herb)
	g.RenderHerb(len(g.herbs) - 1)
}

// Check how many herbs there should be
func (g *Game) NeededHerbCount() int {
	return g.settings.options.herbs + g.CalculateIncRateByScore(g.settings.options.herbsIncRate)
}

// Refresh Herbs as needed
func (g *Game) RefreshHerbs() {
	herbCount := g.NeededHerbCount()
	for len(g.herbs) < herbCount {
		g.CreateRandomHerb()
	}
}

// Check if an herb is at Point
func (g *Game) HerbAt(p Point) bool {
	for i := len(g.herbs) - 1; i >= 0; i-- {
		herb := g.herbs[i]
		if herb.At(p) {
			return true
		}
	}
	return false
}

func (h *Herb) At(p Point) bool {
	topLeftPoint := Point{h.position.x, h.position.y}
	bottomRightPoint := Point{
		x: h.position.x + h.drawing.width,
		y: h.position.y + h.drawing.height,
	}
	return InDimensions(p, topLeftPoint, bottomRightPoint)
}

func (g *Game) HerbAtPoint(p Point) (int, Herb) {
	for i := len(g.herbs) - 1; i >= 0; i-- {
		herb := g.herbs[i]
		if herb.At(p) {
			return i, herb
		}
	}
	return -1, Herb{}
}

func (g *Game) KillHerb(i int) {
	if i < 0 || i >= len(g.herbs) {
		return
	}
	g.UnRenderHerb(i)
	g.herbs = append(g.herbs[:i], g.herbs[i+1:]...)
}

// Handle HerbTimer

func (g *Game) HerbTimer() {
	for i := len(g.herbs) - 1; i >= 0; i-- {
		herb := g.herbs[i]
		g.herbs[i].timer--
		if herb.timer <= 0 {
			g.KillHerb(i)
		}
	}
}
