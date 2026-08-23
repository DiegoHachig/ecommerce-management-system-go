# Ecommerce Management System Go

Sistema de Gestión de E-commerce desarrollado en Go (Golang) aplicando Programación Funcional, Programación Orientada a Objetos, Concurrencia, Persistencia de Datos y Servicios Web.

---

# Universidad Internacional del Ecuador

## Programación Orientada a Objetos 2-CIB-3B

### Proyecto Integrador Final

**Autor:** Diego Alexander Hachig Zapata

**Fecha:** Agosto 2026

**Repositorio GitHub:**

https://github.com/DiegoHachig/ecommerce-management-system-go

---

# Introducción

El presente proyecto integrador consiste en el desarrollo de un Sistema de Gestión de E-commerce utilizando Go (Golang). El sistema permite administrar usuarios, productos, inventario, pedidos y reportes mediante una arquitectura modular basada en buenas prácticas de desarrollo de software.

Durante el desarrollo se aplicaron los conocimientos adquiridos en las cuatro unidades de la asignatura, incorporando funciones, paquetes, estructuras de datos, programación orientada a objetos, encapsulación, interfaces, concurrencia, persistencia de datos, serialización JSON y servicios web.

El proyecto fue desarrollado utilizando Visual Studio Code y administrado mediante GitHub para el control de versiones.

---

# Objetivo General

Desarrollar un Sistema de Gestión de E-commerce capaz de administrar productos, usuarios, inventario y pedidos mediante una arquitectura modular desarrollada en Go, aplicando los conocimientos adquiridos durante las cuatro unidades de la asignatura.

---

# Objetivos Específicos

- Gestionar usuarios y perfiles.
- Gestionar el catálogo de productos.
- Administrar el inventario disponible.
- Gestionar pedidos realizados por los clientes.
- Generar reportes básicos del sistema.
- Aplicar programación funcional en la lógica de negocio.
- Implementar Programación Orientada a Objetos mediante structs, interfaces y encapsulación.
- Utilizar persistencia de datos mediante MySQL y GORM.
- Implementar servicios web utilizando JSON.
- Aplicar concurrencia mediante goroutines y canales.

---

# Módulos Implementados

## Módulo de Usuarios

Funcionalidades:

- Registro de usuarios.
- Consulta de usuarios.
- Actualización de información.
- Persistencia en MySQL.

---

## Módulo de Productos

Funcionalidades:

- Registro de productos.
- Modificación de productos.
- Eliminación de productos.
- Consulta de catálogo.
- Persistencia en MySQL.

---

## Módulo de Inventario

Funcionalidades:

- Registro de existencias.
- Consulta de stock.
- Actualización de inventario.
- Verificación de disponibilidad.

---

## Módulo de Pedidos

Funcionalidades:

- Creación de pedidos.
- Consulta de pedidos.
- Actualización de estados.
- Gestión de órdenes.

---

## Módulo de Reportes

Funcionalidades:

- Reporte de ventas.
- Reporte de inventario.
- Historial de operaciones.
- Información consolidada del sistema.

---

# Arquitectura del Proyecto

```text
ecommerce-management-system-go
│
├── cmd
│   └── main.go
│
├── configs
│
├── database
│   └── connection.go
│
├── internal
│   ├── users
│   ├── products
│   ├── inventory
│   ├── orders
│   └── reports
│
├── routes
│   └── routes.go
│
├── tests
│
├── go.mod
├── go.sum
└── README.md
```

---

# Tecnologías Utilizadas

- Go (Golang)
- Visual Studio Code
- GitHub
- Gin Framework
- GORM
- MySQL
- JSON
- REST API

---

# Persistencia de Datos

El sistema utiliza MySQL junto con GORM para almacenar y gestionar información.

Características implementadas:

- Conexión a MySQL.
- Migraciones automáticas.
- Operaciones CRUD.
- Gestión de entidades.
- Persistencia de usuarios.
- Persistencia de productos.

Estado actual de persistencia:

| Módulo | Persistencia |
|---------|---------|
| Usuarios | MySQL |
| Productos | MySQL |
| Pedidos | Memoria |
| Inventario | Memoria |
| Reportes | Memoria |

---

# Integración de las Cuatro Unidades

## Unidad 1 - Fundamentos de Go

Temas aplicados:

- Sintaxis.
- Condicionales.
- Estructuras de control.
- Funciones.
- Paquetes.

Aplicación:

La lógica del sistema se desarrolló mediante funciones reutilizables organizadas en paquetes independientes para cada módulo.

---

## Unidad 2 - Estructuras de Datos y Objetos

Temas aplicados:

- Arrays.
- Slices.
- Maps.
- Structs.
- Métodos.
- Constructores.

Aplicación:

Se desarrollaron las principales entidades del sistema:

- User
- Product
- Order
- InventoryRecord
- ProductEntity
- UserEntity

---

## Unidad 3 - Programación Orientada a Objetos

Temas aplicados:

- Encapsulación.
- Interfaces.
- Polimorfismo.
- Manejo de errores.

Interfaces implementadas:

- ProductRepository
- UserRepository
- OrderRepository
- InventoryRepository

Manejo de errores:

- errors.New()
- fmt.Errorf()

Encapsulación:

Se utilizaron atributos privados y métodos getter para acceso controlado a la información.

---

## Unidad 4 - Concurrencia, Testing y Web

Temas aplicados:

- Concurrencia.
- Goroutines.
- Canales.
- Servicios Web.
- Serialización JSON.
- Testing.

