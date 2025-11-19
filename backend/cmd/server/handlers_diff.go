package main

import (
	"net/http"

	"lawlens-g/internal/services"

	"github.com/gin-gonic/gin"
)

type CompareRequest struct {
	ContractA uint `json:"contractA"`
	ContractB uint `json:"contractB"`
}

func (app *AppContext) CompareContracts(c *gin.Context) {
	var req CompareRequest
	if err := c.ShouldBindJSON(&req); err != nil || req.ContractA == 0 || req.ContractB == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid payload"})
		return
	}

	diff, err := services.CompareContracts(app.DB, req.ContractA, req.ContractB)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to compare contracts", "details": err.Error()})
		return
	}

	c.JSON(http.StatusOK, diff)
}
