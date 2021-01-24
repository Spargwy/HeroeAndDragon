package main

import (
	"fmt"
)

func events(dragon Dragon, hero Hero, chance int) (Hero, Dragon) {

	if chance <= 5 {
		fmt.Print("The cold snap affected the dragon(dragon hp reduced by 20 percent and -2 when hp < 10\n")
		if dragon.health > 10 {
			dragon.health = dragon.health - dragon.health/10*2
		} else if dragon.health < 10 {
			dragon.health -= 2
		}

	} else if chance > 5 && chance <= 10 {
		fmt.Print("Magnetic storms began, the hero had a headache(reduce damage by 20 percent and -2 wheh hp < 10 \n")
		if hero.health > 10 {
			hero.health = hero.health - hero.health/10*2
		} else if hero.health < 10 {
			hero.health -= 2
		}

	} else if chance > 10 && chance <= 15 {
		fmt.Print("You told the dragon that no one loves him. The dragon's health has decreased by 50 percent and if dragon hase smaller than 10hp, reduce = 4.")
		fmt.Print("\nAll is fair in war, but this is too much...\n")
		if dragon.health > 10 {
			dragon.health = dragon.health - dragon.health/10*5
		} else if dragon.health < 10 {
			dragon.health -= 4
		}
	} else if chance > 15 && chance <= 20 {
		fmt.Print("a lot of snow fell and a huge snowball flew into the dragon's mouth(dragon health reduces by 10 percent and -1 if hp<10\n")
		if dragon.health > 10 {
			dragon.health = dragon.health - dragon.health/10*1
		} else if dragon.health < 10 {
			dragon.health -= 1
		}
	} else if chance > 20 && chance <= 25 {
		fmt.Print("You are found the elixir! (Attack increased by 30 percent and +3 if damage < 10)\n")
		if hero.damage > 10 {
			hero.damage = hero.damage + hero.damage/10*3
		} else if hero.damage < 10 {
			hero.damage += 3
		}
	}
	return hero, dragon
}
