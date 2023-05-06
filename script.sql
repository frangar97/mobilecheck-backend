CREATE TABLE public.cliente (
	id serial4 NOT NULL,
	codigocliente varchar(100) NULL,
	nombre varchar(100) NOT NULL,
	telefono varchar(50) NULL,
	email varchar(50) NULL,
	direccion varchar(150) NULL,
	activo bool NOT NULL DEFAULT true,
	latitud float8 NULL,
	longitud float8 NULL,
	usuariocrea int4 NULL,
	fechacrea timestamptz NULL,
	usuariomodifica int4 NULL,
	fechamodifica timestamptz NULL,
	CONSTRAINT cliente_pkey PRIMARY KEY (id)
);

CREATE TABLE public.tarea (
	id serial4 NOT NULL,
	fecha timestamptz NULL,
	completada bool NULL,
	clienteid int4 NULL,
	visitaid int4 NULL,
	usuarioid int4 NULL,
	imagenrequerida bool NOT NULL DEFAULT false,
	tipovisitaid int4 NULL,
	meta text NULL,
	CONSTRAINT tarea_pkey PRIMARY KEY (id)
);

CREATE TABLE public.tipovisita (
	id serial4 NOT NULL,
	nombre varchar(50) NULL,
	color varchar(50) NULL,
	activo bool NULL,
	requieremeta bool NULL DEFAULT false,
	CONSTRAINT tipovisita_pkey PRIMARY KEY (id)
);

CREATE TABLE public.usuario (
	id serial4 NOT NULL,
	nombre varchar(50) NULL,
	apellido varchar(50) NULL,
	telefono varchar(50) NULL,
	email varchar(50) NULL,
	activo bool NULL,
	usuario varchar(50) NULL,
	"password" text NULL,
	web bool NULL,
	movil bool NULL,
	CONSTRAINT usuario_pkey PRIMARY KEY (id)
);


CREATE TABLE public.visita (
	id serial4 NOT NULL,
	comentario text NULL,
	latitud float8 NULL,
	longitud float8 NULL,
	fecha timestamptz NULL,
	imagen text NULL,
	usuarioid int4 NULL,
	clienteid int4 NULL,
	meta text NULL,
	CONSTRAINT visita_pkey PRIMARY KEY (id)
);

ALTER TABLE public.visita ADD CONSTRAINT visita_clienteid_fkey FOREIGN KEY (clienteid) REFERENCES public.cliente(id);
ALTER TABLE public.visita ADD CONSTRAINT visita_usuarioid_fkey FOREIGN KEY (usuarioid) REFERENCES public.usuario(id);
ALTER TABLE public.tarea ADD CONSTRAINT tarea_clienteid_fkey FOREIGN KEY (clienteid) REFERENCES public.cliente(id);
ALTER TABLE public.tarea ADD CONSTRAINT tarea_tipovisitaid_fkey FOREIGN KEY (tipovisitaid) REFERENCES public.tipovisita(id);
ALTER TABLE public.tarea ADD CONSTRAINT tarea_usuarioid_fkey FOREIGN KEY (usuarioid) REFERENCES public.usuario(id);
ALTER TABLE public.tarea ADD CONSTRAINT tarea_visitaid_fkey FOREIGN KEY (visitaid) REFERENCES public.visita(id);

alter table Visita alter column fecha type timestamp WITH TIME ZONE USING fecha AT TIME ZONE 'GTM'
alter table Tarea alter column fecha type timestamp WITH TIME ZONE USING fecha AT TIME ZONE 'GTM'

INSERT INTO public.usuario(
 nombre, apellido, telefono, email, activo, usuario, password, web, movil)
	VALUES ('admin', 'admin','', '', true, 'admin', '$2a$10$xIWwKcbQCZ9wAqkt7SZ7yOcqmwDPFkWypOpNGY9HQpudTOJGRy222', true, true);