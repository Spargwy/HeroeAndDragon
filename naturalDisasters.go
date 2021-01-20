package main

import (
	"fmt"
	"math/rand"
	"time"
)

func naturalDisasters(dragon Dragon, hero Hero) (int, int) {

	rand.Seed(time.Now().UnixNano())
	chance := rand.Intn(100)
	if chance <= 5 {
		fmt.Println("Похолодание сказалось на драконе...\n")
		dragon.health = dragon.health - dragon.health/10*2
	} else if chance > 5 && chance <= 90 {
		hero.health = hero.health - hero.health/10*2
	} else if chance > 10 && chance <= 30 {
		fmt.Println("Дракон назвал рыцаря живодером, из-за чего тот окончательно пригорюнился. Дракон этот шанс, конечно не упустил\n")
		hero.health = hero.health - hero.health/10*7
	} else if chance > 30 && chance <= 50 {
		fmt.Printf("Произошло землятрясение, хвост дракона зажало в щели между двумя стенками разверзнувшейся бездны\n")
		dragon.health = dragon.health - dragon.health/10*9
	} else if chance > 50 && chance < 60 {
		fmt.Printf("Произошло землятрясение, хвост дракона зажало в щели между двумя стенками разверзнувшейся бездны. А герой туда провалился...\n")
		hero.health = hero.health - hero.health/10*9
	}
	return dragon.health, hero.health
}
