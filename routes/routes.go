package routes

import (
	"censusV/controllers"
	"censusV/middlewares"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	// Middleware para manejar CORS
	router.Use(middlewares.CORSMiddleware())

	// Ruta principal para verificar el estado del servidor
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Welcome to Census Data API!",
		})
	})

	// Rutas de autenticación y manejo de usuarios
	userRoutes := router.Group("/api")
	{
		userRoutes.POST("/register", controllers.RegisterUser) // Registro de usuarios
		userRoutes.POST("/login", controllers.LoginUser)       // Inicio de sesión
	}

	// Rutas para manejar datos censales
	censusRoutes := router.Group("/api/census")
	{
		censusRoutes.GET("/", controllers.GetCensusData)     // Filtrado, ordenamiento y paginación
		censusRoutes.GET("/summary", controllers.GetSummary) // Resumen de datos
		censusRoutes.GET("/export", controllers.ExportData)  // Exportación a CSV
		//censusRoutes.GET("/visuals", controllers.GetVisuals) // Visualizaciones interactivas
	}

	return router
}
