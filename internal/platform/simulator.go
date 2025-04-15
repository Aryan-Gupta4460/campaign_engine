package platform

import (
	"campaign_engine/internal/db"
	"campaign_engine/internal/models"
	"campaign_engine/internal/utils"

	"github.com/lib/pq"
)

func SeedDemoCampaigns() {
	var count int64
	db.DB.Model(&models.Campaign{}).Count(&count)
	if count == 0 {
		utils.LogInfo("Seeding demo campaigns...")

		demoCampaigns := []models.Campaign{
			{
				ID:        "cmp123",
				Budget:    1000,
				ReachGoal: 5000,
				Platforms: pq.StringArray{"Meta", "Google"},
			},
			{
				ID:        "cmp456",
				Budget:    800,
				ReachGoal: 2000,
				Platforms: pq.StringArray{"Google"},
			},
			{
				ID:        "cmp789",
				Budget:    1500,
				ReachGoal: 10000,
				Platforms: pq.StringArray{"Meta", "TikTok"},
			},
		}

		for _, c := range demoCampaigns {
			db.DB.Create(&c)
		}

		utils.LogInfo("Demo campaigns seeded successfully!")
	} else {
		utils.LogInfo("Campaigns already exist. Skipping seeding.")
	}
}
