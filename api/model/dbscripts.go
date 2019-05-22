package campaign

// SQLCreateCampaignsTable SQL to create the campaigns table
const SQLCreateCampaignsTable string = `
	CREATE TABLE campaigns (
		id SERIAL PRIMARY KEY,
		ID varchar(60) UNIQUE NOT NULL,
		Name varchar(60) NOT NULL,
		Country varchar(60) NOT NULL,
		Budget float NOT NULL, 
		Goal varchar(200) NOT NULL,
		Category varchar(60) NOT NULL,
		URL varchar(200)
	);
`

// SQLDropCampaignsTable SQL to drop the campaings table
const SQLDropCampaignsTable string = `
	DROP TABLE IF EXISTS campaigns;
`

// Defining all the CRUD operations from the campaign resource
// CRUD => Create - Read - Update - Delete

// SQLCreateCampaign SQL to create a new campaign record
const SQLCreateCampaign string = `
	INSERT INTO campaigns (ID, Name, Country, Budget, Goal, Category, URL)
	VALUES ($1, $2, $3, $4, $5, $6, $7)
`

// SQLReadCampaing SQL to read a campaign record
const SQLReadCampaing string = `
	SELECT * FROM campaigns
	WHERE ID = $1
`

// SQLUpdateCampaign SQL to update a campaign record
const SQLUpdateCampaign string = `
	UPDATE campaigns
	SET Name = $1, Country = $2, Budget = $3, Goal = $4, Category = $5, URL = $6
	WHERE ID = $7
`

// SQLDeleteCampaign SQL to delete a campaign record
const SQLDeleteCampaign string = `
	DELETE FROM campaigns
	where ID = $1
`
