package main

import (
	"fmt"
	"math/rand"
	"time"
)

func heroAttack(hero Hero, dragon Dragon, weapon Weapon) int {
	damageReduction := heroTiredness(hero)
	if weapon.minDamage > weapon.damage-damageReduction {
		hero.damage = weapon.minDamage
	} else {
		hero.damage = weapon.damage - damageReduction
	}
	rand.Seed(time.Now().UnixNano())
	chance := rand.Intn(100)
	if chance <= 100-weapon.missChance {
		dragon.health = dragon.health - hero.damage
	} else {
		fmt.Printf("You missed with chance %d!\n", chance)
	}
	fmt.Println(hero.damage)
	return dragon.health
}
