package main

import (
	"testing"
)

func TestHeroAttack(t *testing.T) {
	hero := Hero{}
	hero.maxHealth = 100
	hero.health = 40
	hero.damage = 100

	hero2 := Hero{}
	hero.maxHealth = 100
	hero.health = 30
	hero.damage = 20

	dragon := Dragon{}
	dragon.health = 1000
	dragon.damage = 5

	crossbow := Weapon{}
	crossbow.name = "crossbow"
	crossbow.damage = 30
	crossbow.minDamage = 15
	crossbow.missChance = 0
	crossbow.numberOfUsing = 100
	crossbow.ItemUsed = 0

	pan := Weapon{}
	pan.name = "pan"
	pan.damage = 10
	pan.missChance = 100
	pan.numberOfUsing = 10
	pan.ItemUsed = 0

	sword := Weapon{}
	sword.name = "sword"
	sword.damage = 30
	sword.missChance = 0
	sword.numberOfUsing = 1
	sword.ItemUsed = 0

	damageReduction := heroTiredness(hero)
	returnedDragonHealth := heroAttack(hero, dragon, crossbow, true)
	wantDragonHealth := dragon.health - (crossbow.damage - damageReduction)
	if returnedDragonHealth != wantDragonHealth {
		t.Fatalf("ERROR! Dragon's health must be equal %d, but equal %d", wantDragonHealth, returnedDragonHealth)
	}
	returnedDragonHealth = heroAttack(hero2, dragon, crossbow, true)
	wantDragonHealth = dragon.health - (crossbow.damage - 5)
	if returnedDragonHealth != wantDragonHealth {
		t.Fatalf("ERROR! Dragon's health must be equal %d, but equal %d", wantDragonHealth, returnedDragonHealth)
	}

}
