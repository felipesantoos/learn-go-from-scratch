package main

import (
	"fmt"
)

// Function to recover program in case of panic
func recoverProgram() {
	if r := recover(); r != nil {
		fmt.Println("Programa recuperado com sucesso!")
	}
}

// Panic/recover
func givePoints(points int) {
	defer recoverProgram()

	fmt.Println("Analisando pontuação...")
	if points <= 10 && points >= 0 {
		fmt.Println("Nota válida!")
	} else {
		fmt.Println("Entrando em pânico...")
		panic("Nota inválida!")
	}
}

func main() {
	givePoints(1)
	fmt.Println("Fim do programa...")
}

// OK
