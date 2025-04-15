package api

import (
	"campaign_engine/internal/db"
	"campaign_engine/internal/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
	r.GET("/campaigns", getAllCampaigns)
	r.GET("/bids", getAllBids)
}

func getAllCampaigns(c *gin.Context) {
	var campaigns []models.Campaign
	if err := db.DB.Find(&campaigns).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch campaigns"})
		return
	}
	c.JSON(http.StatusOK, campaigns)
}

func getAllBids(c *gin.Context) {
	var bids []models.Bid
	if err := db.DB.Find(&bids).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch bids"})
		return
	}
	c.JSON(http.StatusOK, bids)
}
