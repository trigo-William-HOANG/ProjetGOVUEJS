package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/schema"
	log "github.com/sirupsen/logrus"
	"github.com/trigo-William-HOANG/ProjetGOVUEJS/api"
	"github.com/trigo-William-HOANG/ProjetGOVUEJS/internal/tools"
)

func GetAppDataHandler(w http.ResponseWriter, r *http.Request) {
	var params = api.AppDataParams{}
	var decoder *schema.Decoder = schema.NewDecoder()
	var err error
	err = decoder.Decode(&params, r.URL.Query())
	if err != nil {
		log.Error(err)
		api.InternalErrorHandler(w)
		return
	}

	var database *tools.DatabaseInterface
	database, err = tools.NewDatabase()
	if err != nil {
		log.Error(err)
		api.InternalErrorHandler(w)
		return
	}

	var appData *tools.AppDetails = (*database).GetAppData(params.Id)
	if appData == nil {
		api.InternalErrorHandler(w)
		return
	}

	var response = api.AppDataResponse{
		Code:        http.StatusOK,
		Title:       (*appData).Title,
		Description: (*appData).Description,
		Link:        (*appData).Link,
		Image:       (*appData).Image,
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		log.Error(err)
		api.InternalErrorHandler(w)
		return
	}

}
