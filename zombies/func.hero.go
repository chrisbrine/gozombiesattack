package zombies

// Initialize Hero

func (h *Hero) Init(settings Settings) Hero {
	h.name = settings.start.name
	h.health = settings.start.health
	h.healthMax = settings.start.healthMax
	h.position = Point{1, 1}
	h.herbs = settings.start.herbs
	h.herbsMax = settings.start.herbsMax
	h.damageHerbs = settings.start.damageHerbs
	h.damageHerbsMax = settings.start.damageHerbsMax
	h.score = 0
	h.weapon1 = Weapon{}
	h.weapon2 = Weapon{}
	h.weapon1 = h.weapon.CreateGun()
	h.weapon2 = h.weapon.CreateStrongGun()
	h.currentWeapon = settings.start.weapon
	if h.currentWeapon == 1 {
		h.weapon = h.weapon1
	} else {
		h.weapon = h.weapon2
	}
	h.drawing = h.Draw()
	return *h
}

func (g *Game) InitHero() {
	g.hero = Hero{}
	g.hero = g.hero.Init(g.settings)
}

// Check if Point on hero

func (g *Game) HeroAt(p Point) bool {
	return g.hero.At(p)
}

func (h *Hero) At(p Point) bool {
	topLeftPoint := Point{h.position.x, h.position.y}
	bottomRightPoint := Point{
		x: h.position.x + h.drawing.width,
		y: h.position.y + h.drawing.height,
	}
	return InDimensions(p, topLeftPoint, bottomRightPoint)
}

// Change guns

func (g *Game) HeroWeapon1() {
	g.hero.weapon = g.hero.weapon1
	g.hero.currentWeapon = 1
}

func (g *Game) HeroWeapon2() {
	g.hero.weapon = g.hero.weapon2
	g.hero.currentWeapon = 2
}

// Move Hero

func (g *Game) MoveHero(y int) {
	newY := g.hero.position.y + y
	if g.InFrame(Point{g.hero.position.x, newY}, g.hero.drawing.height, g.hero.drawing.width) {
		g.UnRenderHero()
		g.hero.position.y = newY
		g.RenderHero()
	}
}

// Hero gets herb

func (g *Game) HeroGetHerb(herb Herb) {
	if herb.damage > 0 {
		g.HeroHurt(herb.damage)
	}
	if herb.health > 0 {
		g.hero.herbs += herb.health
		if (g.hero.herbs > g.hero.herbsMax) {
			g.hero.herbs = g.hero.herbsMax
		}
	}
	if herb.zombieDamage > 0 {
		g.hero.damageHerbs += herb.zombieDamage
		if (g.hero.damageHerbs > g.hero.damageHerbsMax) {
			g.hero.damageHerbs = g.hero.damageHerbsMax
		}
	}
	g.hero.score += g.settings.options.herbGetPoints
}

// Hero Dies, Hurt, or Healed

func (g *Game) HeroDies() {
	g.hero.health = 0
	g.status.gameOver = true
	g.UnRenderHero()
}

func (g *Game) HeroHurt(damage int) {
	g.hero.health -= damage
	if g.hero.health <= 0 {
		g.HeroDies()
	}
}

// hero use herb

func (g *Game) HeroUseHerb() {
	if g.hero.herbs > 0 && g.hero.health < g.hero.healthMax {
		g.hero.herbs--
		g.hero.health++
	}
}

func (g *Game) HeroUseDamageHerb() {
	if g.hero.damageHerbs > 0 {
		g.hero.damageHerbs--
		g.DamageAllZombies(1)
	}
}
