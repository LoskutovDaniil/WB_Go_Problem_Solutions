package main

import "fmt"

// Базовая структура
type Human struct {
	Name string
	Age  int
}

// Метод базовой структуры
func (a Human) Defend() {
	fmt.Println("The human defends himself")
}

// Наследуемая структура
type Action struct {
	Human
	Attack string
}

// Метод структуры Action
func (a Action) AttackAction() {
	fmt.Println(a.Attack)
}

func main() {
	humanAction := Action{
		Human: Human{
			Name: "Daniil",
			Age:  25,
		},
		Attack: "Be careful, he attacks!",
	}

	humanAction.Defend()
	humanAction.AttackAction()
}
