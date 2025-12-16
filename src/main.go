package main

import (
  "4_practica_ssdd_dist/taller"
  "math/rand"
)

const sims = 1
const NUM_COCHES = 1
const NUM_MECANICOS = 1

func valoresBasicos() (taller.Taller){
  var t taller.Taller

  t.Inicializar()
  for n := 0; n < NUM_MECANICOS; n++{
    t.CrearMecanico("Mecanico", 0, 0)
  }

  for n := 0; n < NUM_COCHES; n++{
    t.CrearVehiculo(rand.Int()%taller.MAX_MATRICULA + 1, rand.Int()%3 + 1, "Incidencia")
  }

  return t
}

func main(){
  var t taller.Taller
  
  for i := 0; i < sims; i++{
    t = valoresBasicos()
    t.Liberar()
  }
}
