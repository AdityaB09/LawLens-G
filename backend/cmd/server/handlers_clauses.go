package main

import (
	"net/http"
	"strconv"

	"lawlens-g/internal/services"

	"github.com/gin-gonic/gin"
)

func (app *AppContext) GetContractClauses(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil || id <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	clauses, err := services.GetClausesForContract(app.DB, uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to load clauses"})
		return
	}

	c.JSON(http.StatusOK, clauses)
}

func (app *AppContext) GetContractObligations(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil || id <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	obligations, err := services.GetObligationsForContract(app.DB, uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to load obligations"})
		return
	}

	c.JSON(http.StatusOK, obligations)
}
