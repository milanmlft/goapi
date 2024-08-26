package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/schema"
	"github.com/milanmlft/goapi/api"
	"github.com/milanmlft/goapi/internal/tools"
	log "github.com/sirupsen/logrus"
)

func GetFlavours(w http.ResponseWriter, r *http.Request) {
	params := api.IceCreamParams{}
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
		api.InternalErrorHandler(w)
		return
	}

	var flavours *tools.Flavours
	flavours = (*database).GetFlavours()
	if flavours == nil {
		log.Error(err)
		api.InternalErrorHandler(w)
		return
	}

	var flavourNames []string
	for i, details := range *flavours {
		flavourNames[i] = details.Name
	}

	response := api.IceCreamResponse{
		Flavours: flavourNames,
		Code:     http.StatusOK,
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		log.Error(err)
		api.InternalErrorHandler(w)
		return
	}
}
