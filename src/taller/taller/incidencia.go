package taller

type Incidencia struct{
  Tipo int
  Fase int
  Estado int // 0 Cerrado, 1 Abierto, 2 En proceso
}

const MAX_TIPO = 3
const MAX_FASE = 4
const MAX_ESTADO_INCIDENCIA = 2
const TIEMPO_MECANICA = 5
const TIEMPO_ELECTRICA = 3
const TIEMPO_CARROCERIA = 1
