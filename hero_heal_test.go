package main

import "testing"

func TestHeroHeal(t *testing.T) {
	hero := Hero{}
	hero.maxHealth = 100
	hero.health = 40
	hero.damage = 20

	healthBeforeHeal := hero.health
	hero.health = heroHeal(hero, hero.maxHealth)
	if healthBeforeHeal >= hero.health && healthBeforeHeal != hero.maxHealth {
		t.Fatalf("Health after healing must be bigger than before. health before heal = %d health after heal = %d", healthBeforeHeal, hero.health)
	}
}
