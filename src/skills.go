package main

import "fmt"

func useSkill(m *Monster) {
	for _, s := range p1.Skills {

		if s == "Boule de Feu" {
			m.HP -= 30
			fmt.Println("🔥 Boule de Feu !")
		}
	}
}
