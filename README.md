# swiftedge project 

SwiftEdge es una herramienta CLI y SDK modular diseñada para monitorear bases de datos MySQL/MariaDB y recolectar métricas del sistema, como el uso de CPU, memoria, disco y conexiones activas. 

Este proyecto está inspirado y basado en el trabajo original de [Luis Contreras](https://github.com/luiscontrerasdo/golang-mariadbconsole/tree/main), con mejoras significativas en modularidad, extensibilidad y usabilidad.

## Características

- **Métricas del Sistema**: Recolección en tiempo real de métricas de CPU, memoria y disco.
- **Monitoreo de Base de Datos**:
  - Conexiones activas
  - Estado de replicación
  - Versión de la base de datos
  - Consultas lentas
- **Interfaz Gráfica en Terminal**: Construida con `termui` para mostrar métricas en tiempo real.
- **Modularidad**: Diseño extensible como SDK para integrarse en otros proyectos.

## Requisitos Previos

- **Go**: Versión 1.20 o superior.
- **MySQL/MariaDB**: Configuración de acceso remoto habilitada.
- **Dependencias**: El proyecto utiliza las siguientes bibliotecas:
  - [`github.com/shirou/gopsutil`](https://github.com/shirou/gopsutil): Recolección de métricas del sistema.
  - [`github.com/gizak/termui/v3`](https://github.com/gizak/termui): Interfaz gráfica en terminal.
  - [`github.com/spf13/cobra`](https://github.com/spf13/cobra): Framework para CLI.
  - [`github.com/go-sql-driver/mysql`](https://github.com/go-sql-driver/mysql): Controlador para MySQL/MariaDB.

## Instalación

1. **Clona el Repositorio**:

   ```bash
   git clone https://github.com/jclanas2019/swiftedge.git
   cd swiftedge
   ```

   ```bash
   go mod tidy
  ```

  ```bash
go build -o swiftedge
```
```bash
./swiftedge monitor --user=<usuario> --password=<contraseña> --host=<host> --database=<base_de_datos>
```
Ejemplo:
```bash
./swiftedge monitor --user=root --password=1234 --host=localhost --database=testdb
```


