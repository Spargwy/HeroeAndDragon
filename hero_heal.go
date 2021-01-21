package main

import "fmt"

func heroHeal(hero Hero, maxHealth int) int {
	health := hero.health
	health = health + 20
	fmt.Printf("You healed your hero to 20 HP\n")

	return health

}
