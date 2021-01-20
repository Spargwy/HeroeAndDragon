package main

import "fmt"

func comments(hero Hero, dragon Dragon) {
	if hero.health < hero.maxHealth/10*5 {
		fmt.Print("Герой устал, но полон решительности\n")
	} else if hero.health < hero.maxHealth/10*4 {
		fmt.Print("Силы со временем битвы покидают героя...\n")
	} else if hero.health < hero.maxHealth/10*3 {
		fmt.Print("Герой истощен, но ещё в состоянии продолжать битву!\n")
	} else if hero.health < hero.maxHealth/10*2 {
		fmt.Print("Герой борется на грани своих возможностей!\n")
	} else if hero.health <= 0 {
		fmt.Print("R.I.P.\n")
	}

	if dragon.health < dragon.maxHealth/10*5 {
		fmt.Print("Слишком стрёмно проигрывать какому-то чуваку в железе\n")
	} else if dragon.health < dragon.maxHealth/10*4 {
		fmt.Print("Ну не, ну нееееееее, бред какой-то\n")
	} else if dragon.health < dragon.maxHealth/10*3 {
		fmt.Print("Дракон выбивается из сил")
	} else if dragon.health < dragon.maxHealth/10*2 {
		fmt.Print("Ещё немного и кто-то отобедует дракошей...\n")
	} else if hero.health <= 0 {
		fmt.Print("https://deezer.page.link/F7wYcc1Voe9ZiCHP7\n")
	}
}
