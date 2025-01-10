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
2. Instala las Dependencias:
   ```bash
   go mod tidy
  ```
3. Compila el Proyecto:
  ```bash
go build -o swiftedge
```
4. Ejecuta el Comando CLI:
```bash
./swiftedge monitor --user=<usuario> --password=<contraseña> --host=<host> --database=<base_de_datos>
```
Ejemplo:
```bash
./swiftedge monitor --user=root --password=1234 --host=localhost --database=testdb
```

Uso

CLI
El comando principal es monitor, que inicia el monitoreo de métricas de la base de datos y del sistema en tiempo real. Opciones disponibles:

```bash
./swiftedge monitor --user=<usuario> --password=<contraseña> --host=<host> --database=<base_de_datos>
```
Como SDK
Puedes integrar SwiftEdge en tus proyectos importando el módulo correspondiente:

```go
package main

import (
    "log"
    "swiftedge/sdk/monitor"
)

func main() {
    err := monitor.Run("root", "1234", "localhost", "testdb")
    if err != nil {
        log.Fatalf("Error iniciando el monitor: %v", err)
    }
}
```
Contribuciones

¡Contribuciones son bienvenidas! Por favor, sigue estos pasos para colaborar:

1. Haz un fork del repositorio.
2. Crea una rama para tus cambios:
```bash
git checkout -b feature/nueva-funcionalidad
```
3. Realiza un commit con tus cambios:
```bash
git commit -m "Añadir nueva funcionalidad"
```
4. Sube tus cambios:
```bash
git push origin feature/nueva-funcionalidad
```   
5. Abre un pull request en el repositorio principal.

Licencia

Este proyecto está bajo la licencia MIT. 


