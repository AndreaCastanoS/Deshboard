# Proyecto de Análisis de Datos Demográficos

## **Descripción**

Este proyecto es una herramienta para analizar y visualizar datos demográficos a partir de un dataset basado en el censo de 1994 creado por Barry Becker. Utiliza **Vue** para el frontend, **Go** para el backend, y **PostgreSQL** como base de datos.

## **La aplicación permite a los usuarios:**

- Filtrar y ordenar los datos de manera eficiente.
- Visualizar estadísticas y gráficos interactivos sobre los datos.
- Guardar configuraciones personalizadas de visualización.

## **Funcionalidades**

- **Registro e inicio de sesión**: Los usuarios pueden registrarse e iniciar sesión para persistir sus configuraciones de visualización.
- **Filtrado y ordenación de datos**: Los usuarios pueden filtrar y ordenar los registros por cualquier campo (como education, occupation, etc.) en orden ascendente o descendente.
- **Paginación**: Los datos se muestran con paginación para facilitar la visualización de grandes volúmenes de registros.
- **Estadísticas resumidas**: La aplicación muestra el promedio de edad y el promedio de horas trabajadas por semana de los registros.
- **Gráficas interactivas**: Se muestran gráficos que evidencian las ocupaciones y los niveles educativos con mayores ingresos.
- **Configuraciones personalizadas**: Los usuarios pueden guardar sus filtros personalizados, y estos se aplicarán automáticamente cuando inicien sesión nuevamente.

## **Instalación y Configuración**

### **Requisitos previos**

- **Go**: Asegúrate de tener Go instalado en tu máquina.
- **Node.js y npm**: Necesarios para el frontend de Vue.

### **Paso 1: Instalación de dependencias en el backend (Go)**

1. Clona el repositorio:
   ```
   git clone https://github.com/AndreaCastanoS/trucode3-challenge-final
   ```
   
Navega al directorio del backend:

 ```
   cd backend_challenge
```

Instala las dependencias necesarias. Para manejar las dependencias en Go, utiliza go mod:

```
go mod tidy
```

Paso 2: Instalación de dependencias en el frontend (Vue.js)

Navega al directorio del frontend:

```
cd frontend_challenge
```

Instala las dependencias:
```
npm install
```

Paso 3: Ejecutar la aplicación
Backend (Go): Navega al directorio del backend y ejecuta el servidor:

   ```
   go run .
   ```

Frontend (Vue): Navega al directorio del frontend y ejecuta el servidor de desarrollo:

```
npm run dev
```

Esto iniciará la aplicación.

## **Despliegue**
La aplicación está desplegada y accesible en la web en el siguiente enlace: http://trucode.s3-website.us-east-2.amazonaws.com
