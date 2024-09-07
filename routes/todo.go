package routes

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"gitlab.com/olooeez/nooter/handlers"
	"gitlab.com/olooeez/nooter/storage"
)

func SetupRouter(storage storage.Storage) *mux.Router {
	router := mux.NewRouter()

	router.Use(loggingMiddleware)
	router.Use(recoveryMiddleware)

	h := &handlers.Handlers{
		Storage: storage,
	}

	api := router.PathPrefix("/api/v1").Subrouter()
	api.HandleFunc("/todos", h.GetTodos).Methods(http.MethodGet)
	api.HandleFunc("/todos", h.CreateTodo).Methods(http.MethodPost)
	api.HandleFunc("/todos/{id:[0-9]+}", h.GetTodoByID).Methods(http.MethodGet)
	api.HandleFunc("/todos/{id:[0-9]+}", h.UpdateTodo).Methods(http.MethodPut)
	api.HandleFunc("/todos/{id:[0-9]+}", h.DeleteTodoByID).Methods(http.MethodDelete)

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
