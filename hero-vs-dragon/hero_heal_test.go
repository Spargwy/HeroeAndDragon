package main

import (
	"fmt"
	"testing"
)

func TestHeroHeal(t *testing.T) {
	hero := Hero{
		maxHealth: 100,
		health:    40,
		damage:    20,
	}

	healthBeforeHeal := hero.health
	hero.health = heroHeal(hero, hero.maxHealth)
	if healthBeforeHeal >= hero.health && healthBeforeHeal != hero.maxHealth {
		t.Fatalf("Health after healing must be bigger than before. health before heal = %d health after heal = %d", healthBeforeHeal, hero.health)
	}
	fmt.Print("OK\n")

}
