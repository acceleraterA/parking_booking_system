package router

import (
	"github.com/acceleraterA/parking_booking_system/server/middleware"

	"github.com/gorilla/mux"
)

// Router is exported and used in main.go
func Router() *mux.Router {

	router := mux.NewRouter()

	//router.HandleFunc("/api/spot", middleware.GetAllSpot).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/spot", middleware.CreateSpot).Methods("POST", "OPTIONS")
	//router.HandleFunc("/api/spot/{id}", middleware.SpotComplete).Methods("PUT", "OPTIONS")
	//router.HandleFunc("/api/undoSpot/{id}", middleware.UndoSpot).Methods("PUT", "OPTIONS")
	router.HandleFunc("/api/deleteSpot/{id}", middleware.DeleteSpot).Methods("DELETE", "OPTIONS")
	//router.HandleFunc("/api/deleteAllSpot", middleware.DeleteAllSpot).Methods("DELETE", "OPTIONS")
	return router
}
