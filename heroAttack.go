package main

import (
	"fmt"
	"math/rand"
	"time"
)

func heroAttack(hero Hero, dragon Dragon, weapon Weapon) int {
	rand.Seed(time.Now().UnixNano())
	chance := rand.Intn(100)
	weaponDamage := weapon.damage
	if hero.health <= hero.maxHealth/100*30 {
		weaponDamage = weaponDamage - weaponDamage/2
	}
	if chance <= 100-weapon.missChacnce {
		dragon.health = dragon.health - weaponDamage
	} else if chance >= 100-weapon.missChacnce {
		fmt.Printf("You missed with chance %d!\n", chance)
	}

	return dragon.health
}
