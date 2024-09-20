package main

import (
	"database/sql"
	"fmt"
	"log"

	"gitlab.com/olooeez/nooter/controllers"
	"gitlab.com/olooeez/nooter/middlewares"
	"gitlab.com/olooeez/nooter/routes"

	"github.com/gin-gonic/gin"
	"github.com/golang-migrate/migrate"
	"github.com/golang-migrate/migrate/database/sqlite3"
	_ "github.com/golang-migrate/migrate/source/file"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "gitlab.com/olooeez/nooter/docs"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// @title Nooter API
// @version 0.1.0
// @description Esta é uma API para gerenciar notas.

// @contact.name Luiz Felipe de Castro Vilas Boas
// @contact.url https://olooeez.gitlab.io
// @contact.email luizfelipecastrovb@gmail.com

// @license.name MIT
// @license.url https://gitlab.com/olooeez/nooter/-/blob/main/LICENSE
// @host localhost:8080
// @BasePath /api/v1
func main() {
	db, err := gorm.Open(sqlite.Open("nooter.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("Falha ao conectar ao banco de dados: ", err)
	}

	controllers.DB = db

	initDB("nooter.db")

	r := gin.Default()

	r.Use(middlewares.LoggerMiddleware())

	routes.NoteRoutes(r)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	r.Run(":8080")
}

func initDB(filepath string) {
	db, err := sql.Open("sqlite3", filepath)
	if err != nil {
		log.Fatalf("Failed to open database: %v", err)
	}

	if err := db.Ping(); err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	fmt.Println("Database connection established")

	runMigrations(db)
}

func runMigrations(db *sql.DB) {
	if db == nil {
		log.Fatal("Database is not initialized")
	}

	driver, err := sqlite3.WithInstance(db, &sqlite3.Config{})
	if err != nil {
		log.Fatalf("Could not start SQL driver: %v", err)
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://migrations",
		"sqlite3", driver)

	if err != nil {
		log.Fatalf("Migration setup failed: %v", err)
	}

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatalf("Could not run up migrations: %v", err)
	}

	fmt.Println("Migrations ran successfully")
}
