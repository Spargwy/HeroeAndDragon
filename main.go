package main

import (
	"fmt"
)

type Hero struct {
	health int
	damage int
}

type Dragon struct {
	health      int
	damage      int
	missChacnce int
}

type Weapon struct {
	name         string
	damage       int
	numberOfUses int
	missChacnce  int
	usedTimes    int
}

func main() {

	pan := Weapon{}
	pan.name = "pan"
	pan.damage = 10
	pan.missChacnce = 0
	pan.numberOfUses = 2
	pan.usedTimes = 0

	crossbow := Weapon{}
	crossbow.name = "crossbow"
	crossbow.damage = 30
	crossbow.missChacnce = 70
	crossbow.numberOfUses = 2
	crossbow.usedTimes = 0

	standartSword := Weapon{}
	standartSword.name = "standart"
	standartSword.damage = 20
	standartSword.missChacnce = 50
	standartSword.numberOfUses = 1
	standartSword.usedTimes = 0

	hero := Hero{}
	hero.health = 100
	hero.damage = 20

	dragon := Dragon{}
	dragon.health = 100
	dragon.damage = 20
	dragon.missChacnce = 50

	var action string
	maxHealth := hero.health
	move := 1
	for hero.health > 0 && dragon.health > 0 {
		fmt.Printf("Choose the next action: ")
		fmt.Scanf("%s", &action)
		switch action {
		case "attack":
			fmt.Printf("If you want to use weapon, enter it's name: ")
			weapon := usingWeapon(pan.usedTimes)
			dragon.health = heroAttack(hero, dragon, weapon)
			pan.usedTimes++
		case "heal":
			hero.health = heroHeal(hero, dragon, maxHealth)

		}
		hero.health = dragonAttack(hero, dragon)

		fmt.Printf("\n\n\n#++++++++++++++++++++++++++++++++++++\n")
		fmt.Printf("#move: %d\n", move)
		fmt.Printf("#Your Hero's health = %d. \n", hero.health)
		fmt.Printf("#Dragon's health = %d\n", dragon.health)
		fmt.Printf("#++++++++++++++++++++++++++++++++++++\n\n\n")

		if hero.health <= 0 {
			fmt.Printf("Dragon win! The battle lasted %d moves!\n", move)
		} else if dragon.health <= 0 {
			fmt.Printf("Hero win! The battle lasted %d moves!\n", move)
		}

		move++

	}

}
