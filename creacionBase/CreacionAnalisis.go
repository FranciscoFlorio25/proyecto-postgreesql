package main
import (
	"encoding/json"
	"fmt"
	bolt "go.etcd.io/bbolt"
	"database/sql"1
	_ "github.com/lib/pq"
)

type cliente struct {
	nrocliente int
	nombre string
	apellido string
	domicilio string
	telefono string
}

type tarjeta struct {
	nrotarjeta string
	validadesde string
	validahasta string
	
}

func s
