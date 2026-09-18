package bucles

import "fmt"

func Bucles() {
	// aca lo que se esta haciendo es recorrer i hasta que sea menor que 6
	for i := 0; i < 6; i++ {
		fmt.Println(i)
	}
	// aca estamos definiendo una lista de datos
	nums := []int{1, 2, 3}

	//aca estamos recorriendo la lista de datos definida anteriormente
	for _, num := range nums {
		fmt.Println(num)
	}

}
