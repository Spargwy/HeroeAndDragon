package main

import "fmt"

func increaseWeaponDamage(weapon Weapon) Weapon {
	increaseDamage := weapon.damage / 10
	if weapon.name != "sword" {
		if weapon.damage-increaseDamage > 0 {
			weapon.damage -= increaseDamage
		} else if weapon.damage-increaseDamage <= 0 {
			increaseDamage = (0 - weapon.damage) * -1
			weapon.damage = weapon.damage - increaseDamage
			fmt.Print("This weapon is broke: !", weapon.name)
		}
	}
	return weapon
}
