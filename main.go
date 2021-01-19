package main

import (
	"fmt"
)

type Hero struct {
	maxHealth int
	health    int
	damage    int
	armor     int
}

type Dragon struct {
	health      int
	damage      int
	missChacnce int
}

type Weapon struct {
	name         string
	damage       int
	numberOfUses int
	missChacnce  int
	usedTimes    int
}

func main() {

	pan := Weapon{}
	pan.name = "pan"
	pan.damage = 10
	pan.missChacnce = 0
	pan.numberOfUses = 2
	pan.usedTimes = 0

	crossbow := Weapon{}
	crossbow.name = "crossbow"
	crossbow.damage = 30
	crossbow.missChacnce = 70
	crossbow.numberOfUses = 3
	crossbow.usedTimes = 0

	standartSword := Weapon{}
	standartSword.name = "standart"
	standartSword.damage = 15
	standartSword.missChacnce = 50
	standartSword.numberOfUses = 1
	standartSword.usedTimes = 0

	hero := Hero{}
	hero.maxHealth = 20
	hero.health = 100
	hero.damage = 20
	hero.armor = 50

	dragon := Dragon{}
	dragon.health = 2000
	dragon.damage = 20
	dragon.missChacnce = 0

	var action string

	maxHealth := hero.maxHealth
	move := 1
	for hero.health > 0 && dragon.health > 0 {
		dragon.health, hero.health = naturalDisasters(dragon, hero)
		if dragon.health > 0 && hero.health > 0 {
			fmt.Printf("Choose the next action: ")
			fmt.Scanf("%s", &action)
			switch action {
			case "attack":
				fmt.Printf("If you want to use weapon, enter it's name: ")
				var weaponName string
				fmt.Scanf("%s", &weaponName)
				if weaponName == "pan" {
					weapon := usingWeapon(pan.usedTimes, pan.name, pan.damage)
					dragon.health, pan.damage = heroAttack(hero, dragon, weapon)
					pan.usedTimes++
				} else if weaponName == "crossbow" {
					weapon := usingWeapon(crossbow.usedTimes, crossbow.name, crossbow.damage)
					dragon.health, crossbow.damage = heroAttack(hero, dragon, weapon)

					crossbow.usedTimes++
					fmt.Println(crossbow.usedTimes)
				} else {
					weapon := usingWeapon(standartSword.usedTimes, standartSword.name, standartSword.damage) //можно не вызывать. Хочу, чтобы default выполнился. Я сним час мучился. Надо же себя утещить
					dragon.health, standartSword.damage = heroAttack(hero, dragon, weapon)
				}

			case "heal":
				hero.health = heroHeal(hero, dragon, maxHealth)

			default:
				fmt.Printf("If you will not choose the correct actions, heal will be execute automatically")
				if hero.armor <= 0 {
					hero.health = heroHeal(hero, dragon, maxHealth)
				} else {

				}

			}
			if dragon.health > 0 {
				hero.health, hero.armor = dragonAttack(hero, dragon)
			}

			fmt.Printf("\n\n\n#++++++++++++++++++++++++++++++++++++\n")
			fmt.Printf("#move: %d\n", move)
			fmt.Printf("#Your Hero's health = %d. \n", hero.health)
			fmt.Printf("#Dragon's health = %d\n", dragon.health)
			fmt.Printf("Your Hero's armor = %d. \n", hero.armor)
			fmt.Printf("#++++++++++++++++++++++++++++++++++++\n\n\n")

			if hero.health <= 0 {
				fmt.Printf("Dragon win! The battle lasted %d moves!\n", move)
			} else if dragon.health <= 0 {
				fmt.Printf("Hero win! The battle lasted %d moves!\n", move)
			}

			move++

		}
	}
}
