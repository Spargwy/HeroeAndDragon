package main

import (
	"fmt"
	"math/rand"
)

func dragonSuccessfullyAttack(daragonMissChance, seed int) bool {
	var successfulAttack bool
	rand.Seed(int64(seed))
	chance := rand.Intn(100)
	fmt.Println(chance)
	if chance >= 100-daragonMissChance {
		fmt.Printf("Dragon missed with chance%d!\n", chance)
	} else {
		fmt.Printf("Dragon attack with chance %d!\n", chance)
		successfulAttack = true
	}
	return successfulAttack
}
func dragonAttack(hero Hero, dragon Dragon, successfulAttack bool) int {
	if successfulAttack == true {
		hero.health = hero.health - dragon.damage
	}
	return hero.health

}
