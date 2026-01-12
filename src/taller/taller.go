package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net"
	"strconv"
	"taller_main/taller"
)

const NUM_PLAZAS = 1
const NUM_VEHICULOS = 2

var (
	buf    bytes.Buffer
	logger = log.New(&buf, "logger: ", log.Lshortfile)
	msg    string
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8000")
	if err != nil {
		logger.Fatal(err)
	}
	defer conn.Close()
	buf := make([]byte, 512)
	var taller taller.Taller
  taller.Inicializar(NUM_PLAZAS, NUM_VEHICULOS)
	for {
		n, err := conn.Read(buf)
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println(err)
			continue
		}
		if n > 0 {
			msg = string(buf[:n])

			estado, err := strconv.Atoi(msg[:len(msg) - 1]) // Quitar saltos de linea y convertir a entero
			if err == nil{
				taller.CambiarEstado(estado)
			} else {
				fmt.Println(err)
			}
			fmt.Println("len: " + strconv.Itoa(n) + " msg: " + msg)
		}
	}
}
