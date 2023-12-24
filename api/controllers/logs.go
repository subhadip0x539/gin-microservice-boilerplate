package controllers

import (
	"fmt"
	"go-gin-boilerplate/api/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type LogsController struct{}

func (controller LogsController) Index(context *gin.Context) {
	var request models.Index

	if err := context.ShouldBindJSON(&request); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	fmt.Println(request.Data)

	context.JSON(http.StatusOK, gin.H{"message": "Successful"})
}
