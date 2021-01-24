package main

func damageToArmor(hero Hero, dragon Dragon) (heroArmor int, heroHealth int) {
	damageToArmor := dragon.damage
	if hero.armor-damageToArmor > 0 {
		hero.armor -= damageToArmor
	} else {
		damageToArmor = (0 - hero.armor) * (-1)
		dragon.damage = dragon.damage - damageToArmor
		hero.health = hero.health - dragon.damage
		hero.armor -= damageToArmor
	}

	return hero.armor, hero.health
}
