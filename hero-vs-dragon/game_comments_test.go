package main

import "testing"

func TestGameComments(t *testing.T) {
	hero := Hero{
		maxHealth: 100,
		health:    25,
		damage:    20,
		armor:     50,
	}
	dragon := Dragon{
		health:     1000,
		damage:     20,
		missChance: 20,
	}

	hero.health = 20
	wantComment := "Hero's health equal" + string(hero.health) + "! One more hit and dragon will be win! Maybe it's time to heal?"
	comment := gameComments(hero, dragon)
	if wantComment != comment {
		t.Fatalf("ERROR!!! Comment must be equal %s, but equal %s", wantComment, comment)
	}
	hero.health = 29
	wantComment = "The hero is weak, but he still has a chance"
	comment = gameComments(hero, dragon)
	if wantComment != comment {
		t.Fatalf("ERROR!!! Comment must be equal %s, but equal %s", wantComment, comment)
	}
	hero.health = 35
	wantComment = "The dragon inflicted some injuries on the hero, but these are all trifles!"
	comment = gameComments(hero, dragon)
	if wantComment != comment {
		t.Fatalf("ERROR!!! Comment must be equal %s, but equal %s", wantComment, comment)
	}

	hero.health = hero.maxHealth
	wantComment = "Congratulations! Your hero has maximum health. Now you can fight with new forces!"
	comment = gameComments(hero, dragon)
	if wantComment != comment {
		t.Fatalf("ERROR!!! Comment must be equal %s, but equal %s", wantComment, comment)
	}
}
