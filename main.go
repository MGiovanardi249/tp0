package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"tp0/ejercicios"
)

func lectura(Archivo string) []int {
	archivo, _ := os.Open(Archivo)
	defer archivo.Close()

	var numeros []int

	scanner := bufio.NewScanner(archivo)

	for scanner.Scan() {
		texto := scanner.Text()
		numero, _ := strconv.Atoi(texto)

		numeros = append(numeros, numero)

	}
	return numeros
}

func main() {
	v1 := lectura("archivo1.in")
	v2 := lectura("archivo2.in")

	if ejercicios.Comparar(v1, v2) == -1 {
		ejercicios.Seleccion(v2)
		for _, v := range v2 {
			fmt.Println(v)
		}
	} else {
		ejercicios.Seleccion(v1)
		for _, v := range v1 {
			fmt.Println(v)
		}
	}

}
