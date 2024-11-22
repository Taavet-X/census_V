package controllers

import (
	"censusV/database"
	"censusV/models"
	"encoding/csv"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ExportData(c *gin.Context) {
	var censusData []models.CensusData

	// Obtener los datos
	if err := database.DB.Find(&censusData).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al obtener los datos del censo"})
		return
	}

	// Configurar encabezados de respuesta
	fileName := "census_data.csv"
	c.Header("Content-Disposition", fmt.Sprintf("attachment; filename=%s", fileName))
	c.Header("Content-Type", "text/csv")

	// Crear CSV writer
	writer := csv.NewWriter(c.Writer)
	defer writer.Flush()

	// Escribir encabezados
	headers := []string{
		"ID", "Edad", "Clase de Trabajo", "Educación", "Estado Civil",
		"Ocupación", "Relación", "Raza", "Sexo", "Horas por Semana", "Ingreso",
	}
	if err := writer.Write(headers); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al escribir encabezados del CSV"})
		return
	}

	// Escribir datos
	for _, data := range censusData {
		row := []string{
			fmt.Sprintf("%d", data.ID),
			fmt.Sprintf("%d", data.Age),
			data.Workclass,
			data.Education,
			data.MaritalStatus,
			data.Occupation,
			data.Relationship,
			data.Race,
			data.Sex,
			fmt.Sprintf("%d", data.HoursPerWeek),
			data.Income,
		}
		if err := writer.Write(row); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al escribir datos en el CSV"})
			return
		}
	}
}