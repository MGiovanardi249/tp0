package main

import (
	"tp0/ejercicios"
)


func main() {
	v1 := ejercicios.Lectura("archivo1.in")
	v2 := ejercicios.Lectura("archivo2.in")

	var ganador []int

	if ejercicios.Comparar(v1,v2) == -1 {
		ganador = v2
	}else{
		ganador = v1
	}
	ejercicios.Seleccion(ganador)
	ejercicios.Imprimir(ganador)
	}



