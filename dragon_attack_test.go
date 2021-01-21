package main

import "testing"

func TestDragonAttack(t *testing.T) {
	hero := Hero{}
	hero.maxHealth = 100
	hero.health = 40
	hero.damage = 20

	dragon := Dragon{}
	dragon.health = 1000
	dragon.damage = 5

	hero.health = dragonAttack(hero, dragon, true)
	if hero.health > hero.maxHealth {
		t.Fatalf(`dragonAttack(hero, dragon) = %d, hero health must be lower or equal max health`, hero.health)
	}

}
