package taller

import(
	"fmt"
)

type Taller struct{
	Estado int
	Plazas chan *int
	Exclusividad int
	Prioridad int
}

const NUM_PLAZAS = 10
const MAX_ESTADO = 9
const TALLER_CERRADO = MAX_ESTADO
const MAX_TIPOS = 3

func (t *Taller) Inicializar(){
	t.Plazas = make(chan *int, NUM_PLAZAS)
	fmt.Println("Taller Inicializado")
}

func (t *Taller) Liberar(){
	fmt.Println("Taller Inactivo")
}

func (t *Taller) CambiarExclusividad(exclusividad int){
	if exclusividad >= 1 && exclusividad <= MAX_TIPOS{
		fmt.Println("Exclusivo para prioridad", exclusividad)
		t.Exclusividad = exclusividad
	} else {
		t.Exclusividad = 0
	}
}

func (t *Taller) CambiarPrioridad(prioridad int){
	if prioridad >= 1 && prioridad <= MAX_TIPOS{
		fmt.Println("Mayor prioridad para tipo", prioridad)
		t.Prioridad = prioridad
	} else {
		t.Prioridad = 0
	}
}

func (t Taller) Cerrar(){
	close(t.Plazas)
	fmt.Println("Taller Cerrado")
}

func (t Taller) Actualizar(){
	switch(t.Estado){
		case 0:
			t.Liberar()
		case 1:
			t.CambiarExclusividad(1)
		case 2:
			t.CambiarExclusividad(2)
		case 3:
			t.CambiarExclusividad(3)
		case 4:
			t.CambiarPrioridad(1)
		case 5:
			t.CambiarPrioridad(2)
		case 6:
			t.CambiarPrioridad(3)
		case 9:
			t.Cerrar()
		default:
			fmt.Println("Mantiene estado")
	}
}

func (t *Taller) CambiarEstado(estado int){
	if estado >= 0 && estado <= MAX_ESTADO{
		fmt.Println(estado)
		if estado > 0 && estado <= TALLER_CERRADO{
			t.Estado = estado
			t.Actualizar()
		} else {
			if t.Plazas == nil{
				t.Inicializar()
			} else {
				t.Liberar()
			}
		}
	} else {
		fmt.Println("Estado no valido:", estado)
	}
}