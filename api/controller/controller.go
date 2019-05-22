package main

import (
	campaign "Cognitive-Backend-Challenge/api/model"
	dbwrapper "Cognitive-Backend-Challenge/utils/database"
	logger "Cognitive-Backend-Challenge/utils/logger"
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// db A Database access resource
var db *sql.DB

// handleCreateCampaign A function to handle create requests
func handleCreateCampaign(writer http.ResponseWriter, req *http.Request) {
	var campaignObj campaign.Campaign

	json.NewDecoder(req.Body).Decode(&campaignObj)

	if campaignObj.Category == "" {
		category, ok := getCategory(campaignObj.Category)
		if ok == false {
			return
		}

		campaignObj.Category = category
	}

	campaign.CreateCampaign(db, campaignObj)
}

func getCategory(_url string) (string, bool) {
	url := fmt.Sprintf("https://ngkc0vhbrl.execute-api.eu-west-1.amazonaws.com/api/?url=%s", _url)

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("cache-control", "no-cache")
	req.Header.Add("Postman-Token", "0b4cc901-cf79-4ceb-900f-ec94d2452dcc")

	var category string
	var ok bool

	res, err := http.DefaultClient.Do(req)
	if err == nil {
		category = ""
		ok = false
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err == nil {
		category = ""
		ok = false
	}

	var apiResponse campaign.CategoryExtractionResponse
	err = json.Unmarshal(body, &apiResponse)
	if err == nil {
		category = ""
		ok = false
	}

	category = apiResponse.Category.Name
	ok = true

	return category, ok
}

// handleReadCampaign A function to handle read requests
func handleReadCampaign(writer http.ResponseWriter, req *http.Request) {
	params := req.URL.Query()
	id := params.Get("id")

	if id == "" {
		campaigns := campaign.ReadCampaigns(db)
		json.NewEncoder(writer).Encode(campaigns)
	} else {
		campaigns, _ := campaign.ReadCampaign(db, id)

		json.NewEncoder(writer).Encode(campaigns)
	}
}

// handleUpdateCampaign A function to handle update requests
func handleUpdateCampaign(writer http.ResponseWriter, req *http.Request) {
	var campaignObj campaign.Campaign

	json.NewDecoder(req.Body).Decode(&campaignObj)

	campaign.UpdateCampaign(db, campaignObj)
}

// handleDeleteCampaign A function to handle delete requests
func handleDeleteCampaign(writer http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	id := params["id"]

	campaign.DeleteCampaign(db, id)
}

func handleChartData(writer http.ResponseWriter, req *http.Request) {

}

func main() {
	//Establish connection to the DB
	db = dbwrapper.ConnectDB()
	defer db.Close()

	//Setup the DB
	//dbwrapper.CleanUP(db, campaign.SQLDropCampaignsTable)
	//dbwrapper.Migrate(db, campaign.SQLCreateCampaignsTable)

	//Acquire router
	router := mux.NewRouter()

	//Specify the different API end points
	router.HandleFunc("/new", handleCreateCampaign).Methods("POST")
	router.HandleFunc("/read", handleReadCampaign).Methods("GET")
	router.HandleFunc("/update", handleUpdateCampaign).Methods("PUT")
	router.HandleFunc("/delete/{id}", handleDeleteCampaign).Methods("DELETE")
	router.HandleFunc("/chart", handleChartData).Methods("GET")

	logger.LogMsg("[Controller]", 0, "Listening on port 3000")
	log.Fatal(http.ListenAndServe(":3000", router))
}
