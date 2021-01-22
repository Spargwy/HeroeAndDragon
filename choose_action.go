package main

import "fmt"

func askAboutnextAction() string {
	var action string
	fmt.Print("Enter number of the next action\n")
	fmt.Print("attack: 1\n")
	fmt.Print("heal: 2\n")
	fmt.Scan(&action)

	return action
}

func chooseAction(hero Hero, dragon Dragon, maxHealth int, action string) (Hero, Dragon) {

	switch action {
	case "1":
		chooseWeapon := chooseWeapon()
		weapon := usingWeapon(chooseWeapon)
		seed := randomSeed()
		chanceToAttack := heroChanceToAttack(weapon.missChance, seed)
		dragon.health = heroAttack(hero, dragon, weapon, chanceToAttack)
	case "2":
		fmt.Println(hero.health, maxHealth)
		if hero.health < maxHealth {
			hero.health = heroHeal(hero, maxHealth)
		} else {
			fmt.Printf("Sorry, but you cant heal your hero, because HP is max. Please, choose another action\n")
			action = askAboutnextAction()
			chooseAction(hero, dragon, maxHealth, action)

		}

	default:
		fmt.Printf("Please, choose the correct action\n")
		action = askAboutnextAction()
		chooseAction(hero, dragon, maxHealth, action)
	}
	return hero, dragon
}
