package routes

import (
	controller "ConnectApp/src/controllers"
	"net/http"

	"github.com/gorilla/mux"
)

func Routes() http.Handler {
	router := mux.NewRouter()
	// router.HandleFunc("/Vaccines/List", controller.Vaccines_List).Methods("GET")
	router.HandleFunc("/Appointments/List", controller.Appointments_List).Methods("GET")
	router.HandleFunc("/Doctors/List", controller.Doctors_List).Methods("GET")
	router.HandleFunc("/Patient/Doctors/Appointments", controller.Doctor_Wise_Patient_appointments).Methods("GET")
	router.HandleFunc("/Patients/List", controller.All_Patients).Methods("GET")
	router.HandleFunc("/Clinics/List", controller.Clinics_List).Methods("GET")
	router.HandleFunc("/Chat", controller.WsEndpoint).Methods("GET")
	return router
}
