package main

import "fmt"

func heroHeal(hero Hero, dragon Dragon, maxHealth int) int {
	health := hero.health
	fmt.Println(health)
	if health <= maxHealth-20 {
		health = health + 20
		fmt.Printf("You healed your hero to 20 HP\n")

	} else {
		fmt.Printf("Sorry, but you cant heal you hero, because HP is max. You automatically attack the dragon\n")
		dragonAttack(hero, dragon)
	}
	return health

}
