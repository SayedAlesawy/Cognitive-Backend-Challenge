package main

import (
	campaign "Cognitive-Backend-Challenge/api/model"
	dbwrapper "Cognitive-Backend-Challenge/utils/database"
)

func main() {
	db := dbwrapper.ConnectDB()
	defer db.Close()

	dbwrapper.CleanUP(db, campaign.SQLDropCampaignsTable)
	dbwrapper.Migrate(db, campaign.SQLCreateCampaignsTable)
}
