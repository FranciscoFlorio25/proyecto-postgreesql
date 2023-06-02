create or replace  function  autorizarTarjeta (buscado_nrotarjeta char(16)) returns bool as $$
	declare
		resultado record;
	begin
		select * into resultado
   	 	from tarjeta
		where tarjeta.nrotarjeta = buscado_nrotarjeta ;
		if resultado.estado = 'vigente' then
			return true;
		else
			
			return false;
		end if;
	end;
	$$ language plpgsql;

create or replace function autorizarCodigoSeg(buscado_nrotarjeta char(16),codseguridad_buscado char(4)) returns bool as $$
	declare
		resultado record;
	begin
		select * into resultado
   	 	from tarjeta
		where tarjeta.nrotarjeta = buscado_nrotarjeta;
	
		if codseguridad_buscado != resultado.codseguridad then
			return false;
		else
			return true;
		end if;
	end;
	$$ language plpgsql;


create or replace function tarjeta_Suspendida(buscado_nrotarjeta char(16)) returns text as $$
 	declare
		resultado record;
		suspendida text;
	begin
		select * into resultado
		from tarjeta
		where tarjeta.nrotarjeta = buscado_nrotarjeta;

		if resultado.estado = 'suspendida' then
			suspendida= resultado.estado;
			return suspendida;
		end if;
		return null ;
	end;
	$$ language plpgsql;
	
create or replace function autorizarSuspendida(buscado_nrotarjeta char(16)) returns bool as $$
 	declare
		resultado record;
	begin
		select * into resultado
		from tarjeta
		where tarjeta.nrotarjeta = buscado_nrotarjeta;

		if resultado.estado = 'suspendida' then
			return false;
		else
			return true;
		end if;
	end;
	$$ language plpgsql;

create or replace function autorizarLimiteCompra(buscado_tarjeta char(16), 
							montoPedido decimal(7, 2)) returns bool as $$
	declare
		pendiente record;
--		compra_pendiente decimal;
		limite record;
--
	begin
		select cast(sum(monto) as integer) into pendiente
		from compra where compra.nrotarjeta = buscado_tarjeta and compra.pagado='f';
--		compra_pendiente:=pendiente.sum;
		raise notice 'monto: %', pendiente.sum;
	
		select limitecompra into limite
		from tarjeta where tarjeta.nrotarjeta = buscado_tarjeta;
		raise notice 'limite: %',limite.limitecompra;
		
		if (pendiente.sum+montoPedido) >= limite.limitecompra then
			raise notice ' su nuevo monto dara: %', pendiente.sum+montoPedido;
			return false;
		else
			return true;
		end if;

	end;
	$$ language plpgsql;


create or replace function autorizarVigencia(nrotarjeta_buscado char(16),fecha_buscada date) returns bool as $$
	declare
		resultado record;
		tarjetaBuscada record;
		ultimo_digito integer;
		mes_digitos integer;
		año_digitos integer;
	begin
		select * into tarjetaBuscada from tarjeta where tarjeta.nrotarjeta=nrotarjeta_buscado;
		select right(nrotarjeta,1) into ultimo_digito from tarjeta where tarjeta.nrotarjeta=nrotarjeta_buscado;
		raise notice 'tarjeta: %',tarjetabuscada.nrotarjeta ;
		raise notice 'ultimo digito %', ultimo_digito;
		select right(validahasta,2) into mes_digitos from tarjeta where tarjeta.nrotarjeta=nrotarjeta_buscado;
		select left(validahasta,4) into año_digitos from tarjeta where tarjeta.nrotarjeta=nrotarjeta_buscado;
		raise notice 'mes %',mes_digitos;
		raise notice 'año %',año_digitos;
		select *
		into resultado 
		from cierre
		where  mes_digitos=cierre.mes  and ultimo_digito= cierre.terminacion;
		raise notice ' fecha vto en resultado %',resultado.fechavto;
		raise notice ' resultado mes: %', resultado.mes;

	--	end if;
		if resultado.fechavto >= fecha_buscada then
			raise notice 'no se vencio, fecha ingresada es menor a fecha Vto';
			return false;
		else
			raise notice 'la fecha si se expiro';
			
			return true;
		end if;
	end;
	$$ language plpgsql;

create or replace function contarfilas() returns integer as $func$
	declare
		filas record;
		fila integer;
	begin 
	--	count:=select count(nrorechazo) from rechazo;
		select count(nrorechazo)
				into filas 
				from rechazo;
				
		fila=filas.count+1;
	--if fila==0 then
	--	fila=1;
--	end if;
	return fila;
	end;
	$func$ language plpgsql;

		
