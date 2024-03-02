package zombies

// Initialize starting zombies
func (g *Game) InitZombies() {
	g.zombies = []Zombie{}
	for i := 0; i < g.settings.start.zombies; i++ {
		g.CreateZombie()
	}
}

// Check for more zombies to create
func (g *Game) CheckZombies() {
	needZombieCount := g.settings.start.zombies + g.CalculateIncRateByScore(g.settings.options.zombiesIncRate)
	if needZombieCount > g.settings.options.zombiesMax {
		needZombieCount = g.settings.options.zombiesMax
	}
	if len(g.zombies) < g.settings.start.zombies {
		g.CreateZombie()
	}
	for len(g.zombies) < needZombieCount {
		g.CreateZombie()
	}
}

// Check if any zombies in Point given
func (g *Game) ZombieAt(p Point) bool {
	for _, z := range g.zombies {
		if z.At(p) {
			return true
		}
	}
	return false
}

func (z *Zombie) At(p Point) bool {
	topLeftPoint := Point{z.position.x, z.position.y}
	BottomRightPoint := Point{
		z.position.x + z.drawing.width,
		z.position.y + z.drawing.height,
	}
	return InDimensions(p, topLeftPoint, BottomRightPoint)
}

// initialize new zombie
func (g *Game) NewZombie() Zombie {
	zombie := Zombie{}
	min1 := g.settings.start.zombieMinHealth
	calcMax1 := g.CalculateIncRateByScore(g.settings.options.zombieHealthIncRate)
	if calcMax1 > g.settings.options.zombieHealthMaxRate {
		calcMax1 = g.settings.options.zombieHealthMaxRate
	}
	max1 := g.settings.start.zombieMaxHealth + calcMax1
	min2 := g.settings.start.zombieMinSpeed
	calcMax2 := g.CalculateIncRateByScoreFloat(g.settings.options.zombieSpeedIncRate)
	if calcMax2 > g.settings.options.zombieSpeedMaxRate {
		calcMax2 = g.settings.options.zombieSpeedMaxRate
	}
	max2 := g.settings.start.zombieMaxSpeed + calcMax2
	calcMax3 := g.CalculateIncRateByScore(g.settings.options.zombieDamageIncRate)
	if calcMax3 > g.settings.options.zombieDamageMaxRate {
		calcMax3 = g.settings.options.zombieDamageMaxRate
	}
	min3 := g.settings.start.zombieMinDamage
	max3 := g.settings.start.zombieMaxDamage + calcMax3
	zombie.health = g.RandomBetween(min1, max1)
	zombie.speed = g.RandomFloatBetween(min2, max2)
	zombie.damage = g.RandomBetween(min3, max3)
	zombie.drawing = zombie.Draw()
	zombie.maxBiteTimer = g.RandomBetween(g.settings.options.zombieBiteTimerMin, g.settings.options.zombieBiteTimerMax)
	zombie.biteTimer = zombie.maxBiteTimer
	x := (g.frame.position.x + g.frame.width) - 1 - zombie.drawing.width
	y := g.RandomBetween(1, g.frame.height-zombie.drawing.height - 1)
	zombie.position = Point{x, y}
	return zombie
}

// Create a random zombie in a random position at the far right of the frame
func (g *Game) CreateZombie() {
	zombie := g.NewZombie()
	g.zombies = append(g.zombies, zombie)
	g.RenderZombie(len(g.zombies) - 1)
}

// Move zombies
func (g *Game) MoveZombies() {
	for i, z := range g.zombies {
		g.MoveZombie(i, z)
	}
}

func (g *Game) CheckZombieHeightPoints(i int, z Zombie, newX int) bool {
	if i < 0 || i >= len(g.zombies) {
		return false
	}
	alreadyHurtHero := false
	result := false
	for newY := z.position.y; newY < z.position.y+z.drawing.height; newY++ {
		if g.BulletAt(Point{newX, newY}) || g.BulletAt(Point(Point{z.position.x, newY})) {
			idx, bullet := g.GetBulletAtPoint(Point{newX, newY})
			damage := bullet.damage
			if damage < 1 {
				damage = 1
			}
			g.ZombieHurt(i, damage)
			g.KillBullet(idx)
			result = true
		} else if g.HeroAt(Point{newX, newY}) && !alreadyHurtHero {
			if z.biteTimer >= z.maxBiteTimer {
				alreadyHurtHero = true
				damage := z.damage
				if damage < 1 {
					damage = 1
				}
				g.HeroHurt(damage)
				g.zombies[i].biteTimer = 0
			} else {
				g.zombies[i].biteTimer++
			}
			result = true
		} else if g.HerbAt(Point{newX, newY}) {
			g.ZombieEatsHerb(i, Point{newX, newY})
			result = true
		}
	}
	return result
}

func (g *Game) MoveZombie(i int, z Zombie) Zombie {
	if i < 0 || i >= len(g.zombies) {
		return z
	}
	newX := z.position.x - int(z.speed)
	if !g.CheckZombieHeightPoints(i, z, newX) {
		if InDimensions(Point{newX, z.position.y}, Point{g.frame.position.x, g.frame.position.y}, Point{g.frame.position.x + 2, g.frame.height}) {
			newX = g.frame.position.x + 1
			g.HeroDies()
			g.UnRenderZombie(i)
			g.zombies[i].position.x = newX
			g.RenderZombie(i)
		} else {
			if newX <= g.frame.position.x {
				newX = g.frame.position.x + 1
			}
			g.UnRenderZombie(i)
			g.zombies[i].position.x = newX
			g.RenderZombie(i)
		}
	}
	return z
}

// Zombie Eats Herb

func (g *Game) ZombieEatsHerb(i int, p Point) {
	if i < 0 || i >= len(g.zombies) {
		return
	}
	for i := len(g.herbs) - 1; i >= 0; i-- {
		h := g.herbs[i]
		if h.At(p) {
			if h.zombieDamage > 0 {
				g.ZombieHurt(i, h.zombieDamage)
			}
			if h.zombieHeal > 0 {
				g.ZombieHeal(i, h.zombieHeal)
				if g.zombies[i].health > g.settings.options.zombieHealthMaxRate {
					g.zombies[i].health = g.settings.options.zombieHealthMaxRate
				}
			}
			g.KillHerb(i)
		}
	}
}

// Zombie Hurt or Killed or Healed

func (g *Game) ZombieDies(i int) {
	if i < 0 || i >= len(g.zombies) {
		return
	}
	g.UnRenderZombie(i)
	g.zombies = append(g.zombies[:i], g.zombies[i+1:]...)
	g.hero.score += g.settings.options.zombieKillPoints
}

func (g *Game) ZombieHurt(i int, damage int) {
	if i < len(g.zombies) && i >= 0 {
		g.zombies[i].health -= damage
		if g.zombies[i].health <= 0 {
			g.ZombieDies(i)
		}
	}
}

func (g *Game) ZombieHurtAt(p Point, damage int) {
	for i := len(g.zombies) - 1; i >= 0; i-- {
		z := g.zombies[i]
		if z.At(p) {
			g.ZombieHurt(i, damage)
		}
	}
}

func (g *Game) ZombieHeal(i int, heal int) {
	if i < 0 || i >= len(g.zombies) {
		return
	}
	g.zombies[i].health += heal
}

func (g *Game) DamageAllZombies(damage int) {
	for i := len(g.zombies) - 1; i >= 0; i-- {
		g.ZombieHurt(i, damage)
	}
}
