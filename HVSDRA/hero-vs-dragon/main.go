package main

import "fmt"

type Hero struct {
	maxHealth int
	health    int
	damage    int
	armor     int
}

type Dragon struct {
	health     int
	damage     int
	missChance int
}

func main() {
	hero := Hero{
		maxHealth: 1000,
		health:    10000,
		damage:    20,
		armor:     50,
	}
	dragon := Dragon{
		health:     1000,
		damage:     40,
		missChance: 30,
	}
	crossbow := Weapon{
		name:          "crossbow",
		damage:        30,
		minDamage:     0,
		missChance:    20,
		numberOfUsing: 5,
		ItemUsed:      0,
	}
	pan := Weapon{
		name:          "pan",
		damage:        10,
		missChance:    0,
		numberOfUsing: 10,
		ItemUsed:      0,
	}
	sword := Weapon{
		name:          "sword",
		damage:        15,
		missChance:    20,
		numberOfUsing: 1,
		ItemUsed:      0,
	}
	maxHealth := hero.maxHealth
	move := 1
	for hero.health > 0 && dragon.health > 0 {
		action := askAboutnextAction()
		hero, dragon, crossbow, pan, sword = chooseAction(crossbow, pan, sword, hero, dragon, maxHealth, action)
		if dragon.health > 0 {
			chance := chanceToAttack()
			chanceIsGood := dragonSuccessfullyAttack(dragon.missChance, chance)
			hero.health, hero.armor = dragonAttack(hero, dragon, chanceIsGood)
		}
		fmt.Println(crossbow)
		outputMessage(move, hero, dragon)
		move++

	}

}
