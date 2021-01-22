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
		chooseWeapon := chooseWeapon()
		weapon := usingWeapon(chooseWeapon)
		chanceToAttack := heroChanceToAttack(weapon.missChance)
		dragon.health = heroAttack(hero, dragon, weapon, chanceToAttack)
	case "2":
		fmt.Println(hero.health, maxHealth)
		if hero.health <= maxHealth-20 {
			hero.health = heroHeal(hero, maxHealth)
		} else {
			fmt.Printf("Sorry, but you cant heal your hero, because HP is max. You automatically attack the dragon\n")
			chooseWeapon := chooseWeapon()
			weapon := usingWeapon(chooseWeapon)
			chanceToAttack := heroChanceToAttack(weapon.missChance)
			dragon.health = heroAttack(hero, dragon, weapon, chanceToAttack)
		}

	default:
		weapon := usingWeapon("sword")
		chanceToAttack := heroChanceToAttack(weapon.missChance)
		dragon.health = heroAttack(hero, dragon, weapon, chanceToAttack)
	}
	return hero, dragon
}
