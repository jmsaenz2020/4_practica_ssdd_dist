package taller

import(
  "taller/incidencia"
)

type Vehiculo struct{
  Matricula int
  Incidencia taller.Incidencia
}

const MAX_MATRICULA = 100000 - 1