create or replace function autoriza(nrotarjeta_buscado char(16),codseguridad_buscado char(4),
										nrocomercio_buscado int,monto_buscado decimal(7,2)) returns bool as  $$ 

	declare
		tarjeta varchar(16);
		resultado record;
		fecha date;
		nrorechazo_compra int :=1;
		incrementa int;
		suspendida_tarj text;
	begin
		select *
		into resultado 
		from consumo
    	where 
    	monto_buscado =consumo.monto
    	and nrotarjeta_buscado = consumo.nrotarjeta
    	and nrocomercio_buscado = consumo.nrocomercio
    	and codseguridad_buscado = consumo.codseguridad;
    	tarjeta:=resultado.nrotarjeta;
    	suspendida_tarj=tarjeta_Suspendida(nrotarjeta_buscado);
 --	select * 
   --- 	into resultado
   -- 	from tarjeta
   -- 	where tarjeta.nrotarjeta = buscado_nrotarjeta;

    	--select nrotarjeta into tarjeta from consumo where nrotarjeta=nrotarjeta_buscado;
    --	tarjeta:=	select nrotarjeta from consumo where nrotarjeta=nrotarjeta_buscado;
    	raise notice ' la tarjeta es: %',resultado.nrotarjeta;
--		
		raise notice ' el nrorechazo_compra es: %',nrorechazo_compra;
	--	if nrorechazo_compra=1 then
--			nrorechazo_compra=nrorechazo_compra+1;
--				raise notice ' el nrorechazo_compra es: %',nrorechazo_compra;
--		end if;
    	select CURRENT_TIMESTAMP into fecha;
		--tarjeta:= select nrotarjeta into consumo where nrotarjeta=nrotarjeta_buscado;
    	if autorizarTarjeta(nrotarjeta_buscado)='f' then
    		incrementa:=contarfilas();
    		raise notice ' registro de rechazo es: %',incrementa;

    		if suspendida_tarj='suspendida' then
				insert into rechazo values (incrementa,nrotarjeta_buscado,nrocomercio_buscado,fecha,monto_buscado,'la tarjeta se encuentra suspendida');
    			return false;
    		end if;
    		insert into rechazo values (incrementa,nrotarjeta_buscado,nrocomercio_buscado,fecha,monto_buscado,'?tarjeta no vàlida ò no vigente');
    		return false;
    	end if;	
 
    	if autorizarCodigoSeg(nrotarjeta_buscado,codseguridad_buscado)='f' then
    		raise notice 'su codigo es falso %',tarjeta;
    		insert into rechazo values (contarfilas(),nrotarjeta_buscado,nrocomercio_buscado,fecha,monto_buscado,'?còdigo de seguridad invàlido');
    		return false;
    	end if;

    	if autorizarLimiteCompra(nrotarjeta_buscado,monto_buscado)='f' then
			insert into rechazo values (contarfilas(),nrotarjeta_buscado,nrocomercio_buscado,fecha,monto_buscado,'?Supera lìmite de tarjeta');
			return false;
		end if;
		
 	  	if autorizarVigencia(nrotarjeta_buscado,fecha)='f' then
    		insert into rechazo values (nrorechazo_compra,nrotarjeta_buscado,nrocomercio_buscado,fecha,monto_buscado,'?plazo de vigencia expirado');
   			return false;

		end if;

		
--		if autorizarSuspendida(resultado.nrotarjeta)=false then
--			insert into rechazo values (nrorechazo_compra,resultado.nrotarjeta,resultado.nrocomercio,fecha,resultado.monto,'la tarjeta se encuentra suspendida');
--			return false;
  --  	end if;
--
--    	if autorizarLimiteCompra(nrotarjeta_buscado,monto_buscado)='f' then
--			insert into rechazo values (contarfilas(),nrotarjeta_buscado,nrocomercio_buscado,fecha,monto_buscado,'?Supera lìmite de tarjeta');
--			return false;
--		end if;
	
 		insert into compra values (nrorechazo_compra,resultado.nrotarjeta,resultado.nrocomercio,fecha,resultado.monto);
    	return true;
	end;
	$$ language plpgsql;

create or replace function autorizar(nrotarjeta_buscado char(16),codseguridad_buscado char(4),
										nrocomercio_buscado int,monto_buscado decimal(7,2)) returns bool as  $$ 
	declare 
		autorizado boolean;
		fecha date;
	begin
		select CURRENT_TIMESTAMP into fecha;
		autorizado:=autoriza(nrotarjeta_buscado,codseguridad_buscado,nrocomercio_buscado ,monto_buscado);
	if autorizado='t' then
		raise notice 'aca se guarda a compra';
		insert into compra values(contarfilasCompra(),nrotarjeta_buscado,nrocomercio_buscado,fecha,monto_buscado,false);
		return true;
	else 
		return false;
	end if;
	return false;
	end;
	$$ language plpgsql;


create or replace function contarfilasCompra() returns integer as $func$
	declare
		filas record;
		fila integer;
	begin 
	--	count:=select count(nrorechazo) from rechazo;
		select count(nrooperacion)
				into filas 
				from compra;
				
		fila=filas.count+1;
	--if fila==0 then
	--	fila=1;
--	end if;
	
	return fila;
	end;
	$func$ language plpgsql;
