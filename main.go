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
	hero := Hero{
		maxHealth: 100,
		health:    40,
		damage:    20,
	}
	dragon := Dragon{
		health:     1000,
		damage:     5,
		missChance: 0,
	}
	maxHealth := hero.maxHealth
	move := 1
	for hero.health > 0 && dragon.health > 0 {
		action := askAboutnextAction()
		hero, dragon = chooseAction(hero, dragon, maxHealth, action)
		if dragon.health > 0 {
			chance := chanceToAttack()
			chanceIsGood := dragonSuccessfullyAttack(dragon.missChance, chance)
			hero.health = dragonAttack(hero, dragon, chanceIsGood)
		}
		outputMessage(move, hero.health, dragon.health)
		move++

	}

}
