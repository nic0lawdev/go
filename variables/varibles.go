package variables

import "fmt"

//aca estamos declarando una varible
var nombre = "Nicolas Cea"

// aca estamos declarando varias varibles al mismo tiempo
var a, b, c = 3, 3, 4

// en go no es necesario declarar el tipo de dato que se guardara en al varible
// para definir una variable de esta manera es necesario hacerlo dentro de una funncion

// nombre  := "Nicolas Cea"

func MostrarVariable(nombre string) {
	// esta es la manera mas facil/comun de declarar una variable en go
	Nombre := nombre
	fmt.Println("mi nombre es: " + Nombre)
}

//en go tambien podremos declarar contantes con la palabra clave CONST

const Pi = 3.1415

// tambien las constantes se pueden definir en un grupo de constantes

const (
	g = 1
	h = 2
	j = 3
)
