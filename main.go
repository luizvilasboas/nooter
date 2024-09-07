package main

import (
	"fmt"
	"log"
	"net/http"

	_ "github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"
	_ "gitlab.com/olooeez/nooter/docs"
	"gitlab.com/olooeez/nooter/routes"
	"gitlab.com/olooeez/nooter/storage"
)

// @title Todo API
// @version 1.0
// @description API para gerenciar tarefas (todos) com Go e SQLite.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /api/v1
func main() {
	storage.InitDB("todos.db")

	r := routes.SetupRouter()

	r.PathPrefix("/docs").Handler(httpSwagger.WrapHandler)

	fmt.Println("Server is running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
