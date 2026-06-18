package main

import (
	"encoding/json"
	"net/http"
	"strings"
)

var resiValidator ResiValidator = RealResiValidator{}
var trackingRepo TrackingRepository

func insertTrackingEventHandler(w http.ResponseWriter, r *http.Request) {
	var req TrackingEvent
	json.NewDecoder(r.Body).Decode(&req)

	result, err := InsertTrackingEvent(
		req,
		resiValidator,
		trackingRepo,
	)

	if err != nil {
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(map[string]string{
			"error": err.Error(),
		})
		return
	}

	json.NewEncoder(w).Encode(result)
}

func getTrackingHandler(w http.ResponseWriter, r *http.Request) {
	resi := r.URL.Query().Get("resi")
	if resi == "" {
		resi = strings.TrimPrefix(r.URL.Path, "/tracking/")
	}

	resp := GetTrackingStatus(resi)
	if resp == nil {
		w.WriteHeader(404)
		return
	}

	json.NewEncoder(w).Encode(resp)
}

func calculateDistanceHandler(w http.ResponseWriter, r *http.Request) {
	var req DistanceRequest
	json.NewDecoder(r.Body).Decode(&req)

	resp := CalculateDistance(req)
	if resp == nil {
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "Failed to calculate distance",
		})
		return
	}

	json.NewEncoder(w).Encode(resp)
}

func calculateRouteHandler(w http.ResponseWriter, r *http.Request) {
	var req RouteRequest
	json.NewDecoder(r.Body).Decode(&req)

	resp := CalculateRoute(req)
	if resp == nil {
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "Failed to calculate route",
		})
		return
	}

	json.NewEncoder(w).Encode(resp)
}

func getCourierLocationHandler(w http.ResponseWriter, r *http.Request) {
	courierID := r.URL.Query().Get("courier_id")
	if courierID == "" {
		courierID = strings.TrimPrefix(r.URL.Path, "/location/")
	}

	resp := GetCourierLocation(courierID)
	if resp == nil {
		w.WriteHeader(404)
		return
	}

	json.NewEncoder(w).Encode(resp)
}
