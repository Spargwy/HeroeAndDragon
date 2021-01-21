package main

type Hero struct {
	maxHealth int
	health    int
	damage    int
}

type Dragon struct {
	health int
	damage int
}

var crossbowItemUsed int // переменные использующиеся как счетчик использования оружия
var panItemUsed int

func main() {
	hero := Hero{}
	hero.maxHealth = 100
	hero.health = 40
	hero.damage = 20

	dragon := Dragon{}
	dragon.health = 1000
	dragon.damage = 5

	maxHealth := hero.maxHealth
	move := 1
	for hero.health > 0 && dragon.health > 0 {
		hero, dragon = chooseAction(hero, dragon, maxHealth)
		if dragon.health > 0 {
			hero.health = dragonAttack(hero, dragon)
		}
		outputMessage(move, hero.health, dragon.health)
		move++

	}

}
