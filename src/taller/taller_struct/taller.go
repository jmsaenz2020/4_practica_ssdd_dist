package taller

import(
	"fmt"
  "sync"
  "taller/vehiculo"
)

type Taller struct{
	Estado int
	Plazas chan taller.Vehiculo
	Cola chan taller.Vehiculo
	Exclusividad int
	Prioridad int
  Cerrado bool
  Cierre sync.RWMutex
  Grupo sync.WaitGroup
}

const NUM_PLAZAS = 5
const NUM_MECANICOS = 1
const MAX_ESTADO = 9
const TALLER_CERRADO = MAX_ESTADO
const MAX_TIPOS = 3

func (t *Taller) Operar(){

  for vehiculo := range t.Cola{
    t.Cierre.Lock()
    t.Plazas <- vehiculo
    fmt.Printf("Coche %05d en una plaza\n", vehiculo.Matricula)
    t.Cierre.Unlock()
  }
}

func (t *Taller) GenerarVehiculos(num_vehiculos int){
  var v taller.Vehiculo

  for i := 0; i < num_vehiculos; i++{
    v.Matricula = i + 1
    v.Incidencia.Tipo = 1
    //if v.Valido(){
      t.Cola <- v
      fmt.Printf("Coche %05d en la cola\n", v.Matricula)
    //}
  }
  close(t.Cola)
}

func (t *Taller) Inicializar(num_plazas int){
  //var num_mecanicos = 1
  var num_vehiculos = 10

	t.Plazas = make(chan taller.Vehiculo, num_plazas)
	t.Cola = make(chan taller.Vehiculo)
	fmt.Println("Taller Inicializado")

  go t.GenerarVehiculos(num_vehiculos)
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
		fmt.Println("Prioridad elevada para tipo", prioridad)
		t.Prioridad = prioridad
	} else {
		t.Prioridad = 0
	}
}

func (t Taller) Cerrar(){
  t.Cerrado = true
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
				t.Inicializar(NUM_PLAZAS)
        go t.Operar()
			} else {
				t.Liberar()
			}
		}
	} else {
		fmt.Println("Estado no valido:", estado)
	}
}
