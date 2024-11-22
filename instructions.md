# Contexto:

Gracias a un dataset creado por Barry Becker basado en un censo de 1994 que contiene datos demográficos, tenemos los datos necesarios para comprender qué tipo de personas ganaban más o menos de $50,000 USD.
El problema es que los datos recopilados sólo están disponibles en un archivo de texto plano y actualmente carecemos de una herramienta que nos permita visualizarlos y derivar hipótesis y conclusiones a partir de esa información.

## Antes de comenzar.

Renombra este archivo como instructions.md para que puedas agregar tu propio README.md

La fuente de datos se encuentra en la carpeta data. Es un archivo de texto plano con más de 30,000 registros por lo que como primer recomendación pasarla a una base de datos estructurada sería lo adecuado. Otras recomendaciones:

1. Utiliza una base de datos Postgres.
2. El uso de enums podría serte útil.
3. Al ser una gran cantidad de datos por cargar a la DB, podrías considerar utilizar [goroutines](https://medium.com/@sumairz/golang-intro-to-goroutines-39fe44d6bf8d) que utilizamos también durante [el curso](https://github.com/codeableorg/trucode3-go-live-coding/tree/main/16.concurrency).

## Historia de Usuario 1: Como usuario, quiero filtrar y ordenar los datos para obtener información específica.

- Contexto:
  - necesitamos una herramienta que nos permita ver, filtrar y ordenar los datos para obtener información específica sobre ingresos, educación, estado civil, ocupación, edad y nivel de ingresos.
- Descripción:
  - Quisiera poder visualizar los datos en una tabla.
  - Desearía poder filtrar los datos por educación, estado civil, ocupación, rango de edad y nivel de ingresos.
  - Necesito la capacidad de ordenar los datos de manera ascendente o descendente según cualquiera de los atributos de la tabla.
  - Necesito poder ver los elementos con paginación para tener una mejor presentación de los mismos.
- Criterios de Aceptación:
  - Para optimizar la velocidad de los resultados, ejecutar los filtros y la paginación en el backend.
  - Debe haber opciones de filtrado visible en la interfaz.
  - Los datos deben reaccionar dinámicamente a los filtros seleccionados.

## Historia de Usuario 2: Como usuario registrado, quiero poder autenticarme para persistir mis configuraciones de visualización.

- Descripción:
  - Quiero tener la opción de registrarme e iniciar sesión para guardar mis preferencias de visualización -filtros y ordenamiento-.
  - La autenticación debería permitirme acceder a mis configuraciones personalizadas desde cualquier dispositivo.
- Criterios de Aceptación:
  - La interfaz debe tener opciones de registro e inicio de sesión.
  - Después de iniciar sesión, mis configuraciones personalizadas deben ser cargadas automáticamente.

## Historia de Usuario 3: Como usuario, quiero que la aplicación sea accesible desde diferentes dispositivos.

- Contexto
  - necesitamos una aplicación que sea accesible y funcional en diversos dispositivos, incluyendo escritorios, tablets y teléfonos móviles.
- Descripción:
  - La aplicación debe ser responsive y funcionar correctamente en escritorios, tablets y teléfonos móviles.
  - Quiero una experiencia de usuario consistente independientemente del dispositivo que esté utilizando.
- Criterios de Aceptación:
  - La aplicación debe ser accesible y usable en diferentes tamaños de pantalla.
  - Todas las funcionalidades deben ser fácilmente utilizables en dispositivos móviles.

## Historia de Usuario 4 -opcional-: Como usuario, quiero poder exportar datos censales.

- Contexto:
  - necesitamos herramientas para exportar datos de acuerdo a la información filtrada y ordenada.
- Descripción:
  - quiero la opción de exportar los datos en formato CSV para su análisis externo.
- Criterios de Aceptación:
  - Debería haber una opción clara para importar y actualizar datos.
  - La exportación a CSV debe proporcionar un archivo descargable con todos los datos seleccionados.

## Historia de Usuario 5 -opcional-: Como usuario, quiero ver estadísticas resumidas para comprender mejor los datos.

- Contexto:
  - necesitamos obtener estadísticas resumidas para comprender mejor la distribución demográfica, incluyendo edad promedio y distribución de ingresos por nivel de educación y ocupación.
- Descripción:
  - Quiero ver estadísticas resumidas como edad promedio y horas de trabajo promedio.
  - Además, deseo obtener insights sobre la distribución de ingresos dentro de diferentes niveles de educación y ocupaciones.
- Criterios de Aceptación:
  - Las estadísticas resumidas deben ser claramente visibles en la interfaz.
  - Las visualizaciones deben proporcionar información sobre la distribución de ingresos según la educación y la ocupación.

## Historia de Usuario 6 -opcional-: Como usuario, quiero ver visualizaciones interactivas de datos para comprender la distribución demográfica.

- Descripción:
  - Al acceder a la aplicación, deseo ver gráficos interactivos que representen la distribución de datos demográficos.
- Criterios de Aceptación:
  - La página principal debe mostrar al menos cuatro visualizaciones interactivas.

## Implementación y Verificaciones Finales

- Despliegue
  - Despliega la aplicación en cualquier plataforma para que sea accesible desde la web.
- Verificaciones Finales y Documentación
  - Documenta cualquier paso adicional necesario para el despliegue.
  - Asegúrate de que la aplicación cumpla con los requisitos establecidos como obligatorios y cualesquiera de los opcionales.

¡Éxito!
