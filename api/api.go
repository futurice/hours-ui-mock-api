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
	queryValues := r.URL.Query()
	response, _ := MockHoursResponse(queryValues.Get("day__lte"), queryValues.Get("day__gte"))

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(response); err != nil {
		log.Panic(err)
	}
}

type test struct {
	ID          int     `json:"id"`
	Description string  `json:"description"`
	Day         string  `json:"day"`
	Hours       float64 `json:"hours"`
}

func hoursPOST(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	log.Println(r.Method)
	decoder := json.NewDecoder(r.Body)
	var t test
	err := decoder.Decode(&t)
	if err != nil {
		log.Panic(err)
	}
	log.Println(t)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(t); err != nil {
		log.Panic(err)
	}
}
