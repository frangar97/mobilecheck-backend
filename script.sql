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
    nombre varchar(50),
    telefono varchar(50),
    email varchar(50),
    direccion varchar(100),
    activo bool,
    latitud float,
    longitud float,
    usuarioId int,
    foreign key(usuarioId) references Usuario(id)
)