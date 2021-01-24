package main

func heroTiredness(hero Hero) int {
	damageReduction := 0
	if hero.health <= (hero.maxHealth/100*30) && hero.health > hero.maxHealth/100*20 {
		damageReduction = 5
	} else if hero.health <= hero.maxHealth/100*20 && hero.health > hero.maxHealth/100*10 {
		damageReduction = 10
	} else if hero.health <= hero.maxHealth/100*10 {
		damageReduction = 20
	} else {
		damageReduction = 0
	}

	return damageReduction
}
