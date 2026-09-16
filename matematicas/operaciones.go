package matematicas

// esta es una funcion  que esta definida con mayusculas por lo tanto se esta exportando con el package

func Sumar(x, y int) int {
	return x + y
}

/*  esta es una funcion  que esta definida con minusculas por lo tanto
esto quieres decir que solo puede ser utilizada dentro del package*/
func sumar(x, y int) int {
	return x + y
}

func Resta(x, y int) int {
	return x - y
}
