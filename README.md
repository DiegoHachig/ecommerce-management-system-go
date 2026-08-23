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

## Ecommerce Management System 2030

La evolución futura del sistema contempla la transformación de la aplicación actual hacia una plataforma empresarial inteligente basada en tecnologías modernas.

```text
Sistema Actual
       │
       ▼
Servicios Web REST
       │
       ▼
Aplicación Web y Móvil
       │
       ▼
Computación en la Nube
       │
       ▼
Business Intelligence
       │
       ▼
Inteligencia Artificial
       │
       ▼
Microservicios
       │
       ▼
Analítica Predictiva
```

### Tecnologías Futuras Consideradas

- APIs REST completas para integración entre sistemas.
- Aplicaciones móviles para acceso remoto.
- Infraestructura en la nube.
- Bases de datos distribuidas.
- Inteligencia Artificial para predicción de ventas.
- Business Intelligence para análisis de información.
- Arquitectura basada en microservicios.
- Automatización de procesos empresariales.

---

# Explicación de la Visualización del Futuro

La visualización muestra la evolución tecnológica proyectada para el Sistema de Gestión de E-commerce.

Actualmente el sistema permite administrar usuarios, productos, inventario, pedidos y reportes mediante una arquitectura modular desarrollada en Go.

Como siguiente etapa, el sistema podría incorporar servicios web REST que permitan la comunicación con aplicaciones externas y clientes web.

Posteriormente se podrían desarrollar aplicaciones móviles para facilitar el acceso desde cualquier dispositivo.

La incorporación de infraestructura en la nube permitiría mejorar la disponibilidad, seguridad y escalabilidad del sistema, facilitando el crecimiento de la plataforma.

Mediante herramientas de Business Intelligence y Analítica de Datos, la organización podría obtener indicadores de rendimiento, conocer tendencias de consumo y mejorar la toma de decisiones.

Finalmente, la integración de Inteligencia Artificial permitiría generar predicciones de ventas, recomendaciones de productos y automatización de procesos, convirtiendo la aplicación en una solución empresarial inteligente y escalable.
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

# Justificación del Proyecto

## Criterio de Selección del Tema

Se seleccionó el desarrollo de un Sistema de Gestión de E-commerce debido a la creciente importancia del comercio electrónico en los entornos empresariales actuales.

Las organizaciones requieren herramientas que permitan administrar productos, usuarios, inventarios y pedidos de forma eficiente, reduciendo tiempos operativos y mejorando el control de la información.

Además, este tipo de sistema permite aplicar de forma práctica los conceptos estudiados durante la asignatura, integrando programación funcional, estructuras de datos, programación orientada a objetos, persistencia de datos y servicios web.

---

# Aplicaciones Prácticas

El sistema desarrollado puede ser utilizado en diferentes contextos empresariales y comerciales.

Entre sus principales aplicaciones se encuentran:

- Tiendas virtuales.
- Comercios electrónicos.
- Negocios de venta de productos físicos.
- Gestión de inventarios.
- Control de pedidos.
- Administración de clientes.
- Seguimiento de ventas.
- Generación de reportes operativos.

La arquitectura implementada también permite futuras ampliaciones para adaptarse a diferentes necesidades organizacionales.

---

# Reflexión sobre lo Aprendido

Durante el desarrollo del proyecto se aplicaron los conocimientos adquiridos en las cuatro unidades de la asignatura.

La implementación del sistema permitió comprender la importancia de la organización del código mediante paquetes y funciones, el uso de estructuras de datos para el manejo de información y la aplicación de principios de Programación Orientada a Objetos como encapsulación e interfaces.

Asimismo, se fortalecieron conocimientos relacionados con persistencia de datos utilizando MySQL y GORM, así como la preparación de la arquitectura para futuras implementaciones de servicios web y concurrencia.

---

# Dificultades Encontradas

Durante el desarrollo se presentaron algunos desafíos técnicos:

- Organización de la arquitectura del proyecto.
- Implementación de interfaces y repositorios.
- Integración con MySQL mediante GORM.
- Gestión de dependencias en Go.
- Diseño de estructuras reutilizables.
- Coordinación de los diferentes módulos del sistema.

Estos desafíos permitieron adquirir experiencia práctica en el desarrollo de aplicaciones empresariales utilizando Go.

---

# Aplicaciones Futuras

Como líneas de mejora futura se plantean:

- Implementación completa de APIs REST.
- Aplicación móvil para clientes.
- Integración con servicios en la nube.
- Facturación electrónica.
- Dashboards de Business Intelligence.
- Inteligencia Artificial para análisis predictivo.
- Arquitectura basada en microservicios.

Estas mejoras permitirán transformar el sistema en una solución empresarial más completa y escalable.

---

# Ambiente de Desarrollo y Pruebas

Para la implementación de los temas correspondientes a la Unidad 4 (Concurrencia, Servicios Web, JSON y Testing), se utilizará un entorno de pruebas independiente con el objetivo de preservar la estabilidad de la aplicación principal.

Esta estrategia permite validar nuevas funcionalidades sin afectar los módulos operativos de Usuarios, Productos, Inventario, Pedidos y Reportes que actualmente se encuentran funcionando correctamente.

Las pruebas realizadas en este entorno incluyen:

- Implementación de Servicios Web REST.
- Serialización de datos mediante JSON.
- Pruebas unitarias.
- Pruebas de concurrencia con Goroutines.
- Validación de nuevas funcionalidades antes de su integración definitiva.

---

# Estrategia de Desarrollo

Con el objetivo de preservar la estabilidad del sistema principal, las actividades correspondientes a la Unidad 4 (Concurrencia, Servicios Web, JSON y Testing) se desarrollan en un entorno de pruebas independiente.

Para este propósito se creó la rama:

feature/unidad4-testing

Esta rama permite implementar y validar nuevas funcionalidades sin afectar la versión estable del proyecto.

Una vez concluidas las pruebas y verificaciones necesarias, las mejoras serán integradas al proyecto principal mediante Git y GitHub.

Esta estrategia garantiza la continuidad operativa del sistema y facilita una evolución controlada del software.

---
# Concurrencia

Como parte de la Unidad 4 se desarrolló un ambiente de pruebas para la implementación de concurrencia utilizando Goroutines y Channels.

Funcionalidades evaluadas:

- Ejecución concurrente de tareas.
- Comunicación entre procesos mediante canales.
- Generación simultánea de reportes.

Tecnologías utilizadas:

- Goroutines.
- Channels.
- Package time.

Ejemplo conceptual:

Reporte de Ventas
       │
       ▼
 Goroutine
       │
       ▼
 Channel
       │
       ▼
 Resultado

Reporte de Inventario
       │
       ▼
 Goroutine
       │
       ▼
 Channel
       │
       ▼
 Resultado

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