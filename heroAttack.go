package main

import (
	"fmt"
	"math/rand"
	"time"
)

func heroAttack(hero Hero, dragon Dragon, weapon Weapon) (dragonHealth int, weaponDamage int) {
	rand.Seed(time.Now().UnixNano())
	chance := rand.Intn(100)
	weaponDamage = weapon.damage
	if hero.health <= hero.maxHealth/100*30 {
		weaponDamage = weaponDamage - weaponDamage/2
	}
	if chance <= 100-weapon.missChacnce {
		dragon.health = dragon.health - weaponDamage
		weapon.damage = weaponDamage - weaponDamage/10
	} else if chance >= 100-weapon.missChacnce {
		fmt.Printf("You missed with chance %d!\n", chance)
	}

	return dragon.health, weapon.damage
}
