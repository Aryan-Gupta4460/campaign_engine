package campaign

import (
	"campaign_engine/internal/bidder"
	"campaign_engine/internal/db"
	"campaign_engine/internal/models"
	"campaign_engine/internal/utils"
	"time"
)

func StartCampaignSimulation() {
	ticker := time.NewTicker(10 * time.Second)

	go func() {
		for range ticker.C {
			var campaigns []models.Campaign
			if err := db.DB.Find(&campaigns).Error; err != nil {
				utils.LogError(err)
				continue
			}

			for _, c := range campaigns {
				bid := bidder.GenerateBid(c)
				db.DB.Create(&bid)
			}
		}
	}()
}
