package main

import (
	"fmt"
	"math/rand"
	"time"
)

func heroAttack(hero Hero, dragon Dragon, weapon Weapon) int {
	rand.Seed(time.Now().UnixNano())
	chance := rand.Intn(100)
	if chance <= 100-weapon.missChacnce {
		dragon.health = dragon.health - weapon.damage
	} else if chance > 100 {
		fmt.Printf("You missed with chance %d!\n", chance)
	}

	return dragon.health
}
