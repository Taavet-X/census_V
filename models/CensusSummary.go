package models
import "gorm.io/gorm"

// CensusSummary define los campos para el resumen de los datos censales
type CensusSummary struct {
	gorm.Model
	Education string `json:"education"`
	Count     int    `json:"count"`
}
