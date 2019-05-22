package campaign

// Campaign Represents the campaign resource
type Campaign struct {
	ID       string  `json:"id,omitempty"`       //Unique identifier for the Campaign record
	Name     string  `json:"name,omitempty"`     //The Campaign name
	Country  string  `json:"country,omitempty"`  //The country at which the Campaign orginates
	Budget   float32 `json:"budget,omitempty"`   //The budget allocated to the Campaign
	Goal     string  `json:"goal,omitempty"`     //The market goal of the Campaign
	Category string  `json:"category,omitempty"` //Fetched from the category extraction service if not provided
	URL      string  `json:"url,omitempty"`      //Used to fetched the category from the extraction service if not provided
}
