package main

import (
	"fmt"

	"github.com/mauriGaray/GoPrueba.git/modelo"
)

func main (){
	p := modelo.Persona{
		Id: 1,
		Nombre: "Carlos",
		Edad: 36,
		Estado: true,
	}
	fmt.Println(p.imprimirNombre())
}