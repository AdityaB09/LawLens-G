package main

import (
	"net/http"

	"lawlens-g/internal/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type AppContext struct {
	DB *gorm.DB
}

func SetupRouter(db *gorm.DB) *gin.Engine {
	r := gin.Default()
	r.Use(corsMiddleware())

	appCtx := &AppContext{DB: db}

	r.GET("/api/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok", "service": "lawlens-g"})
	})

		api := r.Group("/api")
	{
		api.POST("/contracts", appCtx.CreateContract)
		api.POST("/contracts/pdf", appCtx.CreateContractFromPDF) // NEW

		api.GET("/contracts", appCtx.ListContracts)
		api.GET("/contracts/:id", appCtx.GetContract)
		api.GET("/contracts/:id/clauses", appCtx.GetContractClauses)
		api.GET("/contracts/:id/obligations", appCtx.GetContractObligations)
		api.POST("/contracts/compare", appCtx.CompareContracts)
		api.GET("/contracts/:id/risk-summary", appCtx.GetRiskSummary)
	}


	return r
}

func corsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET,POST,OPTIONS")
		if c.Request.Method == http.MethodOptions {
			c.AbortWithStatus(200)
			return
		}
		c.Next()
	}
}

// Small helper to preload associations where needed
func preloadContract(db *gorm.DB) *gorm.DB {
	return db.Preload("Clauses").Preload("Obligations").Preload("Clauses.Triggers")
}

// just to ensure imports are used
var _ = models.Contract{}
