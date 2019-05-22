package campaign

import (
	dbwrapper "Cognitive-Backend-Challenge/utils/database"
	logger "Cognitive-Backend-Challenge/utils/logger"
	"database/sql"
	"fmt"
)

// CreateCampaign A function to create a campaign record
func CreateCampaign(db *sql.DB, campaign Campaign) bool {
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

// ReadCampaign A function to read a campaign record
func ReadCampaign(db *sql.DB, id string) (Campaign, bool) {
	sqlStatement := sqlReadCampaign

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

// ReadCampaigns A function to read all campaigns
func ReadCampaigns(db *sql.DB) []Campaign {
	sqlStatement := sqlReadCampaigns

	logMsgs := logger.LogInfo{
		Success: "Campaigns selected Successfully",
		Error:   "Campaigns selection failed",
	}

	rows, ok := dbwrapper.ExecuteRowsQuery(db, sqlStatement, logMsgs, false)
	defer rows.Close()

	var campaigns []Campaign
	for rows.Next() {
		var serialID int
		var ID string
		var Name string
		var Country string
		var Budget float32
		var Goal string
		var Category string
		var URL string

		err := rows.Scan(&serialID, &ID, &Name, &Country, &Budget, &Goal, &Category, &URL)
		logger.LogDBErr(err, dbwrapper.LogSign, "ReadCampaigns(): Error while extracting results", false)

		res := Campaign{
			ID:       ID,
			Name:     Name,
			Country:  Country,
			Budget:   Budget,
			Goal:     Goal,
			Category: Category,
			URL:      URL,
		}

		campaigns = append(campaigns, res)
	}

	err := rows.Err()
	logger.LogDBErr(err, dbwrapper.LogSign, "ReadCampaigns(): Error while extracting results", false)
	logger.LogDBSuccess(err, dbwrapper.LogSign, "Campaigns extracted successfully")

	if ok == false {
		campaigns = []Campaign{}
	}

	return campaigns
}

// UpdateCampaign A function to update a campaign record
func UpdateCampaign(db *sql.DB, campaign Campaign) bool {
	sqlStatement := sqlUpdateCampaign

	logMsgs := logger.LogInfo{
		Success: "Campaign updated Successfully",
		Error:   "Campaign update failed",
	}

	ok := dbwrapper.ExecuteQuery(db, sqlStatement, logMsgs, false,
		campaign.Name, campaign.Country, campaign.Budget, campaign.Goal,
		campaign.Category, campaign.URL, campaign.ID)

	return ok
}

// DeleteCampaign A function to delete a campaign record
func DeleteCampaign(db *sql.DB, id string) bool {
	sqlStatement := sqlDeleteCampaign

	logMsgs := logger.LogInfo{
		Success: fmt.Sprintf("Campaign #%s deleted Successfully", id),
		Error:   fmt.Sprintf("Failed to delete Campaign #%s", id),
	}

	ok := dbwrapper.ExecuteQuery(db, sqlStatement, logMsgs, false, id)

	return ok
}
