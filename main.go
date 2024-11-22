package main

import (
	"log"

	"censusV/database"
	"censusV/loaders"

	"github.com/joho/godotenv"
)

var loadDataOnRun bool = true

func main() {

	// Cargar las variables de entorno
	if err := godotenv.Load(); err != nil {
		log.Println("No se encontró el archivo .env, usando variables de entorno del sistema.")
	}

	// Inicializar la base de datos
	log.Println("Inicializando la base de datos...")
	database.InitDB()

	if loadDataOnRun {
		// Proporciona la ruta completa al archivo source_data.csv
		const filePath = "./data/test.csv"
		// read a file
		data := loaders.ReadData(filePath)
		// Se intenta guardar en la base de datos
		if err := loaders.SaveDataIntoDB(data); err != nil {
			return
		}
	}

    /*
	// Cargar los datos del censo
	log.Println("Cargando los datos del censo desde 'source_data.csv'...")
	if err := loaders.LoadCensusData("source_data.csv"); err != nil {
		log.Fatalf("Error al cargar los datos del censo: %v", err)
	}

	// Proporciona la ruta completa al archivo source_data.csv
	filePath := "censusV/data/source.data.csv"

	// Cargar datos del censo desde el archivo source_data.csv
	log.Println("Cargando los datos del censo desde 'source_data.csv'...")
	if err := loaders.LoadCensusData(filePath); err != nil {
		log.Fatalf("Error al cargar los datos del censo: %v", err)
	} else {
		log.Println("Datos del censo cargados correctamente")
	}

	// Configurar el router
	router := routes.SetupRouter()

	// Obtener el puerto del entorno o usar el puerto 8080 por defecto
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// Iniciar el servidor
	log.Printf("Servidor corriendo en el puerto %s", port)
	if err := router.Run(":" + port); err != nil {
		log.Fatalf("Error al iniciar el servidor: %v", err)
	}
	*/
}
