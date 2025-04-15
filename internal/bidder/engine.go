package bidder

import (
	"campaign_engine/internal/analytics"
	"campaign_engine/internal/kafka"
	"campaign_engine/internal/models"
	"campaign_engine/internal/utils"
	"encoding/json"
	"time"
)

func GenerateBid(c models.Campaign) models.Bid {
	cpc := analytics.PredictCPC()
	cvr := analytics.PredictCVR()

	bid := models.Bid{
		CampaignID:   c.ID,
		Platform:     c.Platforms[0],
		Amount:       cpc * 1.1,
		EstimatedCVR: cvr,
		CreatedAt:    time.Now(),
	}

	data, _ := json.Marshal(bid)
	if err := kafka.SendBidMessage(data); err != nil {
		utils.LogErrorf("Kafka error:", err)
	}

	return bid
}
