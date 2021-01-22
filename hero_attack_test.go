package main

import (
	"testing"
)

func TestHeroAttack(t *testing.T) {
	hero := Hero{}
	hero.maxHealth = 100
	hero.health = 100
	hero.damage = 100

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
	hero.health = hero.health - 70
	damageReduction = heroTiredness(hero)
	returnedDragonHealth = heroAttack(hero, dragon, crossbow, true)
	wantDragonHealth = dragon.health - (crossbow.damage - damageReduction)
	if returnedDragonHealth != wantDragonHealth {
		t.Fatalf("ERROR! Dragon's health must be equal %d, but equal %d. WEAPON %s, %d, %d", wantDragonHealth, returnedDragonHealth, crossbow.name, crossbow.damage, crossbow.missChance)
	}
	hero.health = hero.health - 10
	damageReduction = heroTiredness(hero)
	returnedDragonHealth = heroAttack(hero, dragon, crossbow, true)
	wantDragonHealth = dragon.health - (crossbow.damage - damageReduction)
	if returnedDragonHealth != wantDragonHealth {
		t.Fatalf("ERROR! Dragon's health must be equal %d, but equal %d. WEAPON %s, %d, %d", wantDragonHealth, returnedDragonHealth, crossbow.name, crossbow.damage, crossbow.missChance)
	}
	hero.health = hero.health - 10
	damageReduction = heroTiredness(hero)
	returnedDragonHealth = heroAttack(hero, dragon, crossbow, true)
	wantDragonHealth = dragon.health - crossbow.minDamage
	if returnedDragonHealth != wantDragonHealth {
		t.Fatalf("ERROR! Dragon's health must be equal %d, but equal %d. WEAPON %s, %d, %d", wantDragonHealth, returnedDragonHealth, crossbow.name, crossbow.damage, crossbow.missChance)
	}
}
