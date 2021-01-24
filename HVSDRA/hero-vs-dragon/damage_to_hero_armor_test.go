package main

import (
	"testing"
)

func TestDamageToArmor(t *testing.T) {
	hero := Hero{
		maxHealth: 100,
		health:    40,
		damage:    20,
		armor:     47,
	}
	hero2 := Hero{
		maxHealth: 100,
		health:    40,
		damage:    20,
		armor:     2,
	}

	hero3 := Hero{
		maxHealth: 100,
		health:    40,
		damage:    20,
		armor:     0,
	}
	dragon := Dragon{
		health:     1000,
		damage:     5,
		missChance: 0,
	}

	hero.armor, hero.health = damageToArmor(hero, dragon)
	wantHeroHealth := 40
	wantHeroArmor := 42
	if hero.armor != wantHeroArmor {
		t.Fatalf("ERROR!!! hero armor must be equal %d, but equal %d", wantHeroArmor, hero.armor)
	}
	if hero.health != wantHeroHealth {
		t.Fatalf("ERROR!!! hero health must be equal %d, but equal %d", wantHeroHealth, hero.health)
	}

	hero2.armor, hero2.health = damageToArmor(hero2, dragon)
	wantHeroHealth = 37
	wantHeroArmor = 0
	if hero2.armor != wantHeroArmor {
		t.Fatalf("ERROR!!! hero armor must be equal %d, but equal %d", wantHeroArmor, hero2.armor)
	}
	if hero2.health != wantHeroHealth {
		t.Fatalf("ERROR!!! hero health must be equal %d, but equal %d", wantHeroHealth, hero2.health)
	}

	hero3.armor, hero3.health = damageToArmor(hero3, dragon)
	wantHeroHealth = 35
	wantHeroArmor = 0
	if hero3.armor != wantHeroArmor {
		t.Fatalf("ERROR!!! hero armor must be equal %d, but equal %d", wantHeroArmor, hero3.armor)
	}
	if hero3.health != wantHeroHealth {
		t.Fatalf("ERROR!!! hero health must be equal %d, but equal %d", wantHeroHealth, hero3.health)
	}

}
