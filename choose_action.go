package main

import "fmt"

func chooseAction(hero Hero, dragon Dragon, maxHealth int) (Hero, Dragon) {
	var action string
	fmt.Print("Enter number of the next action\n")
	fmt.Print("attack: 1\n")
	fmt.Print("heal: 2\n")
	fmt.Scan(&action)
	switch action {
	case "1":
		weapon := usingWeapon()
		dragon.health = heroAttack(hero, dragon, weapon)
	case "2":
		fmt.Println(hero.health, maxHealth)
		if hero.health <= maxHealth-20 {
			hero.health = heroHeal(hero, dragon, maxHealth)
		} else {
			fmt.Printf("Sorry, but you cant heal your hero, because HP is max. You automatically attack the dragon\n")
			weapon := usingWeapon()
			dragon.health = heroAttack(hero, dragon, weapon)
		}

	default:
		weapon := usingWeapon()
		dragon.health = heroAttack(hero, dragon, weapon)
	}
	return hero, dragon
}
