package taller

import (
  "fmt"
  "sync"
  "time"
  "4_practica_ssdd_dist/utils"
)

const NUM_MECANICOS = 1
const NUM_PLAZAS = 1

type Taller struct{
  Plazas chan *Vehiculo
  NumPlazas int
  NumMecanicos int
  Mecanicos []Mecanico
  UltimoIdMecanico int
  UltimoIdIncidencia int
  Grupo sync.WaitGroup
  Cerradura sync.RWMutex
  TiempoInicio time.Time
}

func (t *Taller)Inicializar(){
  t.Plazas = make(chan *Vehiculo, NUM_PLAZAS)
  t.TiempoInicio = time.Now()
}

func (t *Taller)Liberar(){
  t.Grupo.Wait()
  select{
    case _, ok := <- t.Plazas:
      if !ok && len(t.Plazas) == 0{
        close(t.Plazas)
      }
    default:
      close(t.Plazas)
  }
}

func (t Taller) HayEspacio() (bool){
  vehiculos := t.ObtenerPlazas()

  return len(vehiculos) <= NUM_PLAZAS
}

func (t *Taller) AsignarPlaza(v *Vehiculo) (bool){
  if t.HayEspacio() && v.Valido() && v.Incidencia.Valido(){
    t.EntrarVehiculo(v)
    return true
  } else if !v.Incidencia.Valido(){
    utils.WarningMsg("El vehiculo no tiene una incidencia definida")
  }
  return false
}

func (t *Taller) EntrarVehiculo(v *Vehiculo){
  t.Cerradura.Lock()
  t.Plazas <- v
  t.Cerradura.Unlock()
}

func (t *Taller) SalirVehiculo(v *Vehiculo){
  t.Cerradura.Lock()
  <- t.Plazas
  t.Cerradura.Unlock()
}

func (t *Taller) AsignarMecanicoAutomatico(v *Vehiculo){
  if t.HayEspacio(){
    for _, m := range(t.ObtenerMecanicosDisponibles()){
      if m.Alta{
        v.Incidencia.AsignarMecanico(m)
        break
      }
    }
  }
}

func (t *Taller) CrearMecanico(nombre string, especialidad int, experiencia int){
  var m Mecanico

  m.Nombre = nombre
  m.Especialidad = especialidad
  m.Experiencia = experiencia
  m.Id = t.UltimoIdMecanico + 1

  if m.Valido(){
    t.UltimoIdMecanico++
    t.Mecanicos = append(t.Mecanicos, m)
  } else{
    utils.ErrorMsg("Mecanico no se ha podido crear")
  }
}

func (t *Taller) CrearVehiculo(matricula int, tipo int, descripcion string){
  var v Vehiculo
  
  v.Matricula = matricula
  v.Marca = "Marca"
  v.Modelo = "Modelo"
  v.ModificarIncidencia(tipo, descripcion)

  if v.Valido(){
    t.Grupo.Add(1)
    go v.Rutina(t)
  } else {
    utils.ErrorMsg("No se ha creado el vehiculo")
    if !v.Incidencia.Valido(){
      utils.ErrorMsg("La incidencia del vehiculo no es válida")
    }
  }
}

func (t Taller) ObtenerIndiceMecanico(m_in Mecanico) (int){
  var res int = -1

  for i, m := range t.Mecanicos{
    if m.Igual(m_in){
      res = i
    }
  }

  return res
}

func (t Taller) ObtenerMecanicoPorId(id int) (Mecanico){
  var res Mecanico

  for i, m := range t.Mecanicos{
    if m.Id == id{
      res = t.Mecanicos[i]
    }
  }

  return res
}

func (t Taller) ObtenerMecanicosDisponibles() ([]Mecanico){
  var mecanicos []Mecanico  

  for _, m := range t.Mecanicos{
    if m.Alta{
      mecanicos = append(mecanicos, m)
    }
  }

  return mecanicos
}

func (t Taller) ObtenerPlazas() ([]Vehiculo){
  var vehiculos []Vehiculo
  var v Vehiculo
  var exit bool = false

  t.Cerradura.Lock()
  for{
    select{
      case p := <- t.Plazas:
        if p != nil{
          v = *p
          if v.Valido(){
            vehiculos = append(vehiculos, v)
          }
        }
        t.Plazas <- p
      default:
        exit = true
    }
    if exit{
      t.Cerradura.Unlock()
      break
    }
  }

  return vehiculos
}

func (t Taller) ObtenerIncidenciasMecanico(m_in Mecanico) ([]Incidencia){
  var incidencias []Incidencia

  for _, inc := range incidencias{
    if inc.TieneMecanico(m_in){
      incidencias = append(incidencias, inc)
    }
  }

  return incidencias
}

func (t Taller) MecanicosDisponibles(){
  for _, m := range t.Mecanicos{
    if m.Alta{
      fmt.Println(m.Info())
    }
  }
}

func (t *Taller) ModificarMecanico(modif Mecanico){
  for i, m := range t.Mecanicos{
    if m.Igual(modif){
      t.Mecanicos[i] = modif
    }
  }
}
