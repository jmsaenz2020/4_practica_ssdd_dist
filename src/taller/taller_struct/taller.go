package taller

import(
	"fmt"
  "sync"
)

type Taller struct{
	Estado int
	Plazas chan int
	Cola chan int
	Exclusividad int
	Prioridad int
  Cerrado bool
  Cierre sync.RWMutex
  Grupo sync.WaitGroup
}

const NUM_PLAZAS = 10
const NUM_MECANICOS = 1
const MAX_ESTADO = 9
const TALLER_CERRADO = MAX_ESTADO
const MAX_TIPOS = 3

func (t *Taller) CrearVehiculo(matricula int){
  t.Cola <- matricula
  fmt.Printf("Coche %05d en la cola\n", matricula)
}

func (t *Taller) Operar(num_mecanicos int, num_vehiculos int, num_plazas int){
  var exit = true  

  for i := 0; i < num_vehiculos; i++{
    go t.CrearVehiculo(i + 1)
  }

  
  for{
    select{
      case vehiculo, ok := <- t.Cola:
        fmt.Println(vehiculo)
        if ok{
          t.Cierre.Lock()
          t.Plazas <- vehiculo
          t.Cierre.Unlock()
        } else {
          exit = true
        }
    }
    if exit{
      break
    }
  }
}

func (t *Taller) Inicializar(num_plazas int){
  var num_mecanicos = 1
  var num_vehiculos = 5

	t.Plazas = make(chan int, num_plazas)
	t.Cola = make(chan int)
	fmt.Println("Taller Inicializado")
  go t.Operar(num_mecanicos, num_vehiculos, num_plazas)
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
			} else {
				t.Liberar()
			}
		}
	} else {
		fmt.Println("Estado no valido:", estado)
	}
}
