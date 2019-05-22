package main

import (
	campaign "Cognitive-Backend-Challenge/api/model"
	dbwrapper "Cognitive-Backend-Challenge/utils/database"
	logger "Cognitive-Backend-Challenge/utils/logger"
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var db *sql.DB

func handleCreateCampaign(writer http.ResponseWriter, req *http.Request) {
	var campaignObj campaign.Campaign

	json.NewDecoder(req.Body).Decode(&campaignObj)

	campaign.CreateCampaign(db, campaignObj)
}

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

func handleUpdateCampaign(writer http.ResponseWriter, req *http.Request) {
	var campaignObj campaign.Campaign

	json.NewDecoder(req.Body).Decode(&campaignObj)

	campaign.UpdateCampaign(db, campaignObj)
}

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
