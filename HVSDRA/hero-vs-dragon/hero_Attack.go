package main

import (
	"fmt"
)

func heroSuccessfullyAttack(weaponMissChance int, chance int) bool {
	TurnOnAngryModeOrNo := askAboutAngryParameter()
	_, reduceMissChance := angryMode(TurnOnAngryModeOrNo)
	var successfulAttack bool
	if chance > 100-weaponMissChance-reduceMissChance {
		fmt.Printf("You missed with chance %d!\n", chance)
	} else {
		successfulAttack = true
	}

	return successfulAttack
}

func askAboutAngryParameter() string {
	fmt.Print("Do you want to use the anger parameter(+30 percent damage and 20 percent misschance for 1 move)?")
	fmt.Print("Yes: 1\n No: 2\n Use: ")
	var useAngerParameterOrNo string
	fmt.Scan(&useAngerParameterOrNo)

	return useAngerParameterOrNo
}

func angryMode(TurnOnAngryModeOrNo string) (reduceDamage, reduceMissChance int) {
	TurnOnAngryModeOrNo = askAboutAngryParameter()

	if TurnOnAngryModeOrNo == "1" {
		reduceDamage = 30
		reduceMissChance = 20
	} else if TurnOnAngryModeOrNo == "2" {
		reduceDamage = 0
		reduceMissChance = 0
	} else {
		angryMode(TurnOnAngryModeOrNo)
	}
	return reduceDamage, reduceMissChance

}

func heroAttack(hero Hero, dragon Dragon, crossbow, pan, sword, weapon Weapon, successfulAttack bool) (int, Weapon, Weapon, Weapon, Weapon) {
	damageReduction := heroTiredness(hero)
	if weapon.minDamage > weapon.damage-damageReduction {
		hero.damage = weapon.minDamage
	} else {
		hero.damage = weapon.damage - damageReduction
	}
	if weapon.damage > 0 {
		if successfulAttack == true {
			dragon.health = dragon.health - hero.damage

			crossbow, pan, sword = increaseWeaponDamage(weapon, crossbow, pan, sword)
		}
	} else if weapon.damage <= 0 {
		fmt.Printf("WEAPON IS BROKEN %s \n", weapon.name)
		choosedWeapon := chooseWeapon()
		usingWeapon, _, _ := usingWeapon(crossbow, pan, sword, choosedWeapon)
		heroAttack(hero, dragon, crossbow, pan, sword, usingWeapon, successfulAttack)
	}

	return dragon.health, weapon, crossbow, pan, sword
}
