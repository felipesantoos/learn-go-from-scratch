package main

import (
	"fmt"
)

func main() {
	// Função anônima simples
	func() {
		fmt.Println("Função anônima 1")
	}()

	// Função anônima com parâmetro
	func(param string) {
		fmt.Println(param)
	}("Função anônima 2")

	// Função anônima com parâmetro e retorno
	ret := func(param interface{}) string {
		return fmt.Sprintf("Função anônima %v", param)
	}(3)
	fmt.Println(ret)
}

// OK
