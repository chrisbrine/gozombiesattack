package zombies

import (
	"math/rand"
)

// Initialize

func (g *Game) InitWeapons() {
	g.weapons = []Weapon{}
}

func (g *Game) InitBullets() {
	g.bullets = []Bullet{}
}

// Create a gun

func (w *Weapon) CreateGun() Weapon {
	w.name = "Gun"
	w.damage = 1
	w.weaponRange = 100
	w.drawing = w.DrawGun()
	w.speed = 3
	w.bulletDrawing = w.DrawGunBullet()
	w.bulletPosition = Point{2, 0}
	return *w
}

// Create a strong gun
func (w *Weapon) CreateStrongGun() Weapon {
	w.name = "Strong Gun"
	w.damage = 3
	w.weaponRange = 100
	w.drawing = w.DrawStrongGun()
	w.speed = 1
	w.bulletDrawing = w.DrawStrongGunBullet()
	w.bulletPosition = Point{2, 0}
	return *w
}

// Create a random weapon
func (w *Weapon) CreateRandomWeapon(rand *rand.Rand) Weapon {
	weapon := rand.Intn(2)
	if weapon == 0 {
		return w.CreateGun()
	}
	return w.CreateStrongGun()
}

// Shoot Gun

func (g *Game) ShootGun() {
	// this will create a new bullet from the gun
	itemPos := g.hero.drawing.ItemPosition(g.hero.position)
	bullet := Bullet{
		position: Point{
			x: itemPos.x + g.hero.weapon.bulletPosition.x,
			y: itemPos.y + g.hero.weapon.bulletPosition.y,
		},
		drawing: g.hero.weapon.bulletDrawing,
		damage: g.hero.weapon.damage,
		speed: g.hero.weapon.speed,
	}

	g.bullets = append(g.bullets, bullet)
	g.RenderBullet(len(g.bullets) - 1)
}

// Move Bullets

func (g *Game) MoveBullets() {
	for i := 0; i < len(g.bullets); i++ {
		newX := g.bullets[i].position.x + int(g.bullets[i].speed)
		if g.InFrame(Point{newX, g.bullets[i].position.y}, g.bullets[i].drawing.width, g.bullets[i].drawing.height) {
			if g.ZombieAt(Point{newX, g.bullets[i].position.y}) {
				g.hero.score += g.settings.options.zombieShootPoints
				g.ZombieHurtAt(Point{newX, g.bullets[i].position.y}, g.bullets[i].damage)
				g.KillBullet(i)
				i--
			} else if g.ZombieAt(Point{g.bullets[i].position.x, g.bullets[i].position.y}) {
				g.hero.score += g.settings.options.zombieShootPoints
				g.ZombieHurtAt(Point{g.bullets[i].position.x, g.bullets[i].position.y}, g.bullets[i].damage)
				g.KillBullet(i)
				i--
			} else if g.HerbAt(Point{newX, g.bullets[i].position.y}) {
				idx, herb := g.HerbAtPoint(Point{newX, g.bullets[i].position.y})
				g.HeroGetHerb(herb)
				g.KillHerb(idx)
				g.KillBullet(i)
				i--
			} else if g.HerbAt(Point{g.bullets[i].position.x, g.bullets[i].position.y}) {
				idx, herb := g.HerbAtPoint(Point{g.bullets[i].position.x, g.bullets[i].position.y})
				g.HeroGetHerb(herb)
				g.KillHerb(idx)
				g.KillBullet(i)
				i--
			} else {
				g.UnRenderBullet(i)
				g.bullets[i].position.x = newX
				g.RenderBullet(i)
			}
		} else {
			g.KillBullet(i)
			i--
		}
	}
}

func (g *Game) KillBullet(i int) {
	if i < 0 || i >= len(g.bullets) {
		return
	}
	g.UnRenderBullet(i)
	g.bullets = append(g.bullets[:i], g.bullets[i+1:]...)
}

func (g *Game) KillBulletAtPoint(p Point) {
	for i := len(g.bullets) - 1; i >= 0; i-- {
		bullet := g.bullets[i]
		if bullet.At(p) {
			g.KillBullet(i)
			break
		}
	}
}

// Check if Bullet at Point

func (g *Game) BulletAt(p Point) bool {
	for i := len(g.bullets) - 1; i >= 0; i-- {
		bullet := g.bullets[i]
		if bullet.At(p) {
			return true
		}
	}
	return false
}

func (b *Bullet) At(p Point) bool {
	topLeftPoint := Point{b.position.x, b.position.y}
	bottomRightPoint := Point{
		x: b.position.x + b.drawing.width,
		y: b.position.y + b.drawing.height,
	}
	return InDimensions(p, topLeftPoint, bottomRightPoint)
}

func (g *Game) GetBulletAtPoint(p Point) (int, Bullet) {
	for i := len(g.bullets) - 1; i >= 0; i-- {
		bullet := g.bullets[i]
		if bullet.At(p) {
			return i, bullet
		}
	}
	return -1, Bullet{}
}
