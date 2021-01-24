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
		maxHealth: 100,
		health:    40,
		damage:    20,
		armor:     50,
	}
	dragon := Dragon{
		health:     1000,
		damage:     5,
		missChance: 0,
	}
	crossbow := Weapon{
		name:          "crossbow",
		damage:        30,
		minDamage:     15,
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
		damage:        30,
		missChance:    20,
		numberOfUsing: 1,
		ItemUsed:      0,
	}
	maxHealth := hero.maxHealth
	move := 1
	for hero.health > 0 && dragon.health > 0 {
		action := askAboutnextAction()
		hero, dragon, crossbow, pan = chooseAction(crossbow, pan, sword, hero, dragon, maxHealth, action)
		fmt.Println(crossbow, pan)
		if dragon.health > 0 {
			chance := chanceToAttack()
			chanceIsGood := dragonSuccessfullyAttack(dragon.missChance, chance)
			hero.health, hero.armor = dragonAttack(hero, dragon, chanceIsGood)
		}
		outputMessage(move, hero.health, dragon.health, hero.armor)
		move++

	}

}
