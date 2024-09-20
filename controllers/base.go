package controllers

import "gorm.io/gorm"

var DB *gorm.DB

type ErrorResponse struct {
	Message string `json:"message"`
}

type SuccessResponse struct {
	Message string `json:"message"`
}
