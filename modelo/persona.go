package modelo

import "fmt"

type Persona struct {
	Id     int64
	Nombre string
	Edad   int64
	Estado bool
}

func (p *Persona) imprimirNombre() {
	fmt.Println(p.Nombre)
}