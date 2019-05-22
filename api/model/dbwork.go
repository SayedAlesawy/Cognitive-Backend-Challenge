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

// ChartCampaigns A function to fetch the data needed for the chart
func ChartCampaigns(db *sql.DB, dimensions []string) []ChartResponse {
	var dim string
	for i := 0; i < len(dimensions); i++ {
		dim += dimensions[i]

		if i < len(dimensions)-1 {
			dim += ", "
		}
	}

	sqlStatement := fmt.Sprintf("SELECT %s, COUNT(ID) FROM campaigns GROUP BY %s", dim, dim)

	logMsgs := logger.LogInfo{
		Success: "Chart selected Successfully",
		Error:   "Chart selection failed",
	}

	rows, ok := dbwrapper.ExecuteRowsQuery(db, sqlStatement, logMsgs, false)
	defer rows.Close()

	var chart []ChartResponse
	for rows.Next() {
		var Key1 string
		var Key2 string
		var Value int

		if len(dimensions) == 2 {
			err := rows.Scan(&Key1, &Key2, &Value)
			logger.LogDBErr(err, dbwrapper.LogSign, "ChartCampaigns(): Error while extracting results", false)
		}

		if len(dimensions) == 1 {
			err := rows.Scan(&Key1, &Value)
			logger.LogDBErr(err, dbwrapper.LogSign, "ChartCampaigns(): Error while extracting results", false)
		}

		res := ChartResponse{
			Key1:  Key1,
			Key2:  Key2,
			Value: Value,
		}

		chart = append(chart, res)
	}

	err := rows.Err()
	logger.LogDBErr(err, dbwrapper.LogSign, "ChartCampaigns(): Error while extracting results", false)
	logger.LogDBSuccess(err, dbwrapper.LogSign, "Chart extracted successfully")

	if ok == false {
		chart = []ChartResponse{}
	}

	return chart
}
