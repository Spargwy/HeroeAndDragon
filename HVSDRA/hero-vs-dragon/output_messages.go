package main

import "fmt"

func outputMessage(move int, heroHealth, dragonHealth int, heroArmor int) {

	fmt.Printf("\n\n#++++++++++++++++++++++++++++++++++++\n")
	fmt.Printf("#move: %d\n", move)
	fmt.Printf("#Your Hero's health = %d. \n", heroHealth)
	fmt.Printf("#Dragon's health = %d\n", dragonHealth)
	fmt.Printf("#Hero's armor = %d\n", heroArmor)
	fmt.Printf("#++++++++++++++++++++++++++++++++++++\n\n")

	if heroHealth <= 0 {
		fmt.Printf("Dragon win! The battle lasted %d moves!\n", move)
	} else if dragonHealth <= 0 {
		fmt.Printf("Hero win! The battle lasted %d moves!\n", move)
	}
}
