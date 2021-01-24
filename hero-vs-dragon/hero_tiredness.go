package main

func heroTiredness(hero Hero) int {
	increaseMissChance := 0
	if hero.health <= (hero.maxHealth/100*30) && hero.health > hero.maxHealth/100*20 {
		increaseMissChance = 5
	} else if hero.health <= hero.maxHealth/100*20 && hero.health > hero.maxHealth/100*10 {
		increaseMissChance = 10
	} else if hero.health <= hero.maxHealth/100*10 {
		increaseMissChance = 20
	} else {
		increaseMissChance = 0
	}

	return increaseMissChance
}
