package main

import (
	"aprende-go/bucles"
	"aprende-go/condicionales"
	"aprende-go/matematicas"
	"aprende-go/variables"
	"fmt"
	"math"
)

func main() {
	fmt.Println(matematicas.Sumar(5, 3))
	fmt.Println(matematicas.Resta(5, 3))
	fmt.Println(math.Sqrt(144))
	variables.MostrarVariable("Nicolas Andres Cea Cerda")
	condicionales.Condicionales(1, 2)
	bucles.Bucles()
	logout := Logout()
	fmt.Println(logout)
}
func Logout() string {
	return "adios"
}
