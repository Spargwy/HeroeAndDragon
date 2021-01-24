package main

func gameComments(hero Hero, dragon Dragon) string {
	var comment string
	if hero.health <= dragon.damage {
		comment = "Hero's health equal" + string(hero.health) + "! One more hit and dragon will be win! Maybe it's time to heal?"
	} else if hero.health > hero.maxHealth/10*2 && hero.health <= hero.maxHealth/10*3 {
		comment = "The hero is weak, but he still has a chance"
	} else if hero.health > hero.maxHealth/10*3 && hero.health <= hero.maxHealth/10*4 {
		comment = "The dragon inflicted some injuries on the hero, but these are all trifles!"
	} else if hero.health == hero.maxHealth {
		comment = "Congratulations! Your hero has maximum health. Now you can fight with new forces!"
	}
	return comment

}
