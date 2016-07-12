package api

import (
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
)

const apiUrl = "/api"
const apiVersion = "/v1"
const baseUrl = apiUrl + apiVersion

func Router() *httprouter.Router {
	router := httprouter.New()

	router.GET(baseUrl+"/user/", userGET)
	router.GET(baseUrl+"/hours/", hoursGET)

	return router
}

func userGET(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	log.Printf("GET\t%v", r.URL)
	response := MockUserResponse()

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(response); err != nil {
		log.Panic(err)
	}
}

func hoursGET(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	log.Printf("GET\t%v", r.URL)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	queryValues := r.URL.Query()
	response, err := MockHoursResponse(queryValues.Get("start-date"), queryValues.Get("end-date"))

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		errorResponse := ErrorResponse{
			Status:     http.StatusInternalServerError,
			StatusText: err.Error(),
		}
		if err := json.NewEncoder(w).Encode(errorResponse); err != nil {
			log.Panic(err)
		}
		return
	}

	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(response); err != nil {
		log.Panic(err)
	}
}
