package main

import "testing"

func TestEvents(t *testing.T) {
	heroForEvents := Hero{
		maxHealth: 100,
		health:    100,
		damage:    20,
		armor:     47,
	}
	dragonForEvents := Dragon{
		health:     1000,
		damage:     5,
		missChance: 0,
	}

	wantDragonHealth := 6
	dragonForEvents.health = 7
	heroForEvents, dragonForEvents = events(dragonForEvents, heroForEvents, 17)
	if dragonForEvents.health != wantDragonHealth {
		t.Fatalf("ERROR, dragon health must be equal %d, but equal %d", wantDragonHealth, dragonForEvents.health)
	}

	wantDragonHealth = 900
	dragonForEvents.health = 1000
	heroForEvents, dragonForEvents = events(dragonForEvents, heroForEvents, 17)
	if dragonForEvents.health != wantDragonHealth {
		t.Fatalf("ERROR, dragon health must be equal %d, but equal %d", wantDragonHealth, dragonForEvents.health)
	}

	dragonForEvents.health = 4
	wantDragonHealth = 2
	heroForEvents, dragonForEvents = events(dragonForEvents, heroForEvents, 5)
	if dragonForEvents.health != wantDragonHealth {
		t.Fatalf("ERROR, dragon health must be equal %d, but equal %d", wantDragonHealth, dragonForEvents.health)
	}

	dragonForEvents.health = 1000
	wantDragonHealth = 800
	heroForEvents, dragonForEvents = events(dragonForEvents, heroForEvents, 5)
	if dragonForEvents.health != wantDragonHealth {
		t.Fatalf("ERROR, dragon health must be equal %d, but equal %d", wantDragonHealth, dragonForEvents.health)
	}

	dragonForEvents.health = 8
	wantDragonHealth = 4
	heroForEvents, dragonForEvents = events(dragonForEvents, heroForEvents, 12)
	if dragonForEvents.health != wantDragonHealth {
		t.Fatalf("ERROR, dragon health must be equal %d, but equal %d", wantDragonHealth, dragonForEvents.health)
	}

	dragonForEvents.health = 1000
	wantDragonHealth = 500
	heroForEvents, dragonForEvents = events(dragonForEvents, heroForEvents, 12)
	if dragonForEvents.health != wantDragonHealth {
		t.Fatalf("ERROR, dragon health must be equal %d, but equal %d", wantDragonHealth, dragonForEvents.health)
	}

	heroForEvents.damage = 8
	wantHeroDamage := 11
	heroForEvents, dragonForEvents = events(dragonForEvents, heroForEvents, 25)
	if heroForEvents.damage != wantHeroDamage {
		t.Fatalf("ERROR, hero damage must be equal %d, but equal %d", wantHeroDamage, heroForEvents.damage)
	}

	wantHeroDamage = 26
	heroForEvents.damage = 20
	heroForEvents, dragonForEvents = events(dragonForEvents, heroForEvents, 25)
	if heroForEvents.damage != wantHeroDamage {
		t.Fatalf("ERROR, hero damage must be equal %d, but equal %d", wantHeroDamage, heroForEvents.damage)
	}

	heroForEvents.health = 8
	wantHeroHealth := 6
	heroForEvents, dragonForEvents = events(dragonForEvents, heroForEvents, 7)
	if heroForEvents.health != wantHeroHealth {
		t.Fatalf("ERROR, hero health must be equal %d, but equal %d", wantHeroHealth, heroForEvents.health)
	}

	wantHeroHealth = 64
	heroForEvents.health = 80
	heroForEvents, dragonForEvents = events(dragonForEvents, heroForEvents, 7)
	if heroForEvents.health != wantHeroHealth {
		t.Fatalf("ERROR, hero health must be equal %d, but equal %d", wantHeroHealth, heroForEvents.health)
	}

}
