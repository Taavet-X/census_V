package models
import 
"gorm.io/gorm"

type CensusData struct {
	gorm.Model
	ID            uint   `gorm:"primaryKey"`
	Age           int    `gorm:"column:age"`
	Workclass     string `gorm:"column:workclass"`
	Education     string `gorm:"column:education"`
	MaritalStatus string `gorm:"column:marital_status"` // Revisar que este nombre coincida
	Occupation    string `gorm:"column:occupation"`     // Revisar que este nombre coincida
	Relationship  string `gorm:"column:relationship"`
	Race          string `gorm:"column:race"`
	Sex           string `gorm:"column:sex"`
	HoursPerWeek  int    `gorm:"column:hours_per_week"`
	Income        string `gorm:"column:income"`
}
