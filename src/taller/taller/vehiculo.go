package taller

type Vehiculo struct{
  Matricula int
  Incidencia Incidencia
}

const MAX_MATRICULA = 100000 - 1

func (v *Vehiculo) Rutina(t *Taller){
  
  for v.Incidencia.Fase = 1; v.Incidencia.Fase <= MAX_FASE; v.Incidencia.Fase++{
    v.Incidencia.Estado = 1
    t.VehiculoFase(v)
    t.Plazas <- *v
  }
}

func (v Vehiculo) ObtenerTiempo() (int){
  switch v.Incidencia.Tipo{
    case 1:
      return TIEMPO_MECANICA
    case 2:
      return TIEMPO_ELECTRICA
    case 3:
      return TIEMPO_CARROCERIA
    default:
      return 0
  }
}
