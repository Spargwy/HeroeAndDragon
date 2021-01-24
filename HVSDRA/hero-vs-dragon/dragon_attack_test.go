package main

import (
	"fmt"
	"testing"
)

func TestDragonAttack(t *testing.T) {
	hero := Hero{
		maxHealth: 100,
		health:    40,
		damage:    20,
		armor:     0,
	}
	dragon := Dragon{
		health: 1000,
		damage: 5,
	}
	wantHeroHealth := hero.health - 5
	hero.health, hero.armor = dragonAttack(hero, dragon, true)
	if hero.health != wantHeroHealth {
		t.Fatalf("ERROR! Hero health must be euqal %d, but equal %d", wantHeroHealth, hero.health)
	}

	fmt.Print("OK\n")

}
