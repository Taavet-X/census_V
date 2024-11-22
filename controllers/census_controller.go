package controllers

import (
	"censusV/database"
	"censusV/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetCensusData(c *gin.Context) {
	var data []models.CensusData
	db := database.DB

	// Filtrado
	query := db
	if education := c.Query("education"); education != "" {
		query = query.Where("education = ?", education)
	}
	if income := c.Query("income"); income != "" {
		query = query.Where("income = ?", income)
	}

	// Ordenamiento
	if order := c.Query("order_by"); order != "" {
		direction := c.DefaultQuery("direction", "asc")
		query = query.Order(order + " " + direction)
	}

	// Paginación
	limit, err := strconv.Atoi(c.DefaultQuery("limit", "10"))
	if err != nil || limit < 1 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Parámetro 'limit' inválido"})
		return
	}
	offset, err := strconv.Atoi(c.DefaultQuery("offset", "0"))
	if err != nil || offset < 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Parámetro 'offset' inválido"})
		return
	}
	query = query.Limit(limit).Offset(offset)

	// Consultar resultados
	if err := query.Find(&data).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al obtener datos del censo"})
		return
	}

	c.JSON(http.StatusOK, data)
}

func GetSummary(c *gin.Context) {
	var summary []struct {
		Education string `json:"education"`
		Count     int    `json:"count"`
	}

	if err := database.DB.Model(&models.CensusData{}).
		Select("education, COUNT(*) as count").
		Group("education").
		Scan(&summary).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al obtener el resumen de datos"})
		return
	}

	c.JSON(http.StatusOK, summary)
}
