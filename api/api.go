package api

import (
	"encoding/json"
	"errors"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
	"time"
)

const apiUrl = "/api"
const apiVersion = "/v1"
const baseUrl = apiUrl + apiVersion

func Router() *httprouter.Router {
	router := httprouter.New()

	router.GET(baseUrl+"/user/", userGET)

	router.GET(baseUrl+"/hours/", hoursGET)

	router.POST(baseUrl+"/entry/", entryPOST)
	router.PUT(baseUrl+"/entry/:id", entryPUT)
	router.DELETE(baseUrl+"/entry/:id", entryDELETE)

	return router
}

func userGET(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	log.Printf("%v\t%v", r.Method, r.URL)
	response := MockUserResponse()
	if RandomFail() {
		sendError(errors.New("Could not load user data"), w)
	} else {
		sendOK(response, w)
	}
}

func hoursGET(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	log.Printf("%v\t%v", r.Method, r.URL)
	queryValues := r.URL.Query()
	response, err := MockHoursResponse(queryValues.Get("start-date"), queryValues.Get("end-date"))

	if err != nil {
		sendError(err, w)
		return
	}

	if RandomFail() {
		sendError(errors.New("Could not load hours"), w)
	} else {
		sendOK(response, w)
	}
}

func entryPOST(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	log.Printf("%v\t%v", r.Method, r.URL)

	var body EntryUpdateRequest
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&body)
	if err != nil {
		sendError(err, w)
		return
	}
	response, err := MockEntryPOSTResponse(body)
	if err != nil {
		sendError(err, w)
		return
	}
	if RandomFail() {
		sendError(errors.New("Could create new entry"), w)
	} else {
		sendOK(response, w)
	}
}

func entryPUT(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	log.Printf("%v\t%v", r.Method, r.URL)

	var body EntryUpdateRequest
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&body)
	if err != nil {
		sendError(err, w)
		return
	}
	response, err := MockEntryPUTResponse(p.ByName("id"), body)
	if err != nil {
		sendError(err, w)
		return
	}
	if RandomFail() {
		sendError(errors.New("Could not update entry"), w)
	} else {
		sendOK(response, w)
	}
}

func entryDELETE(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	log.Printf("%v\t%v", r.Method, r.URL)
	response, err := MockEntryDELETEResponse()
	if err != nil {
		sendError(err, w)
		return
	}
	if RandomFail() {
		sendError(errors.New("Could not delete entry"), w)
	} else {
		sendOK(response, w)
	}
}

func sendOK(response interface{}, w http.ResponseWriter) {
	duration := int(RandomFloat64(1000, 3000))
	time.Sleep(time.Millisecond * time.Duration(duration))

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(response); err != nil {
		log.Panic(err)
	}
}

func sendError(err error, w http.ResponseWriter) {
	duration := int(RandomFloat64(1000, 3000))
	time.Sleep(time.Millisecond * time.Duration(duration))

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	log.Println(err)
	w.WriteHeader(http.StatusInternalServerError)
	errorResponse := ErrorResponse{
		Status: http.StatusInternalServerError,
		Error:  err.Error(),
	}
	if err := json.NewEncoder(w).Encode(errorResponse); err != nil {
		log.Panic(err)
	}
}
