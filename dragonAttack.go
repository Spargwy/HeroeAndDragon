package main

import (
	"fmt"
	"math/rand"
	"time"
)

func dragonAttack(hero Hero, dragon Dragon) (int, int) {
	rand.Seed(time.Now().UnixNano())
	chance := rand.Intn(100)
	if chance <= 100-dragon.missChacnce {
		if hero.armor > 0 && (hero.armor-dragon.damage) > 0 {
			hero.armor = hero.armor - dragon.damage
		} else if hero.armor-dragon.damage < 0 {
			hero.health = hero.health - (dragon.damage - hero.armor)
			hero.armor = 0
			//hero.health = hero.health - (dragon.damage + hero.armor*-1)
			//hero.armor = hero.armor*-1 + hero.armor
		} else {
			hero.health = hero.health - dragon.damage
		}
	} else {
		fmt.Printf("Dragon missed with chance %d!\n", chance)
	}
	return hero.health, hero.armor
}
