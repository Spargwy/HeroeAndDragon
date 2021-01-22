package main

import (
	"fmt"
	"math/rand"
	"time"
)

func heroChanceToAttack(weaponMissChance int) bool {
	var chanceIsGood bool
	rand.Seed(time.Now().UnixNano())
	chance := rand.Intn(100)
	if chance >= 100-weaponMissChance {
		fmt.Printf("You missed with chance %d!\n", chance)
	} else {
		chanceIsGood = true
	}

	return chanceIsGood
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
