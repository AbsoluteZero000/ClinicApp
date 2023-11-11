package main

import (
	"awesomeProject/Controller"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func roleBasedAuthMiddleware(allowedRoles map[string][]string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			session, _ := Controller.Store.Get(r, "Session")

			authenticated, ok := session.Values["authenticated"].(bool)

			fmt.Println(r.Header.Values("Authorization"))
			fmt.Println(authenticated, ok)
			if !ok || !authenticated {
				w.Header().Set("WWW-Authenticate", `Basic realm="Restricted"`)

				http.Error(w, "Unauthorized", http.StatusUnauthorized)
				return
			}

			userRole, _ := session.Values["role"].(string)
			if allowedPaths, ok := allowedRoles[userRole]; ok {
				path := r.URL.Path
				allowed := false
				for _, allowedPath := range allowedPaths {
					if allowedPath == path {
						allowed = true
						break
					}
				}
				if !allowed {
					http.Error(w, "Forbidden", http.StatusForbidden)
					return
				}
			}

			next.ServeHTTP(w, r)
		})
	}
}

func main() {
	allowedRoles := map[string][]string{
		"patient": {"/editpatientslot", "/deleteslot", "/getpatientslots", "/addpatientslot", "/updateuser", "/getusers", "/deleteuser"},
		"doctor":  {"/editdoctorslot", "/deleteslot", "/addslot", "/getdoctorslots", "/updateuser", "/getusers", "/deleteuser"},
		"admin":   {"/editpatientslot", "/getpatientslots", "/addpatientslot", "/editdoctorslot", "/deleteslot", "/addslot", "/getdoctorslots", "/updateuser", "/getusers", "/deleteuser"},
	}

	Controller.InitialData()

	router := mux.NewRouter()

	router.Handle("/signup", http.HandlerFunc(Controller.SignUp)).Methods("POST")
	router.Handle("/login", http.HandlerFunc(Controller.Login)).Methods("POST")
	router.Handle("/signout", http.HandlerFunc(Controller.Signout)).Methods("GET")

	router.Handle("/editdoctorslot", roleBasedAuthMiddleware(allowedRoles)(http.HandlerFunc(Controller.GetPatientSlots))).Methods("PUT")
	router.Handle("/getdoctorslots", roleBasedAuthMiddleware(allowedRoles)(http.HandlerFunc(Controller.GetDoctorSlots))).Methods("GET")
	router.Handle("/addpatientslot", roleBasedAuthMiddleware(allowedRoles)(http.HandlerFunc(Controller.AddPatientSlot))).Methods("POST")
	router.Handle("/getpatientslots", roleBasedAuthMiddleware(allowedRoles)(http.HandlerFunc(Controller.GetPatientSlots))).Methods("GET")
	router.Handle("/addslot", roleBasedAuthMiddleware(allowedRoles)(http.HandlerFunc(Controller.AddSlot))).Methods("POST")
	router.Handle("/deleteslot", roleBasedAuthMiddleware(allowedRoles)(http.HandlerFunc(Controller.DeleteSlot))).Methods("DELETE")
	router.Handle("/editpatientslot", roleBasedAuthMiddleware(allowedRoles)(http.HandlerFunc(Controller.EditPatientSlot))).Methods("PUT")

	router.Handle("/getusers", roleBasedAuthMiddleware(allowedRoles)(http.HandlerFunc(Controller.GetAllUsers))).Methods("GET")
	router.Handle("/updateuser", roleBasedAuthMiddleware(allowedRoles)(http.HandlerFunc(Controller.UpdateUser))).Methods("PUT")
	router.Handle("/deleteuser", roleBasedAuthMiddleware(allowedRoles)(http.HandlerFunc(Controller.DeleteUser))).Methods("DELETE")

	headersOk := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
	originsOk := handlers.AllowedOrigins([]string{"http://localhost:4200"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS"})

	// Wrap the router with the CORS middleware
	log.Fatal(http.ListenAndServe(":8080", handlers.CORS(originsOk, headersOk, methodsOk)(router)))

	http.Handle("/", router)
	fmt.Println("Connected to port 8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
