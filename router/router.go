package router

import (
	"studentregist/controllers"

	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/api/students", controllers.GetAllStudent).Methods("GET")
	router.HandleFunc("/api/students/{id}", controllers.GetSingleStudent).Methods("GET")
	router.HandleFunc("/api/students", controllers.CreateData).Methods("POST")
	router.HandleFunc("/api/students/{id}", controllers.CreateMultipleData).Methods("POST")
	router.HandleFunc("/api/students/{id}", controllers.UpdateData).Methods("PUT")
	router.HandleFunc("/api/students/{id}", controllers.DeleteData).Methods("DELETE")
	router.HandleFunc("/api/students", controllers.DeleteAllData).Methods("DELETE")

	return router

}
