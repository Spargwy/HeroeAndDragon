package main

import "fmt"

func heroHeal(hero Hero, maxHealth int) int {

	health := hero.health
	regen := 20
	if hero.health+20 < hero.maxHealth {
		health += regen
		fmt.Printf("You healed your hero to 20 HP\n")
	} else {
		regen = 100 - hero.health
		health += regen
		fmt.Printf("Heal = %d", regen)
	}
	return health

}
