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

	router.GET(baseUrl+"/hours/", hoursGET)
	router.POST(baseUrl+"/hours/", hoursPOST)
	router.PUT(baseUrl+"/hours/", hoursPOST)
	router.DELETE(baseUrl+"/hours/", hoursPOST)

	//router.GET(baseUrl+"/projects/", projects)
	//router.GET(baseUrl+"/holidays/", holidays)
	//router.GET(baseUrl+"/users/", users)

	return router
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

//func projects(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
//	fmt.Fprint(w, "Welcome!\n")
//}
//
//func holidays(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
//	fmt.Fprint(w, "Welcome!\n")
//}
//
//func users(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
//	fmt.Fprint(w, "Welcome!\n")
//}

//func Hello(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
//	fmt.Fprintf(w, "hello, %s!\n", ps.ByName("name"))
//}
