package main

import (
	"campaign_engine/internal/api"
	"campaign_engine/internal/campaign"
	"campaign_engine/internal/db"
	"campaign_engine/internal/kafka"
	"campaign_engine/internal/platform"
	"campaign_engine/internal/utils"
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	utils.InitLogger()

	db.InitPostgres()
	kafka.InitProducer()

	r := gin.Default()
	api.RegisterRoutes(r)
	platform.SeedDemoCampaigns()
	campaign.StartCampaignSimulation()

	port := ":8085"
	fmt.Printf(" Server is running on http://localhost%s\n", port)
	r.Run(port)
}
