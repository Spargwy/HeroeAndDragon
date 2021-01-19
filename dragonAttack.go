package main

import (
	"fmt"
	"math/rand"
	"time"
)

func dragonAttack(hero Hero, dragon Dragon) int {
	rand.Seed(time.Now().UnixNano())
	chance := rand.Intn(100)
	if chance <= 100-dragon.missChacnce {
		hero.health = hero.health - dragon.damage
	} else {
		fmt.Printf("Dragon missed with chance %d!\n", chance)
	}
	return hero.health
}
