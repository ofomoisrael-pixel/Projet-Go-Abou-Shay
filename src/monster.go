package main

import "fmt"

type Monster struct {
	Name   string
	HP     int
	Attack int
}

func createGoblin() Monster {
	return Monster{
		Name:   "Gobelin🧌",
		HP:     40,
		Attack: 5,
	}
}

func isMonsterDead(m *Monster) bool {
	return m.HP <= 0
}

func monsterAttack(m Monster) {
	p1.CurrentHP -= m.Attack
	fmt.Printf("Le %s attaque ! -%d PV\n", m.Name, m.Attack)
	fmt.Printf("PV joueur : %d/%d\n", p1.CurrentHP, p1.MaxHP)
}
