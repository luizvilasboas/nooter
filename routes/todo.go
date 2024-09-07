package routes

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"gitlab.com/olooeez/nooter/handlers"
)

func SetupRouter() *mux.Router {
	router := mux.NewRouter()

	router.Use(loggingMiddleware)
	router.Use(recoveryMiddleware)

	api := router.PathPrefix("/api/v1").Subrouter()
	api.HandleFunc("/todos", handlers.GetTodos).Methods(http.MethodGet)
	api.HandleFunc("/todos", handlers.CreateTodo).Methods(http.MethodPost)
	api.HandleFunc("/todos/{id:[0-9]+}", handlers.GetTodoByID).Methods(http.MethodGet)
	api.HandleFunc("/todos/{id:[0-9]+}", handlers.UpdateTodo).Methods(http.MethodPut)
	api.HandleFunc("/todos/{id:[0-9]+}", handlers.DeleteTodoByID).Methods(http.MethodDelete)

	return router
}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		log.Printf("Started %s %s", r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
		log.Printf("Completed in %v", time.Since(start))
	})
}

func recoveryMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				log.Printf("Recovered from panic: %v", err)
				http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			}
		}()
		next.ServeHTTP(w, r)
	})
}
