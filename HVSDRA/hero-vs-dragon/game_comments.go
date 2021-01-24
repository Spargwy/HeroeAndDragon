package main

func gameComments(hero Hero, dragon Dragon) (string, string) {
	var commentsAboutHero string
	var commentsAboutDragon string
	if hero.health <= dragon.damage {
		commentsAboutHero = "Hero's health equal" + string(hero.health) + "! One more hit and dragon will be win! Maybe it's time to heal?"
	} else if hero.health > hero.maxHealth/10*2 && hero.health <= hero.maxHealth/10*3 {
		commentsAboutHero = "The hero is weak, but he still has a chance"
	} else if hero.health > hero.maxHealth/10*3 && hero.health <= hero.maxHealth/10*4 {
		commentsAboutHero = "The dragon inflicted some injuries on the hero, but these are all trifles!"
	} else if hero.health == hero.maxHealth {
		commentsAboutHero = "Congratulations! Your hero has maximum health. Now you can fight with new forces!"
	}
	if dragon.health <= dragon.maxHealth/10 {
		commentsAboutDragon = "the dragon has little health. A little more and he will be defeated"
	}

	return commentsAboutHero, commentsAboutDragon

}
