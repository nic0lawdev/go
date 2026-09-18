package condicionales

import "fmt"

func Condicionales(a, b int) int {

	//aca declararemos Condicionales
	if a > b {
		fmt.Println("es mayor")
	} else if a == 0 {
		fmt.Print("es cero")
	} else {
		fmt.Println("es menor")

	}
	return a

}
