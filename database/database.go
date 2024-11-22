package database

import (
	"censusV/models"
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

// CreateDbConnection establece la conexión a la base de datos.
func CreateDbConnection() *gorm.DB {
	connStr := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_DB"),
		os.Getenv("DB_SSL"),
	)

	db, err := gorm.Open(postgres.Open(connStr), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error al conectar a la base de datos: %v", err)
	}
	return db
}


func InitDB() {
	db := CreateDbConnection()

	// Migrar las tablas
	if err := db.AutoMigrate(
		&models.CensusData{}, &models.CensusSummary{}, &models.User{}); err != nil {
		log.Fatalf("Error al migrar las tablas: %v", err)
	}

	log.Println("Tablas migradas exitosamente.")
	DB = db
}
