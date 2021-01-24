package main

import (
	"fmt"
	"testing"
)

func TestHeroAttack(t *testing.T) {
	hero := Hero{
		maxHealth: 100,
		health:    100,
		damage:    100,
	}
	dragon := Dragon{
		health: 1000,
		damage: 5,
	}
	crossbow := Weapon{
		name:          "crossbow",
		damage:        30,
		minDamage:     15,
		missChance:    20,
		numberOfUsing: 100,
		ItemUsed:      0,
	}

	damageReduction := heroTiredness(hero)
	returnedDragonHealth, _ := heroAttack(hero, dragon, crossbow, true)
	wantDragonHealth := dragon.health - (crossbow.damage - damageReduction)
	if returnedDragonHealth != wantDragonHealth {
		t.Fatalf("ERROR! Dragon's health must be equal %d, but equal %d", wantDragonHealth, returnedDragonHealth)
	}
	hero.health = hero.health - 70
	damageReduction = heroTiredness(hero)
	returnedDragonHealth, _ = heroAttack(hero, dragon, crossbow, true)
	wantDragonHealth = dragon.health - (crossbow.damage - damageReduction)
	if returnedDragonHealth != wantDragonHealth {
		t.Fatalf("ERROR! Dragon's health must be equal %d, but equal %d. WEAPON %s, %d, %d", wantDragonHealth, returnedDragonHealth, crossbow.name, crossbow.damage, crossbow.missChance)
	}
	hero.health = hero.health - 10
	damageReduction = heroTiredness(hero)
	returnedDragonHealth, _ = heroAttack(hero, dragon, crossbow, true)
	wantDragonHealth = dragon.health - (crossbow.damage - damageReduction)
	if returnedDragonHealth != wantDragonHealth {
		t.Fatalf("ERROR! Dragon's health must be equal %d, but equal %d. WEAPON %s, %d, %d", wantDragonHealth, returnedDragonHealth, crossbow.name, crossbow.damage, crossbow.missChance)
	}
	hero.health = hero.health - 10
	damageReduction = heroTiredness(hero)
	returnedDragonHealth, _ = heroAttack(hero, dragon, crossbow, true)
	wantDragonHealth = dragon.health - crossbow.minDamage
	if returnedDragonHealth != wantDragonHealth {
		t.Fatalf("ERROR! Dragon's health must be equal %d, but equal %d. WEAPON %s, %d, %d", wantDragonHealth, returnedDragonHealth, crossbow.name, crossbow.damage, crossbow.missChance)
	}
	fmt.Print("OK\n")

}
