package main

import (
	"fmt"
	"database/sql"
	"log"
	"time"

	adminBase "github.com/fernandez-florio-gonzalez-torrico-tp/adminBase"
	creacionBase "github.com/fernandez-florio-gonzalez-torrico-tp/creacionBase"
	crearPkFK "github.com/fernandez-florio-gonzalez-torrico-tp/crearPkFK"
	_ "github.com/lib/pq"
)
func llamarTelefonicaFraude(){
	for{
		
		fmt.Printf("%v: telefonica esta consultando el fraude \n", time.Now().Format("15:04:05"));
		time.Sleep(3*time.Second);		

func llamarTelefonicaFraude() {
	for {

		fmt.Printf("%v: telefonica esta consultando el fraude \n", time.Now().Format("15:04:05"))
		time.Sleep(3 * time.Second)
	}

}
type cliente struct{
	nrocliente int
	nombre, apellido, domicilio, telefono string
}
type tarjeta struct{
	nrotarjeta, validadesde, validahasta, codseguridad, estado string
	nrocliente int
	limitecompra float64
}
type comercio struct{
	nombre, domicilio, codigopostal, telefono string
	nrocomercio int
}
/*type compra struct{
	Offset int
	nrooperacion, nrocomercio int
	nrotarjeta, string string
	monto float64
	pagado bool
}*/
func verCliente(){
	db, err := sql.Open("postgres" , "user=postgres host=localhost dbname=tarjetascredito sslmode=disable")
	if err != nil{
		log.Fatal(err)
	}
	defer db.Close()
	fmt.Printf("Clientes:\n")
	rows, err := db.Query (`select * from cliente`)
	if err != nil{
		log.Fatal(err)
	}
	defer rows.Close()
	var c cliente
	for rows.Next(){
		if err := rows.Scan(&c.nrocliente, &c.nombre, &c.apellido, &c.domicilio, &c.telefono); err != nil{
			log.Fatal(err)
		}
		fmt.Printf("%v %v %v %v %v\n", c.nrocliente, c.nombre, c.apellido, c.domicilio, c.telefono )
	}
	if err = rows.Err(); err!= nil{
		log.Fatal(err)
	}
}
func verComercios(){
	db, err := sql.Open("postgres" , "user=postgres host=localhost dbname=tarjetascredito sslmode=disable")
	if err != nil{
		log.Fatal(err)
	}
	defer db.Close()
	fmt.Printf("Comercios:\n")

	rows, err := db.Query (`select * from Comercio`)
	if err != nil{
		log.Fatal(err)
	}
	defer rows.Close()
	var c comercio
	for rows.Next(){
		if err := rows.Scan(&c.nrocomercio, &c.nombre, &c.domicilio, &c.codigopostal, &c.telefono); err != nil{		log.Fatal(err)
		}
		fmt.Printf("%v %v %v %v %v \n", c.nrocomercio, c.nombre, c.domicilio, c.codigopostal, c.telefono)
	}
	if err = rows.Err(); err!= nil{
		log.Fatal(err)
	}
}
func verTarjetas(){
	db, err := sql.Open("postgres" , "user=postgres host=localhost dbname=tarjetascredito sslmode=disable")
	if err != nil{
		log.Fatal(err)
	}
	defer db.Close()
	fmt.Printf("Tarjetas:\n")

	rows, err := db.Query (`select * from tarjeta`)
	if err != nil{
		log.Fatal(err)
	}
	defer rows.Close()
	var t tarjeta
	for rows.Next(){
		if err := rows.Scan(&t.nrotarjeta, &t.nrocliente, &t.validadesde, &t.validahasta, &t.codseguridad, &t.limitecompra, &t.estado); err != nil{
			log.Fatal(err)
		}
		fmt.Printf("%v %v %v %v %v\n", t.nrotarjeta, t.nrocliente, t.validadesde, t.validahasta, t.codseguridad, t.limitecompra, t.estado )
	}
	if err = rows.Err(); err!= nil{
		log.Fatal(err)
	}
}

func baseJsonNoSql() {
	verTarjetas()
	verComercios()
	verCliente()
}
/*func verCompras(){
	db, err := sql.Open("postgres" , "user=postgres host=localhost dbname=tarjetascredito sslmode=disable")
	if err != nil{
		log.Fatal(err)
	}
	defer db.Close()
	fmt.Printf("Compras:\n")

	rows, err := db.Query (`select * from compra`)
	if err != nil{
		log.Fatal(err)
	}
	defer rows.Close()
	var c compra
	for rows.Next(){
		if err := rows.Scan(&c.nrooperacion, &c.nrotarjeta, &c.nrocomercio, &c.fecha, &c.monto, &c.pagado); err != nil{
			log.Fatal(err)
		}
		fmt.Printf("%v %v %v %v %v %v\n", c.nrooperacion, c.nrotarjeta, c.nrocomercio, c.fecha, c.monto, c.pagado )
	}
	if err = rows.Err(); err!= nil{
		log.Fatal(err)
	}
}*/
func main() {

	var input int;
	var seguir string;
	creacionBase.CrearBase()
	creacionBase.CrearTablas()
	creacionBase.InsertarDatos()
	crearPkFK.CrearPKFK()
	adminBase.ValidacionesFuncionesDeCompra()
	adminBase.DarResumenFunciones()
	for {
		//imprimo menu
		fmt.Printf("Seleccione: " +
			"\n 1 - ver clientes" +
			"\n 2 - ver tarjetas" +
			"\n 3 - ver comercios" +
			"\n 4 - crear PK y FK" +
			"\n 5 - drop PK y FK" +
			"\n 6 - crear funciones de autorizacion" +
			"\n 7 - crear alertas de compra" +
			"\n 8 - telefonica realiza consultas" +
			"\n 9 - exit\n")
		fmt.Printf("ingrese numero: ")
		fmt.Scanf("%d", &input)
		fmt.Printf("selecciono: %d \n", input)
		switch input {
		case 1:
		 	verCliente();
		case 2:
			verTarjetas();
		case 3:
			verComercios();
		case 4:
			fmt.Printf("Creamos las PK y FK \n")
			crearPkFK.CrearPKFK();
		case 5:
			fmt.Printf("Drop de PK y FK \n")
			crearPkFK.DropLlavesFKPK();
		case 6:
			fmt.Printf("Instalar funciones de validacion de compra \n")
			adminBase.ValidacionesFuncionesDeCompra();
		case 7:
			fmt.Printf("Activar alertas de compra\n")
			adminBase.ActivarAlertas();
		case 8:
			fmt.Printf("consultando \n")
			go llamarTelefonicaFraude()
			time.Sleep(10 * time.Second)
		case 9:
			
			break
		default:
			fmt.Printf("opcion no valida\n")
		}
		fmt.Printf("\n desea seguir? s/n: \n")
		fmt.Scanf("%s", &seguir)
		if seguir == "n" {
			break
		}
	}

}
