package main

import (
	"fmt"
	"math/rand"
	"time"
)

func heroAttack(hero Hero, dragon Dragon, weapon Weapon) (dragonHealth int, weaponDamage int) {
	var choose string
	var angryMode bool
	fmt.Print("Angry mode(yes/no):")
	fmt.Scanf("%s", &choose)
	if choose == "yes" {
		angryMode = true
	} else {
		angryMode = false
	}
	if angryMode == true {
		weapon.damage = weapon.damage + weapon.damage
		weapon.missChance = weapon.missChance + weapon.missChance/100*30
		fmt.Print("\n\n\n", weapon.damage, "\n\n\n")

	}

	rand.Seed(time.Now().UnixNano())
	chance := rand.Intn(100)
	weaponDamage = weapon.damage
	if hero.health <= hero.maxHealth/100*30 {
		weaponDamage = weaponDamage - weaponDamage/2
	}
	if chance <= 100-weapon.missChance {
		dragon.health = dragon.health - weaponDamage
		weapon.damage = weaponDamage - weaponDamage/10
	} else if chance >= 100-weapon.missChance {
		fmt.Printf("You missed with chance %d!\n", chance)
	}

	return dragon.health, weapon.damage
}
