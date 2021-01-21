package main

import (
	"fmt"
	"math/rand"
	"time"
)

func dragonAttack(hero Hero, dragon Dragon) int {
	rand.Seed(time.Now().UnixNano())
	chance := rand.Intn(100)
	if chance <= 50 {
		hero.health = hero.health - dragon.damage
	} else {
		fmt.Printf("Dragon missed with chance %d!", chance)
	}
	return hero.health

}
