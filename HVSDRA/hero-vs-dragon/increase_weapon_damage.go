package main

import "fmt"

func increaseWeaponDamage(weapon, crossbow, pan, sword Weapon) (Weapon, Weapon, Weapon) {
	var increaseDamage int
	if weapon.damage < 10 {
		increaseDamage = 1
	} else if weapon.damage >= 10 {
		increaseDamage = weapon.damage / 10
	}
	if weapon.name != "sword" { //sword is a standart weapon
		if weapon.damage-increaseDamage > 0 {
			weapon.damage -= increaseDamage
		} else if weapon.damage-increaseDamage <= 0 {
			increaseDamage = weapon.damage
			weapon.damage = weapon.damage - increaseDamage
			fmt.Print("This weapon is broke: !", weapon.name)
		}
	}
	if weapon.name == "crossbow" {
		crossbow.damage = weapon.damage
	} else if weapon.name == "pan" {
		pan.damage = weapon.damage
	}
	return crossbow, pan, sword
}
