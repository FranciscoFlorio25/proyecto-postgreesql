
create or replace function resumenPeriodo(nrocliente_buscado int, fechaDesde date, fechaHasta date)
returns void as $$
	declare
		nombreCliente record;
		apellidoCliente record;
		domicilioCliente record;
		nrotarjetaCliente char(16);
		fechaVtoTarjeta date;
		terminacionTarjeta char(1);
		totalCompras decimal(7, 2);
		resultado record;
	begin
		select nombre into nombreCliente
		from cliente where cliente.nrocliente = nrocliente_buscado; 
		
		select apellido into apellidoCliente
		from cliente where cliente.nrocliente = nrocliente_buscado;	
	
		select domicilio into domicilioCliente
		from cliente where cliente.nrocliente = nrocliente_buscado;		
	
		select nrotarjeta into nrotarjetaCliente
		from tarjeta where tarjeta.nrocliente = nrocliente_buscado;

		select right(nrotarjetaCliente, 1) into terminacionTarjeta from tarjeta  where  
		tarjeta.nrocliente = nrocliente_buscado;

		select fechaVto into fechaVtoTarjeta 
		from cierre where cierre.terminacion = obtenerNumero(terminacionTarjeta)  and 
		cierre.fechainicio = fechaDesde and cierre.fechacierre = fechaHasta;

		select * into resultado from compra where compra.nrotarjeta = nrotarjetaCliente;
		if resultado.pagado = 'f' then
			select sum(monto) into totalCompras from compra where compra.nrotarjeta = nrotarjetaCliente;
		end if;		
		
		insert into cabecera values(contarFilas(), nombreCliente, apellidoCliente, domicilioCliente, 
					nrotarjetaCliente, fechaDesde, fechaHasta, fechaVtoTarjeta, totalCompras); 
	end;
	$$ language plpgsql;	 

create or replace function obtenerNombreComercio(nrodelcomercio int)
returns text as $$
	declare 
	nombreComercio text;
--	nrotarjetaCliente char(16);
	
	begin
	--	select nrotarjeta into nrotarjetaCliente
--			from tarjeta where tarjeta.nrocliente = nrocliente_buscado;
	--	select nrocomercio into nroDelComercio from compra where compra.nrotarjeta = nrotarjetaCliente;
		select nombre into nombreComercio from comercio where comercio.nrocomercio = nroDelComercio;
		
	return nombreComercio;
	end;
	$$ language plpgsql;
create or replace function comprasDelPeriodo(nrocliente_buscado int, fechaDesde date, fechaHasta date)
returns void as $$
	declare
		nombreComercio record; 
		nroDelComercio integer; 
		montoCompra decimal(7,2);
		fechaCompra date; 
		nrotarjetaCliente char(16);
		resultado record;

		indice record;
	begin
		select nrotarjeta into nrotarjetaCliente
		from tarjeta where tarjeta.nrocliente = nrocliente_buscado;
		select fecha into fechaCompra from compra where compra.nrotarjeta = nrotarjetaCliente;
		select nrocomercio into nroDelComercio from compra where compra.nrotarjeta = nrotarjetaCliente;
		raise notice ' nro tajet ckliente  :%  ',nroDelComercio;
		select nombre into nombreComercio from comercio where comercio.nrocomercio = nroDelComercio;

		select monto into montoCompra from compra where compra.nrotarjeta = nrotarjetaCliente;
	--	for counter in 1..6 by 2 loop 
	--	raise notice 'counter %', counter;
	--	end loop;


		for indice in select * from compra 
					where compra.fecha between fechaDesde and fechaHasta
	--	insert into detalle values(contarFilasDet(), contarLinea(), fechaCompra, nombreComercio, montoCompra);
		loop
			raise notice ' indice: %',indice.nrocomercio;
			insert into detalle values(contarFilasDet(), contarLinea(), indice.fecha,obtenerNombreComercio(indice.nrocomercio) , indice.monto);
	
		end loop;

	end;
	$$ language plpgsql;

create or replace function obtenerNumero(numero char(1))
returns int as $$
	begin
	if numero = '0' then
		return 0;
	end if;
	if numero = '1' then
		return 1;
	end if;
	if numero = '2' then
		return 2;
	end if;
	if numero = '3' then
		return 3;
	end if;
	if numero = '4' then
		return 4;
	end if;
	if numero = '5' then
		return 5;
	end if;
	if numero = '6' then
		return 6;
	end if;
	if numero = '7' then
		return 7;
	end if;
	if numero = '8' then
		return 8;
	end if;
	if numero = '9' then
		return 9;
	end if;
	end;
	$$ language plpgsql;

create or replace function contarFilas() returns integer as $func$
	declare
		filas record;
		fila integer;
	begin
		select count(nroresumen) into filas from cabecera;
		fila = filas.count+1;
		return fila;
	end;
	$func$ language plpgsql;

create or replace function contarFilasDet() returns integer as $func$
	declare
		filas record;
		fila integer;
	begin
		select count(nroresumen) into filas from detalle;
		fila = filas.count+1;
		return fila;
	end;
	$func$ language plpgsql;


create or replace function contarLinea() returns integer as $func$
	declare
		filas record;
		fila integer;
	begin
		select count(nrolinea) into filas from detalle;
		fila = filas.count+1;
		return fila;
	end;
	$func$ language plpgsql;

