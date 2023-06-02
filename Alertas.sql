
create or replace function contarFilasAlertas() returns integer as $func$
	declare
		filas record;
		fila integer;
	begin
		select count(nroresumen) into filas from cabecera;
		fila = filas.count+1;
		return fila;
	end;
	$func$ language plpgsql;
create or replace function diffmin(fecha1 timestamp, fecha2 timestamp) returns int as $$
    declare 
        minutos1 int;
        minutos2 int;
        diffmin int;
    begin
        select date_part('minutes',fecha1::time) into minutos1 as int;
        select date_part('minutes',fecha2::time) into minutos2 as int;
        if minutos2 - minutos1 <0 then
            diffmin = 0;
        else
            diffmin = minutos2 - minutos1;
        return diffmin;
        end if;
    end;
    $$ language plpgsql;

create or replace function diffhoras(fecha1 timestamp, fecha2 timestamp) returns int as $$
    declare 
        horas1 int;
        horas2 int;
        diffhoras int;
    begin
        select date_part('hours',fecha1::time) into horas1 as int;
        select date_part('hours',fecha2::time) into horas2 as int;
        if horas2 - horas1 <0 then
            diffhoras = 0;
        else
            diffhoras = horas2 - horas1;
        end if;
        return diffhoras;
    end;
    $$ language plpgsql;
    
create or replace function insertarAlerta() returns trigger as $$
    begin
        insert into alerta
        values (contarfilas(),NEW.nrotarjeta,NEW.fecha,NEW.nrorechazo,0,'Fallo validaciòn');
        return new;
    end;
    $$ language plpgsql;
    
create or replace  trigger insertarAlerta
after INSERT on rechazo
for each STATEMENT
execute procedure insertarAlerta();


create or replace function  alertaTimer1MIN() returns trigger as $$
    declare
        resultado record;
        resul_comercio record;
    begin
        select c.codigopostal,comp.nrotarjeta, comp.fecha into resultado 
        from comercio c, compra comp 
        where comp.nrocomercio = c.nrocomercio
        and NEW.nrotarjeta = comp.nrotarjeta;

        select comercio.codigopostal into result_comercio  from comercio where resultado.codigpostal = comercio.codigopostal;
        
        if resultado.codigopostal = result_ comercio.codigopostal and resultado.nrocomercio != NEW.nrocomercio and diffmin(resultado.fecha,NEW.fecha) < 1 then
            insert into alerta
            values (contarFilasAlertas(),NEW.nrotarjeta,NEW.fecha,NEW.nrorechazo,1,'Se quizo realizar dos compras en menos de 1min');
        end if;
        return new;
    end;
    $$ language plpgsql;
create or replace  trigger alertaTimer1MIN
before INSERT on compra
for each STATEMENT
execute procedure  alertaTimer1MIN();


create or replace function  alertaTimer5Min() returns trigger as $$
    declare
    	result_comercio record;
        resultado record;
    begin
        select c.codigopostal,comp.nrotarjeta, comp.fecha into resultado 
        from comercio c, compra comp 
        where comp.nrocomercio = c.nrocomercio
        and NEW.nrotarjeta = comp.nrotarjeta; 
		select comercio.codigopostal into result_comercio  from comercio where resultado.codigpostal
		
        if resultado.codigopostal != result_comercio.codigopostal and diffmin(resultado.fecha,NEW.fecha) < 5 then
            insert into alerta 
            values (contarFilasAlertas(),NEW.nrotarjeta,NEW.fecha,NEW.nrorechazo,5,'Se quizo realizar dos compras en menos de 5min');
        end if;
        return new;
    end;
    $$ language plpgsql;
    
create or replace  trigger alertaTimer5Min
before INSERT on compra
for each STATEMENT
execute procedure alertaTimer5Min();


create or replace function rechazosSeguidos() returns trigger as $$
    declare
    	tarjetas record;
        resultado record;
    begin
        select * into resultado from rechazo where NEW.nrotarjeta = rechazo.nrotarjeta;
        select * into tarjetas from tarjeta where resultado.nrotarjeta = tarjeta.nrotarjeta;

        if diffhoras(resultado.fecha,NEW.fecha) < 24 and resultado.motivo = '?supera lìmite de tarjeta' then
            insert into alerta
            values (contarFilasAlertas(),NEW.nrotarjeta,NEW.fecha,NEW.nrorechazo,32,'Se registraron dos rechazos el mismo dia');
            UPDATE  tarjetas.tarjeta set tarjetas.estado = 'suspendida' where tarjetas.nrotarjeta = resultado.nrotarjeta;
        end if;
        return new;
    end;
    $$ language plpgsql;
    
create or replace trigger rechazosSeguidos
before INSERT on rechazo
for each STATEMENT
execute procedure rechazosSeguidos();
