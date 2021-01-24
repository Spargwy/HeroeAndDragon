package main

type Hero struct {
	maxHealth int
	health    int
	damage    int
}

type Dragon struct {
	health     int
	damage     int
	missChance int
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
	dragon.missChance = 0

	maxHealth := hero.maxHealth
	move := 1
	for hero.health > 0 && dragon.health > 0 {
		hero, dragon = chooseAction(hero, dragon, maxHealth)
		if dragon.health > 0 {
			chanceIsGood := dragonChanceToAttack(dragon.missChance)
			hero.health = dragonAttack(hero, dragon, chanceIsGood)
		}
		outputMessage(move, hero.health, dragon.health)
		move++

	}

}
