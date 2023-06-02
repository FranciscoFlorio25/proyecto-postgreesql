package creacionBase

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func insertarComercios() {
	db, err := sql.Open("postgres", "user=postgres host=localhost dbname=tarjetascredito sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	fmt.Print("insertando tarjetas ... \n")

	_, err = db.Exec(`
		insert into comercio values (1, 'Fashion House', 'Av. Presidente Perón 1622/1624/1626, San Miguel, Provincia de Buenos Aires', 'B1662ASX', '541145239875');
		insert into comercio values (2, 'Acr Nomads', 'Calle Constituyentes Ramal Pilar Local G08, Tortuguitas, Provincia de Buenos Aires', 'B1667FYP', '541190874627');
		insert into comercio values (3, 'Feedbox', 'Granaderos A Caballo 4756, José C. Paz, Provincia de Buenos Aires', 'B1665GNZ', '541106376281');
		insert into comercio values (4, 'Rockstar', 'Ruta 202 Y, Aparcamiento, RP8, Provincia de Buenos Aires', 'B1611ABC', '541176310986');
		insert into comercio values (5, 'Senses 4 Fitness', 'Av. Ángel T. de Alvear 2956, CVJ, Provincia de Buenos Aires','B1611ABO', '541162537860');
		insert into comercio values (6, 'Elitesur', 'Arturo Jauretche 978, Hurlingham, Provincia de Buenos Aires', 'B1686FCB', '541104902783');
		insert into comercio values (7, 'Toys R Us', 'Av. Bartolomé Mitre 2857, Moreno, Provincia de Buenos Aires', 'B1744OHM', '541189023287');
		insert into comercio values (8, 'Gamestop', 'Cam. Bancalari, San Fernando, Provincia de Buenos Aires', 'B1646EBS', '541183740387');
		insert into comercio values (9, 'Gamecentric', 'Av. Brig. Gral. Juan Manuel de Rosas 658, Castelar, Provincia de Buenos Aires', 'B1712HIS', '541174382902');
		insert into comercio values (10, 'Veal Furniture', 'Martín Coronado 613, Luis Guillon, Provincia de Buenos Aires', 'B1838EPM', '541130123134');
		insert into comercio values (11, 'Bienestar Cabanas', 'Los Polvorines, Provincia de Buenos Aires', 'B1613HSR', '541110283107');
		insert into comercio values (12, 'A Cats Paw', 'Gutenberg 1829, San Miguel, Provincia de Buenos Aires', 'B1663KPM', '541191759055');
		insert into comercio values (13, 'Papas Cocina', 'Av. del Libertador 7520, Trujui, Provincia de Buenos Aires', 'B1744AEW', '541177390228');
		insert into comercio values (14, 'Yakuza 6', 'Av. Eva Duarte de Perón 1285, Grand Bourg, Provincia de Buenos Aires', 'B1615KUT', '541162556288');
		insert into comercio values (15, 'KyoAi Comics', 'Galería Altube, José C. Paz 1736 Local 132, Gran Buenos Aires, Provincia de Buenos Aires', 'B1665BCF', '541174662411');
		insert into comercio values (16, 'ExileRPG', 'Valparaiso & Mario Bravo, Pablo Nogués, Provincia de Buenos Aires', 'B1616GEK', '541152637188');
		insert into comercio values (17, 'Carroussel', 'Av. Cnel. Niceto Vega 5601, CABA', 'C1414BFE', '541182771637');
		insert into comercio values (18, 'El Valle de Juguetes', 'Av. San Martín 1989, Bella Vista, Provincia de Buenos Aires', 'B1661HVF', '541111263749');
		insert into comercio values (19, 'Peluquerìa Rulito Rebelde', 'Av. Ing. Eduardo Madero 1380, Del Viso, Provincia de Buenos Aires', 'B1669CKU', '541172162516');
		insert into comercio values (20, 'TheWalkingUrban', 'Fray L. Beltrán 2266, Paso del Rey, Provincia de Buenos Aires', 'B1742CBL', '541162621890');
	`)
}

func insertarTarjetas() {
	db, err := sql.Open("postgres", "user=postgres host=localhost dbname=tarjetascredito sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	_, err = db.Exec(`
		insert into tarjeta values ('5952147920589745',1,'201701','202201','7451',92.500000,'vigente');
		insert into tarjeta values ('1235793458027931',1,'202002','202202','9311',92.500000,'suspendida');
		insert into tarjeta values ('7892154698012647',2,'201001','202201','6471',92.500000,'vigente');
		insert into tarjeta values ('6079255489332479',2,'202102','202204','4791',92.500000,'suspendida');
		insert into tarjeta values ('7831554892159706',3,'202201','202201','7061',9.500000,'vigente');
		insert into tarjeta values ('4892008799515300',4,'201005','202205','3001',92.500000,'vigente');
		insert into tarjeta values ('0254795567920014',5,'201506','202206','0141',92.500000,'vigente');
		insert into tarjeta values ('0594122374620017',6,'201207','202207','0191',92.500000,'vigente');
		insert into tarjeta values ('9513570297121457',7,'201108','202208','4571',92.500000,'vigente');
		insert into tarjeta values ('5007877412100889',8,'201409','202209','8891',92.500000,'vigente');
		insert into tarjeta values ('1120489766122002',9,'201110','202210','0221',92.500000,'vigente');
		insert into tarjeta values ('6604789521323411',10,'201711','202211','4111',92.500000,'vigente');
		insert into tarjeta values ('2204879540879044',11,'201912','202212','0441',92.500000,'vigente');
		insert into tarjeta values ('3355541008795201',12,'201002','202202','2011',92.500000,'vigente');
		insert into tarjeta values ('9870965418097626',13,'201602','202202','6261',92.500000,'vigente');
		insert into tarjeta values ('8800484048555666',14,'201405','202205','6661',92.500000,'vigente');
		insert into tarjeta values ('4561234877951000',15,'201806','202206','0001',92.500000,'vigente');
		insert into tarjeta values ('1234879555720448',16,'201402','202207','4481',92.500000,'vigente');
		insert into tarjeta values ('4894519608409840',17,'201003','202212','8401',92.500000,'vigente');
		insert into tarjeta values ('1115557779204122',18,'201401','202209','1221',92.500000,'vigente');
		insert into tarjeta values ('9998804540775553',19,'201309','202209','5531',92.500000,'vigente');
		insert into tarjeta values ('8879208794544830',20,'201502','202201','8301',92.500000,'vigente');
	`)

	fmt.Printf("Tarjetas insertadas \n")
}
func insertarClientes() {
	db, err := sql.Open("postgres", "user=postgres host=localhost dbname=tarjetascredito sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	fmt.Print("insertando clientes ... \n")
	_, err = db.Exec(`
		insert into cliente values (1,'Juan','Carlos','Gelly y obes 320','11504890');
		insert into cliente values (2,'Javier','Marenco','Pte.Peron 520','11514850');
		insert into cliente values (3,'Jose','Carrasco','Pte.Peron 430','11902800');
		insert into cliente values (4,'Jasinto','Gonzales','San martin 300','11204891');
		insert into cliente values (5,'Jaime','Moron','Gelly y obes 900','11401890');
		insert into cliente values (6,'Gerundio','Oracion','Cordoba 604','11103010');
		insert into cliente values (7,'German','Caballieri','Muñoz 207','11302900');
		insert into cliente values (8,'Gonzalo','Gonzales','Moine 705','11910820');
		insert into cliente values (9,'Juana','Gonzales','Moine 705','11712890');
		insert into cliente values (10,'Juliana','Alberti','Alberti 100','11994890');
		insert into cliente values (11,'Julian','Toloza','Sgt.cabral 506','11505891');
		insert into cliente values (12,'Inigo','Montoya','Paunero 2000','11519090');
		insert into cliente values (13,'Ian','Belasques','Paunero 2000','11314590');
		insert into cliente values (14,'Hernan','Rondelli','Pte.Peron 626','11668190');
		insert into cliente values (15,'Hernan','Czemerinski','Alberti 500','11561290');
		insert into cliente values (16,'Maximo','Vida','Uruguay 4400','11504890');
		insert into cliente values (17,'Leon','Rey','Ituizango 1150','11804840');
		insert into cliente values (18,'Leonardo','Hofstadter','Uruguay 2050','11115990');
		insert into cliente values (19,'Sheldon','Cooper','Roca 5550','11811820');
		insert into cliente values (20,'Juan','Nieve','Roca 999','11705990');
		`)
	fmt.Print("tabla cliente insertada \n")
}

func insertarConsumo() {
	db, err := sql.Open("postgres", "user=postgres host=localhost dbname=tarjetascredito sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	fmt.Print("insertando Consumos ... \n")
	_, err = db.Exec(`
		insert into consumo(nrotarjeta,codseguridad,nrocomercio,monto)
		values
		('5952147920589745','7451',1,92.500000),
		('6079255489332479','4791',1,92.500000),
		('7831554892159706','7061',1,92.500000),
		('4892008799515300','1111',1,92.500000),
		('1111111111111111','6471',1,92.500000),
		('8879208794544830','8301',1,92.500000),
		('1235793458027931','9311',1,92.500000);
		`)
	fmt.Print("tabla consumo insertada \n")
}

func insertarCierres() {
	db, err := sql.Open("postgres", "user=postgres host=localhost dbname=tarjetascredito sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	fmt.Print("insertando cierres ... \n")
	_, err = db.Exec(`
			insert into cierre(año, mes, terminacion, fechainicio, fechacierre, fechavto)
			values
			(2022,01,0,'2022-01-01','2022-01-31','2022-02-05'),
			(2022,02,0,'2022-02-01','2022-02-28','2022-03-02'),
			(2022,03,0,'2022-03-01','2022-03-31','2022-04-03'),
			(2022,04,0,'2022-04-01','2022-04-30','2022-05-02'),
			(2022,05,0,'2022-05-01','2022-05-31','2022-06-04'),
			(2022,06,0,'2022-06-01','2022-06-30','2022-07-06'),
			(2022,07,0,'2022-07-01','2022-07-31','2022-08-03'),
			(2022,08,0,'2022-08-01','2022-08-31','2022-09-05'),
			(2022,09,0,'2022-09-01','2022-09-30','2022-10-06'),
			(2022,10,0,'2022-10-01','2022-10-31','2022-11-03'),
			(2022,11,0,'2022-11-01','2022-11-30','2022-12-02'),
			(2022,12,0,'2022-12-01','2022-12-31','2023-01-05'),

			(2022,01,1,'2022-01-01','2022-01-31','2022-02-02'),
			(2022,02,1,'2022-02-01','2022-02-28','2022-03-05'),
			(2022,03,1,'2022-03-01','2022-03-31','2022-04-02'),
			(2022,04,1,'2022-04-01','2022-04-30','2022-05-03'),
			(2022,05,1,'2022-05-01','2022-05-31','2022-06-05'),
			(2022,06,1,'2022-06-01','2022-06-30','2022-07-04'),
			(2022,07,1,'2022-07-01','2022-07-31','2022-08-05'),
			(2022,08,1,'2022-08-01','2022-08-31','2022-09-06'),
			(2022,09,1,'2022-09-01','2022-09-30','2022-10-02'),
			(2022,10,1,'2022-10-01','2022-10-31','2022-11-04'),
			(2022,11,1,'2022-11-01','2022-11-30','2022-12-03'),
			(2022,12,1,'2022-12-01','2022-12-31','2023-01-02'),
	
			(2022,01,2,'2022-01-01','2022-01-27','2022-02-04'),
			(2022,02,2,'2022-02-01','2022-02-28','2022-03-02'),
			(2022,03,2,'2022-03-01','2022-03-30','2022-04-03'),
			(2022,04,2,'2022-04-01','2022-04-28','2022-05-01'),
			(2022,05,2,'2022-05-01','2022-05-31','2022-06-05'),
			(2022,06,2,'2022-06-01','2022-06-29','2022-07-06'),
			(2022,07,2,'2022-07-01','2022-07-31','2022-08-03'),
			(2022,08,2,'2022-08-01','2022-08-31','2022-09-03'),
			(2022,09,2,'2022-09-01','2022-09-30','2022-10-04'),
			(2022,10,2,'2022-10-01','2022-10-31','2022-11-02'),
			(2022,11,2,'2022-11-01','2022-11-30','2022-12-03'),
			(2022,12,2,'2022-12-01','2022-12-30','2023-01-05'),

			(2022,01,3,'2022-01-01','2022-01-26','2022-02-04'),
			(2022,02,3,'2022-02-01','2022-02-28','2022-03-05'),
			(2022,03,3,'2022-03-01','2022-03-27','2022-04-01'),
			(2022,04,3,'2022-04-01','2022-04-29','2022-05-05'),
			(2022,05,3,'2022-05-01','2022-05-30','2022-06-03'),
			(2022,06,3,'2022-06-01','2022-06-26','2022-07-05'),
			(2022,07,3,'2022-07-01','2022-07-30','2022-08-02'),
			(2022,08,3,'2022-08-01','2022-08-31','2022-09-03'),
			(2022,09,3,'2022-09-01','2022-09-30','2022-10-05'),
			(2022,10,3,'2022-10-01','2022-10-28','2022-11-06'),
			(2022,11,3,'2022-11-01','2022-11-29','2022-12-03'),
			(2022,12,3,'2022-12-01','2022-12-31','2023-01-04'),

			(2022,01,4,'2022-01-01','2022-01-27','2022-02-02'),
			(2022,02,4,'2022-02-01','2022-02-28','2022-03-05'),
			(2022,03,4,'2022-03-01','2022-03-30','2022-04-02'),
			(2022,04,4,'2022-04-01','2022-04-29','2022-05-03'),
			(2022,05,4,'2022-05-01','2022-05-30','2022-06-05'),
			(2022,06,4,'2022-06-01','2022-06-28','2022-07-04'),
			(2022,07,4,'2022-07-01','2022-07-27','2022-08-05'),
			(2022,08,4,'2022-08-01','2022-08-30','2022-09-06'),
			(2022,09,4,'2022-09-01','2022-09-26','2022-10-02'),
			(2022,10,4,'2022-10-01','2022-10-29','2022-11-04'),
			(2022,11,4,'2022-11-01','2022-11-28','2022-12-03'),
			(2022,12,4,'2022-12-01','2022-12-31','2023-01-02'),

			(2022,01,5,'2022-01-01','2022-01-29','2022-02-01'),
			(2022,02,5,'2022-02-01','2022-02-28','2022-03-04'),
			(2022,03,5,'2022-03-01','2022-03-30','2022-04-05'),
			(2022,04,5,'2022-04-01','2022-04-28','2022-05-03'),
			(2022,05,5,'2022-05-01','2022-05-30','2022-06-02'),
			(2022,06,5,'2022-06-01','2022-06-27','2022-07-04'),
			(2022,07,5,'2022-07-01','2022-07-30','2022-08-06'),
			(2022,08,5,'2022-08-01','2022-08-31','2022-09-05'),
			(2022,09,5,'2022-09-01','2022-09-27','2022-10-04'),
			(2022,10,5,'2022-10-01','2022-10-31','2022-11-03'),
			(2022,11,5,'2022-11-01','2022-11-26','2022-12-02'),
			(2022,12,5,'2022-12-01','2022-12-31','2023-01-04'),

			(2022,01,6,'2022-01-01','2022-01-28','2022-02-05'),
			(2022,02,6,'2022-02-01','2022-02-28','2022-03-04'),
			(2022,03,6,'2022-03-01','2022-03-27','2022-04-03'),
			(2022,04,6,'2022-04-01','2022-04-30','2022-05-04'),
			(2022,05,6,'2022-05-01','2022-05-26','2022-06-05'),
			(2022,06,6,'2022-06-01','2022-06-27','2022-07-02'),
			(2022,07,6,'2022-07-01','2022-07-31','2022-08-03'),
			(2022,08,6,'2022-08-01','2022-08-31','2022-09-04'),
			(2022,09,6,'2022-09-01','2022-09-29','2022-10-05'),
			(2022,10,6,'2022-10-01','2022-10-31','2022-11-04'),
			(2022,11,6,'2022-11-01','2022-11-28','2022-12-02'),
			(2022,12,6,'2022-12-01','2022-12-31','2023-01-03'),

			(2022,01,7,'2022-01-01','2022-01-27','2022-02-05'),
			(2022,02,7,'2022-02-01','2022-02-28','2022-03-06'),
			(2022,03,7,'2022-03-01','2022-03-30','2022-04-02'),
			(2022,04,7,'2022-04-01','2022-04-29','2022-05-03'),
			(2022,05,7,'2022-05-01','2022-05-31','2022-06-04'),
			(2022,06,7,'2022-06-01','2022-06-27','2022-07-02'),
			(2022,07,7,'2022-07-01','2022-07-30','2022-08-05'),
			(2022,08,7,'2022-08-01','2022-08-29','2022-09-06'),
			(2022,09,7,'2022-09-01','2022-09-26','2022-10-02'),
			(2022,10,7,'2022-10-01','2022-10-29','2022-11-03'),
			(2022,11,7,'2022-11-01','2022-11-30','2022-12-04'),
			(2022,12,7,'2022-12-01','2022-12-30','2023-01-05'),

			(2022,01,8,'2022-01-01','2022-01-28','2022-02-05'),
			(2022,02,8,'2022-02-01','2022-02-28','2022-03-04'),
			(2022,03,8,'2022-03-01','2022-03-30','2022-04-05'),
			(2022,04,8,'2022-04-01','2022-04-28','2022-05-03'),
			(2022,05,8,'2022-05-01','2022-05-29','2022-06-06'),
			(2022,06,8,'2022-06-01','2022-06-29','2022-07-02'),
			(2022,07,8,'2022-07-01','2022-07-27','2022-08-03'),
			(2022,08,8,'2022-08-01','2022-08-31','2022-09-07'),
			(2022,09,8,'2022-09-01','2022-09-28','2022-10-04'),
			(2022,10,8,'2022-10-01','2022-10-31','2022-11-05'),
			(2022,11,8,'2022-11-01','2022-11-26','2022-12-02'),
			(2022,12,8,'2022-12-01','2022-12-31','2023-01-03'),

			(2022,01,9,'2022-01-01','2022-01-30','2022-02-02'),
			(2022,02,9,'2022-02-01','2022-02-28','2022-03-04'),
			(2022,03,9,'2022-03-01','2022-03-28','2022-04-03'),
			(2022,04,9,'2022-04-01','2022-04-28','2022-05-05'),
			(2022,05,9,'2022-05-01','2022-05-30','2022-06-06'),
			(2022,06,9,'2022-06-01','2022-06-29','2022-07-03'),
			(2022,07,9,'2022-07-01','2022-07-28','2022-08-07'),
			(2022,08,9,'2022-08-01','2022-08-30','2022-09-04'),
			(2022,09,9,'2022-09-01','2022-09-30','2022-10-08'),
			(2022,10,9,'2022-10-01','2022-10-27','2022-11-02'),
			(2022,11,9,'2022-11-01','2022-11-30','2022-12-03'),
			(2022,12,9,'2022-12-01','2022-12-27','2023-01-04');		
		`)
	fmt.Print("tabla cierre insertada \n")

}

func InsertarDatos() {
	insertarComercios()
	insertarClientes()
	insertarTarjetas()
	insertarCierres()
	insertarConsumo()
	fmt.Printf("Datos insertados correctamente\n")
}
