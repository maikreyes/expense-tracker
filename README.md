# 💰 Expense Tracker CLI

![Go Version](https://img.shields.io/badge/go-1.25.5-00ADD8?style=flat-square&logo=go)
![License](https://img.shields.io/badge/license-MIT-green?style=flat-square)
![Build Status](https://img.shields.io/badge/build-passing-brightgreen?style=flat-square)

Una herramienta de línea de comandos (CLI) simple y eficiente para gestionar tus finanzas personales. Desarrollada en **Go**, permite realizar un seguimiento de gastos, actualizarlos, eliminarlos y ver resúmenes detallados.

Este proyecto es una solución al desafío [Expense Tracker](https://roadmap.sh/projects/expense-tracker) de **roadmap.sh**.

## ✨ Características

* **Añadir gastos:** Registra nuevos gastos con una descripción y un monto.
* **Listar gastos:** Visualiza una tabla con todos los gastos registrados (ID, fecha, descripción, monto).
* **Actualizar:** Modifica la descripción o el monto de un gasto existente.
* **Eliminar:** Borra gastos específicos mediante su ID.
* **Resumen:** Muestra el total de gastos acumulados.
* **Filtro por mes:** Consulta el total de gastos de un mes específico (ej: todo lo gastado en Agosto).
* **Persistencia:** Los datos se guardan automáticamente en un archivo JSON local (`expenses.json`).

## 🚀 Requisitos e Instalación

Este proyecto fue desarrollado y probado con **Go 1.25.5**.

### Prerrequisitos

* Tener [Go 1.25.5](https://go.dev/dl/) (o superior) instalado en tu sistema.

### Instalación

1.  **Clona el repositorio:**

    ```bash
    git clone [https://github.com/maikreyes/expense-tracker.git](https://github.com/maikreyes/expense-tracker.git)
    cd expense-tracker
    ```

2.  **Compila el proyecto:**

    ```bash
    go build -o expense-tracker
    ```

    > **Nota:** En Windows, esto generará un archivo `expense-tracker.exe`.

## 📖 Uso

Una vez compilado, puedes usar la herramienta directamente desde tu terminal.

### 1. Añadir un nuevo gasto

```bash
./expense-tracker add --description "Almuerzo" --amount 20
# Output: Expense added successfully (ID: 1)
```

### 2. Listar todos los gastos

```Bash
./expense-tracker list
# ID  Date       Description  Amount
# 1   2024-08-01 Almuerzo     $20
# 2   2024-08-02 Café         $5
```

### 3. Ver resumen total de gastos

```Bash
./expense-tracker summary
# Total expenses: $25
```

### 4. Ver resumen de un mes específico

Usa el número del mes (1-12). Ejemplo para Agosto:

```Bash
./expense-tracker summary --month 8
# Total expenses for August: $25
```

### 5. Eliminar un gasto

```Bash
./expense-tracker delete --id 1
# Output: Expense deleted successfully
```

🛠️ Tecnologías

Lenguaje: Go (v1.25.5)

Almacenamiento: Archivo JSON local

🤝 Contribuciones
Las sugerencias y Pull Requests son bienvenidos. Si encuentras algún error, por favor abre un issue en el repositorio.

📄 Licencia
Este proyecto está bajo la licencia MIT.