Aplicación:

El sistema se encuentra preparado para implementar procesamiento concurrente mediante goroutines y canales, permitiendo futuras mejoras de escalabilidad y rendimiento.

Se desarrolló una arquitectura orientada a servicios utilizando estructuras compatibles con APIs REST y serialización JSON.

---

# Servicios Web

El Sistema de Gestión de E-commerce fue diseñado con una arquitectura modular que permite evolucionar hacia una solución basada en Servicios Web REST.

Los servicios web facilitan la comunicación entre aplicaciones mediante protocolos HTTP y permiten que sistemas web, móviles o externos consuman información del sistema de forma segura y estandarizada.

## Servicios Disponibles

### Usuarios

#### Consultar Usuarios

```http
GET /users
```

Obtiene el listado de usuarios registrados.

#### Registrar Usuario

```http
POST /users
```

Registra un nuevo usuario en el sistema.

---

### Productos

#### Consultar Productos

```http
GET /products
```

Obtiene el catálogo de productos registrados.

#### Registrar Producto

```http
POST /products
```

Registra un nuevo producto.

---

### Inventario

#### Consultar Inventario

```http
GET /inventory
```

Permite visualizar las existencias disponibles.

#### Registrar Inventario

```http
POST /inventory
```

Permite agregar registros de inventario.

---

### Pedidos

#### Consultar Pedidos

```http
GET /orders
```

Obtiene los pedidos registrados.

#### Crear Pedido

```http
POST /orders
```

Permite registrar un nuevo pedido.

---

## Resumen de Servicios Web

| Endpoint | Método | Funcionalidad |
|-----------|----------|-------------|
| /users | GET | Consultar usuarios |
| /users | POST | Registrar usuario |
| /products | GET | Consultar productos |
| /products | POST | Registrar producto |
| /inventory | GET | Consultar inventario |
| /inventory | POST | Registrar inventario |
| /orders | GET | Consultar pedidos |
| /orders | POST | Crear pedido |

**Total de servicios web definidos: 8**

---

# Serialización JSON

La comunicación entre los servicios web utiliza JSON (JavaScript Object Notation), permitiendo intercambiar información de forma estructurada y compatible con diferentes plataformas.

## Ejemplo de Producto

```json
{
  "id": 1,
  "name": "Laptop Lenovo",
  "price": 850.00,
  "stock": 10
}
```

## Ejemplo de Usuario

```json
{
  "id": 1,
  "name": "Diego Hachig",
  "email": "diego@email.com",
  "phone": "1234567890"
}
```

## Ejemplo de Pedido

```json
{
  "id": 1,
  "userId": 1,
  "total": 850.00,
  "status": "Pendiente"
}
```

### Ventajas del uso de JSON

- Formato ligero y fácil de interpretar.
- Compatibilidad con aplicaciones web y móviles.
- Integración sencilla con APIs REST.
- Amplio soporte en diferentes lenguajes de programación.
- Facilita el intercambio de información entre sistemas.

---

# Cronograma de Desarrollo

| Semana | Unidad | Tema | Actividad Realizada |
|----------|----------|----------|----------|
| Semana 1 | Unidad 1 | ¿Qué es Go?, Sintaxis y Condicionales | Selección del sistema de gestión y planificación inicial del proyecto. |
| Semana 2 | Unidad 1 | Funciones y Paquetes | Diseño modular y organización de paquetes del sistema. |
| Semana 3 | Unidad 2 | Arrays, Slices y Maps | Implementación de estructuras de datos para almacenar información. |
| Semana 4 | Unidad 2 | Objetos en Go: Structs, Métodos y Constructores | Desarrollo de entidades y modelos del sistema. |
| Semana 5 | Unidad 3 | Encapsulación y Manejo de Errores | Aplicación de getters, validaciones y control de errores. |
| Semana 6 | Unidad 3 | Interfaces y Polimorfismo | Diseño desacoplado mediante repositorios e interfaces. |
| Semana 7 | Unidad 4 | Concurrencia, Goroutines y Channels | Preparación de procesos concurrentes para mejorar rendimiento y escalabilidad. |
| Semana 8 | Unidad 4 | Servicios Web, JSON y Testing | Integración final del sistema, documentación y preparación para entrega. |

---

# Visualización del Futuro

La evolución futura del sistema contempla:

- Implementación completa de APIs REST.
- Aplicación móvil para clientes.
- Integración con Inteligencia Artificial para análisis de ventas.
- Dashboards de Business Intelligence.
- Infraestructura en la nube.
- Facturación electrónica.
- Arquitectura basada en microservicios.
- Procesamiento concurrente avanzado mediante Go.

---

# Aplicaciones Prácticas

El sistema puede utilizarse en:

- Pequeñas empresas.
- Tiendas virtuales.
- Negocios de venta en línea.
- Gestión de inventarios.
- Control de pedidos.
- Administración de catálogos de productos.

---

# Conclusiones

- Se desarrolló un Sistema de Gestión de E-commerce utilizando Go.
- Se aplicaron los conocimientos adquiridos durante las cuatro unidades de la asignatura.
- Se implementó una arquitectura modular basada en buenas prácticas de desarrollo.
- Se integraron conceptos de programación funcional, POO, persistencia de datos y servicios web.
- El proyecto constituye una base sólida para futuras ampliaciones empresariales y tecnológicas.

---

# Licencia

Proyecto académico desarrollado para la asignatura Programación Orientada a Objetos.

Universidad Internacional del Ecuador.