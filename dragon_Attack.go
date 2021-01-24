package main

import (
	"fmt"
)

func dragonSuccessfullyAttack(daragonMissChance, chance int) bool {
	var successfulAttack bool
	if chance >= 100-daragonMissChance {
		fmt.Printf("Dragon missed with chance%d!\n", chance)
	} else {
		fmt.Printf("Dragon attack with chance %d!\n", chance)
		successfulAttack = true
	}
	return successfulAttack
}
func dragonAttack(hero Hero, dragon Dragon, successfulAttack bool) (int, int) {
	if successfulAttack == true {
		hero.armor, hero.health = damageToArmor(hero, dragon)
	}
	return hero.health, hero.armor

}
