package main

import (
	"fmt"
	"math/rand"
)

type Monster struct {
	Name   string
	HP     int
	Attack int
}

func createMonster() Monster {
	if rand.Intn(2) == 0 {
		return Monster{"Gobelin", 40, 5}
	}
	return Monster{"Orc", 80, 10}
}

func isMonsterDead(m *Monster) bool {
	return m.HP <= 0
}

func monsterAttack(m Monster) {
	p1.CurrentHP -= m.Attack
	fmt.Printf("Le %s attaque ! -%d PV\n", m.Name, m.Attack)
	fmt.Printf("PV joueur : %d/%d\n", p1.CurrentHP, p1.MaxHP)
}
func createBoss() Monster {
	return Monster{
		Name:   "Dragon🐉",
		HP:     300,
		Attack: 40,
	}
}
