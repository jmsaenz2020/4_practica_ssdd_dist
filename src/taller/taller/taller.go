package taller

import(
	"fmt"
  "sync"
  "time"
)

type Taller struct{
	Estado int
	Plazas chan Vehiculo
	Cola chan Vehiculo
	Exclusividad int
	Prioridad int
  Cerrado bool
  Cierre sync.RWMutex
  Grupo sync.WaitGroup
  Tiempo time.Time
}

const NUM_MECANICOS = 1
const MAX_ESTADO_TALLER = 9
const TALLER_CERRADO = MAX_ESTADO_TALLER
const MAX_TIPOS = 3

func (t *Taller) Operar(){

  for vehiculo := range t.Cola{
    go vehiculo.Rutina(t)
  }
}

func (t Taller) InfoMsg(v Vehiculo){
  tiempo := time.Now().Sub(t.Tiempo)

  fmt.Printf("Tiempo %s Coche %05d Incidencia %d Fase %d Estado %d\n", tiempo, v.Matricula, v.Incidencia.Tipo, v.Incidencia.Fase, v.Incidencia.Estado)
}

func (t *Taller) VehiculoFase(v *Vehiculo){
  v.Incidencia.Estado = 2
  t.InfoMsg(*v)
  time.Sleep(time.Duration(v.ObtenerTiempo())*time.Second)

  v.Incidencia.Estado = 1
  switch v.Incidencia.Fase{
    case 1:
      t.Plazas <- *v
    case 4:
      v.Incidencia.Estado = 0
  }

  t.InfoMsg(*v)
}

func (t *Taller) GenerarVehiculos(num_vehiculos int){
  var v Vehiculo

  for i := 0; i < num_vehiculos; i++{
    v.Matricula = i + 1
    v.Incidencia.Tipo = 1
    //if v.Valido(){
      t.Cola <- v
    //}
  }
  close(t.Cola)
}

func (t *Taller) Inicializar(num_plazas int){
  //var num_mecanicos = 1
  var num_vehiculos = 10

	t.Plazas = make(chan Vehiculo, num_plazas)
	t.Cola = make(chan Vehiculo)
  t.Tiempo = time.Now()

  go t.GenerarVehiculos(num_vehiculos)
  go t.Operar()
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
	if estado >= 0 && estado <= MAX_ESTADO_TALLER{
		fmt.Println(estado)
		if estado > 0 && estado <= TALLER_CERRADO{
			t.Estado = estado
			t.Actualizar()
		} else {
      t.Liberar()
		}
	} else {
		fmt.Println("Estado no valido:", estado)
	}
}
