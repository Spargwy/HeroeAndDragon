package main

import (
	"fmt"
)

func heroChanceToAttack(weaponMissChance int, chance int) bool {
	var successfulAttack bool
	if chance > 100-weaponMissChance {
		fmt.Printf("You missed with chance %d!\n", chance)
	} else {
		successfulAttack = true
	}

	return successfulAttack
}

func heroAttack(hero Hero, dragon Dragon, weapon Weapon, chanceIsGood bool) int {
	damageReduction := heroTiredness(hero)
	if weapon.minDamage > weapon.damage-damageReduction {
		hero.damage = weapon.minDamage
	} else {
		hero.damage = weapon.damage - damageReduction
	}
	if chanceIsGood == true {
		dragon.health = dragon.health - hero.damage
	}

	return dragon.health
}
