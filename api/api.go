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
	router.POST(baseUrl+"/hours/", hoursPOST)
	router.PUT(baseUrl+"/hours/:id", hoursPUT)
	router.DELETE(baseUrl+"/hours/:id", hoursDelete)

	return router
}

func userGET(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	log.Printf("%v\t%v", r.Method, r.URL)
	response := MockUserResponse()
	sendOK(response, w)
}

func hoursGET(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	log.Printf("%v\t%v", r.Method, r.URL)
	queryValues := r.URL.Query()
	response, err := MockHoursResponse(queryValues.Get("start-date"), queryValues.Get("end-date"))

	if err != nil {
		sendError(err, w)
		return
	}
	sendOK(response, w)
}

func hoursPOST(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	log.Printf("%v\t%v", r.Method, r.URL)

	var body HoursUpdateRequest
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&body)
	if err != nil {
		sendError(err, w)
		return
	}
	response, err := MockHoursPOSTResponse(body)
	if err != nil {
		sendError(err, w)
		return
	}
	sendOK(response, w)
}

func hoursPUT(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	log.Printf("%v\t%v", r.Method, r.URL)

	var body HoursUpdateRequest
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&body)
	if err != nil {
		sendError(err, w)
		return
	}
	response, err := MockHoursPUTResponse(p.ByName("id"), body)
	if err != nil {
		sendError(err, w)
		return
	}
	sendOK(response, w)
}

func hoursDelete(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	log.Printf("%v\t%v", r.Method, r.URL)
	response, err := MockHoursDeleteResponse()
	if err != nil {
		sendError(err, w)
		return
	}
	sendOK(response, w)
}

func sendOK(response interface{}, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(response); err != nil {
		log.Panic(err)
	}
}

func sendError(err error, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	log.Println(err)
	w.WriteHeader(http.StatusInternalServerError)
	errorResponse := ErrorResponse{
		Status:     http.StatusInternalServerError,
		StatusText: err.Error(),
	}
	if err := json.NewEncoder(w).Encode(errorResponse); err != nil {
		log.Panic(err)
	}
}
