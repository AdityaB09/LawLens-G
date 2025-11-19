package main

import (
	"net/http"
	"strconv"
	"time"
    "strings"
	"lawlens-g/internal/models"
	"lawlens-g/internal/services"

	"github.com/gin-gonic/gin"
)

type CreateContractRequest struct {
	Title   string `json:"title" binding:"required"`
	PartyA  string `json:"partyA" binding:"required"`
	PartyB  string `json:"partyB" binding:"required"`
	Text    string `json:"text" binding:"required"` // raw contract text
}

type ContractSummaryResponse struct {
	ID                uint      `json:"id"`
	Title             string    `json:"title"`
	PartyA            string    `json:"partyA"`
	PartyB            string    `json:"partyB"`
	UploadedAt        time.Time `json:"uploadedAt"`
	OverallRiskLevel  string    `json:"overallRiskLevel"`
	OverallRiskScore  float64   `json:"overallRiskScore"`
	ClauseCount       int       `json:"clauseCount"`
	ObligationCount   int       `json:"obligationCount"`
}

func (app *AppContext) CreateContract(c *gin.Context) {
	var req CreateContractRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid payload", "details": err.Error()})
		return
	}

	contract, err := services.CreateContractWithAnalysis(app.DB, req.Title, req.PartyA, req.PartyB, req.Text)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create contract", "details": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"id": contract.ID})
}

func (app *AppContext) ListContracts(c *gin.Context) {
	var contracts []models.Contract
	if err := app.DB.Preload("Clauses").Preload("Obligations").
		Order("uploaded_at DESC").
		Find(&contracts).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to list contracts"})
		return
	}

	resp := make([]ContractSummaryResponse, 0, len(contracts))
	for _, ct := range contracts {
		resp = append(resp, ContractSummaryResponse{
			ID:               ct.ID,
			Title:            ct.Title,
			PartyA:           ct.PartyA,
			PartyB:           ct.PartyB,
			UploadedAt:       ct.UploadedAt,
			OverallRiskLevel: ct.OverallRiskLevel,
			OverallRiskScore: ct.OverallRiskScore,
			ClauseCount:      len(ct.Clauses),
			ObligationCount:  len(ct.Obligations),
		})
	}

	c.JSON(http.StatusOK, resp)
}

func (app *AppContext) GetContract(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil || id <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	var contract models.Contract
	if err := preloadContract(app.DB).First(&contract, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "contract not found"})
		return
	}

	c.JSON(http.StatusOK, contract)
}

func (app *AppContext) GetRiskSummary(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil || id <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	summary, err := services.ComputeRiskSummary(app.DB, uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to compute risk summary"})
		return
	}

	c.JSON(http.StatusOK, summary)
}

func (app *AppContext) CreateContractFromPDF(c *gin.Context) {
	title := c.PostForm("title")
	partyA := c.PostForm("partyA")
	partyB := c.PostForm("partyB")

	if title == "" || partyA == "" || partyB == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "title, partyA, partyB are required"})
		return
	}

	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "PDF file is required"})
		return
	}

	tmpPath := "/tmp/" + file.Filename
	if err := c.SaveUploadedFile(file, tmpPath); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to save PDF"})
		return
	}

	text, err := services.ReadPDFToText(tmpPath)
	if err != nil || strings.TrimSpace(text) == "" {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to extract text from PDF"})
		return
	}

	contract, err := services.CreateContractWithAnalysis(app.DB, title, partyA, partyB, text)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create contract from PDF"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"id": contract.ID})
}

