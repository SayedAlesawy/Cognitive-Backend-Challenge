package campaign

// Campaign Represents the campaign resource
type Campaign struct {
	ID       string  //Unique identifier for the Campaign record
	Name     string  //The Campaign name
	Country  string  //The country at which the Campaign orginates
	Budget   float32 //The budget allocated to the Campaign
	Goal     string  //The market goal of the Campaign
	Category string  //Fetched from the category extraction service if not provided
	URL      string  //Used to fetched the category from the extraction service if not provided
}
