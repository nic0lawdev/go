package main

import (
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

}
