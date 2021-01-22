package main

import (
	"fmt"
	"math/rand"
	"time"
)

func dragonChanceToAttack(dragonMissChance int) bool {
	var chanceIsGood bool
	rand.Seed(time.Now().UnixNano())
	chance := rand.Intn(100)
	if chance >= 100-dragonMissChance {
		fmt.Printf("Dragon missed with chance %d!\n", chance)
	} else {
		chanceIsGood = true
	}

	return chanceIsGood
}
func dragonAttack(hero Hero, dragon Dragon, chanceIsGood bool) int {
	if chanceIsGood == true {
		hero.health = hero.health - dragon.damage
	}
	return hero.health

}
