package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Hero struct {
	health int
	damage int
}

type Dragon struct {
	health int
	damage int
}

func heroattack(hero Hero, dragon Dragon) int {
	rand.Seed(time.Now().UnixNano())
	chance := rand.Intn(100)
	if chance <= 50 {
		dragon.health = dragon.health - hero.damage
	} else {
		fmt.Printf("You missed with chance %d!\n", chance)
	}
	return dragon.health
}

func dragonAttack(hero Hero, dragon Dragon) int {
	rand.Seed(time.Now().UnixNano())
	chance := rand.Intn(100)
	if chance <= 40 {
		hero.health = hero.health - dragon.damage
	} else {
		fmt.Printf("Dragon missed with chance %d!\n", chance)
	}
	return hero.health
}

func heroHeal(hero Hero, dragon Dragon, maxHealth int) int {
	health := hero.health
	if health <= maxHealth-20 {
		health = health + 20
		fmt.Printf("You healed your hero to 20 HP\n")

	} else {
		fmt.Printf("Sorry, but you cant heal you hero, because HP is max. You automatically attack the dragon\n")
		dragonAttack(hero, dragon)
	}
	return health

}

func main() {
	hero := Hero{}
	hero.health = 20
	hero.damage = 20

	dragon := Dragon{}
	dragon.health = 20
	dragon.damage = 20

	var action string
	maxHealth := hero.health
	move := 1
	for hero.health > 0 && dragon.health > 0 {
		fmt.Printf("Choose the next action: ")
		fmt.Scanf("%s", &action)
		switch action {
		case "attack":
			dragon.health = heroattack(hero, dragon)
		case "heal":
			hero.health = heroHeal(hero, dragon, maxHealth)

		}
		if dragon.health > 0 {
			hero.health = dragonAttack(hero, dragon)
		}
		fmt.Printf("#++++++++++++++++++++++++++++++++++++\n")
		fmt.Printf("#move: %d\n", move)
		fmt.Printf("#Your Hero's health = %d. \n", hero.health)
		fmt.Printf("#Dragon's health = %d\n", dragon.health)
		fmt.Printf("#++++++++++++++++++++++++++++++++++++\n")

		if hero.health <= 0 {
			fmt.Printf("Dragon win! The battle lasted %d moves!\n", move)
		} else if dragon.health <= 0 {
			fmt.Printf("Hero win! The battle lasted %d moves!\n", move)
		}
		move++

	}

}
