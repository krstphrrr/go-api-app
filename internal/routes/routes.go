package routes

import (
	"database/sql"
	"net/http"
	
    "go-api-app/internal/middleware"
	"go-api-app/internal/services"
	"go-api-app/config"
)

func RegisterRoutes(mux *http.ServeMux, db *sql.DB) {
	for endpoint := range config.EndpointToTableMap {
		// Handling endpoints that don't require middleware 

		// if endpoint == "/tblProject" {
		// 	mux.Handle(endpoint, http.HandlerFunc(services.tblProjectHandler))

		// } else {
		// 	mux.Handle(endpoint, middleware.AuthMiddleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// 		services.GenericDynamicDataHandler(w, r, db)
		// 	})))
		// }
		// mux.Handle(endpoint, middleware.AuthMiddleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// 	services.GenericDynamicDataHandler(w, r, db)
		// })))
		mux.Handle(endpoint, middleware.AuthMiddleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Differentiate between GET and POST requests in the handler
			switch r.Method {
			case http.MethodGet:
				services.GenericDynamicDataHandler(w, r, db) // Handle GET
			case http.MethodPost:
				services.GenericDynamicDataPostHandler(w, r, db) // Handle POST
			default:
				http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
			}
		})))
	}
}