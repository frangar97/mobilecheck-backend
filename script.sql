CREATE DATABASE mobilecheck

CREATE TABLE Usuario(
    id serial primary key,
    nombre varchar(50),
    apellido varchar(50),
    telefono varchar(50),
    email varchar(50),
    activo bool,
    usuario varchar(50),
    password text,
    web bool,
    movil bool
)

CREATE TABLE TipoVisita(
    id serial primary key,
    nombre varchar(50),
    color varchar(50),
    activo bool
)

CREATE TABLE Cliente(
    id serial primary key,
    nombre varchar(100),
    telefono varchar(50),
    email varchar(50),
    direccion varchar(100),
    activo bool,
    latitud float,
    longitud float,
    usuarioId int,
    foreign key(usuarioId) references Usuario(id)
)

CREATE TABLE Visita(
    id serial primary key,
    comentario text,
    latitud float,
    longitud float,
    fecha timestamp,
    imagen text,
    usuarioId int,
    clienteId int,
    tipoVisitaId int,
    foreign key(usuarioId) references Usuario(id),
    foreign key(clienteId) references Cliente(id),
    foreign key(tipoVisitaId) references TipoVisita(id)
)

CREATE TABLE Tarea(
    id serial primary key,
    descripcion text,
    fecha timestamp,
    completada bool,
    clienteId int,
    visitaId int,
    usuarioId int,
    foreign key(usuarioId) references Usuario(id),
    foreign key(visitaId) references Visita(id),
    foreign key(clienteId) references Cliente(id)
)