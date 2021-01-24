package main

import "fmt"

func damageToArmor(hero Hero, dragon Dragon) (heroArmor int, heroHealth int) {
	damageToArmor := dragon.damage
	if hero.armor-damageToArmor > 0 {
		hero.armor -= damageToArmor
	} else if hero.armor-damageToArmor < 0 {
		damageToArmor = (0 - hero.armor) * (-1)
		dragon.damage = dragon.damage - damageToArmor
		hero.health = hero.health - dragon.damage
		hero.armor -= damageToArmor
		fmt.Print("Your hero's armor is destroyd! Now, dragon's attack will be reduce your health!")
	}
	if hero.armor == 0 {

	}

	return hero.armor, hero.health
}
