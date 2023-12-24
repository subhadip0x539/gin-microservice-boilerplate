package models

type Index struct {
	Data string `json:"data" binding:"required"`
}
