package main

import "fmt"

func outputMessage(move int, hero Hero, dragon Dragon) {

	fmt.Printf("\n\n#++++++++++++++++++++++++++++++++++++\n")
	gameComments(hero, dragon)
	fmt.Printf("\n#move: %d\n", move)
	fmt.Printf("#Your Hero's health = %d. \n", hero.health)
	fmt.Printf("#Dragon's health = %d\n", dragon.health)
	fmt.Printf("#Hero's armor = %d\n", hero.armor)
	fmt.Printf("#++++++++++++++++++++++++++++++++++++\n\n")

	if hero.health <= 0 {
		fmt.Printf("Dragon win! The battle lasted %d moves!\n", move)
	} else if dragon.health <= 0 {
		fmt.Printf("Hero win! The battle lasted %d moves!\n", move)
	}
}
