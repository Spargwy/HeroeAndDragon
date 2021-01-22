package main

import (
	"fmt"
	"testing"
)

func TestHeroTiredness(t *testing.T) {
	hero := Hero{
		maxHealth: 100,
		health:    100,
		damage:    20,
	}

	wantDamageReduction := 0
	DamageReduction := heroTiredness(hero)
	fmt.Print("DR: ", DamageReduction)
	if wantDamageReduction != DamageReduction {
		t.Fatalf("ERROR!!! Damage reduction must be equal %d, but euqal %d", wantDamageReduction, DamageReduction)
	}
	hero.health -= 70
	wantDamageReduction = 5
	DamageReduction = heroTiredness(hero)
	fmt.Print("DR: ", DamageReduction)
	if wantDamageReduction != DamageReduction {
		t.Fatalf("ERROR!!! Damage reduction must be equal %d, but euqal %d", wantDamageReduction, DamageReduction)
	}
	hero.health -= 10
	wantDamageReduction = 10
	DamageReduction = heroTiredness(hero)
	fmt.Print("DR: ", DamageReduction)
	if wantDamageReduction != DamageReduction {
		t.Fatalf("ERROR!!! Damage reduction must be equal %d, but euqal %d", wantDamageReduction, DamageReduction)
	}
	hero.health -= 10
	wantDamageReduction = 20
	DamageReduction = heroTiredness(hero)
	fmt.Print("DR: ", DamageReduction)
	if wantDamageReduction != DamageReduction {
		t.Fatalf("ERROR!!! Damage reduction must be equal %d, but euqal %d", wantDamageReduction, DamageReduction)

	}
}
