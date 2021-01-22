package main

import (
	"fmt"
	"math/rand"
)

func heroChanceToAttack(weaponMissChance int, seed int) bool {
	var successfulAttack bool
	rand.Seed(int64(seed))
	chance := rand.Intn(100)
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
