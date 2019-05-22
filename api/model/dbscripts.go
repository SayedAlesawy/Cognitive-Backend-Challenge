package campaign

// SQLCreateCampaignsTable SQL to create the campaigns table
const SQLCreateCampaignsTable string = `
	CREATE TABLE campaigns (
		id SERIAL PRIMARY KEY,
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
