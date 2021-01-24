package main

import (
	"fmt"
)

func heroSuccessfullyAttack(weaponMissChance int, chance int) bool {
	var successfulAttack bool
	if chance > 100-weaponMissChance {
		fmt.Printf("You missed with chance %d!\n", chance)
	} else {
		successfulAttack = true
	}

	return successfulAttack
}

func heroAttack(hero Hero, dragon Dragon, weapon Weapon, successfulAttack bool) (int, Weapon) {
	damageReduction := heroTiredness(hero)
	if weapon.minDamage > weapon.damage-damageReduction {
		hero.damage = weapon.minDamage
	} else {
		hero.damage = weapon.damage - damageReduction
	}
	if successfulAttack == true {
		if weapon.damage > 0 {
			dragon.health = dragon.health - hero.damage
			weapon = increaseWeaponDamage(weapon)
		} else if weapon.damage <= 0 {
			fmt.Print("WEAPON IS BROKEN")
		}
	}

	fmt.Println("WEAPON : \n", weapon)
	return dragon.health, weapon
}
