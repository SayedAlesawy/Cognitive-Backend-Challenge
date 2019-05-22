package campaign

import (
	dbwrapper "Cognitive-Backend-Challenge/utils/database"
	logger "Cognitive-Backend-Challenge/utils/logger"
	"database/sql"
	"fmt"
)

// createCampaign A function to create a campaign record
func createCampaign(db *sql.DB, campaign Campaign) bool {
	sqlStatement := sqlCreateCampaign

	logMsgs := logger.LogInfo{
		Success: fmt.Sprintf("Campaign #%s created Successfully", campaign.ID),
		Error:   fmt.Sprintf("Failed to create Campaign #%s", campaign.ID),
	}

	ok := dbwrapper.ExecuteQuery(db, sqlStatement, logMsgs, false,
		campaign.ID, campaign.Name, campaign.Country, campaign.Budget, campaign.Goal,
		campaign.Category, campaign.URL)

	return ok
}

// readCampaign A function to read a campaign record
func readCampaign(db *sql.DB, id string) (Campaign, bool) {
	sqlStatement := sqlReadCampaing

	row := dbwrapper.ExecuteRowQuery(db, sqlStatement, id)

	var serialID int
	var res Campaign

	err := row.Scan(&serialID, &res.ID, &res.Name, &res.Country, &res.Budget,
		&res.Goal, &res.Category, &res.URL)

	if err == sql.ErrNoRows {
		return Campaign{}, false
	}

	return res, true
}

// updateCampaign A function to update a campaign record
func updateCampaign(db *sql.DB, campaign Campaign) bool {
	sqlStatement := sqlUpdateCampaign

	logMsgs := logger.LogInfo{
		Success: "Campaign updated Successfully",
		Error:   "Campaign update failed",
	}

	ok := dbwrapper.ExecuteQuery(db, sqlStatement, logMsgs, false,
		campaign.Name, campaign.Country, campaign.Budget, campaign.Goal,
		campaign.Category, campaign.URL)

	return ok
}

// deleteCampaign A function to delete a campaign record
func deleteCampaign(db *sql.DB, id string) bool {
	sqlStatement := sqlDeleteCampaign

	logMsgs := logger.LogInfo{
		Success: fmt.Sprintf("Campaign #%s deleted Successfully", id),
		Error:   fmt.Sprintf("Failed to delete Campaign #%s", id),
	}

	ok := dbwrapper.ExecuteQuery(db, sqlStatement, logMsgs, false, id)

	return ok
}
