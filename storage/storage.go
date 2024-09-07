package storage

import (
	"database/sql"
	"fmt"
	"log"
	"reflect"
	"strings"

	"github.com/golang-migrate/migrate"
	"github.com/golang-migrate/migrate/database/sqlite3"
	_ "github.com/golang-migrate/migrate/source/file"
	_ "github.com/mattn/go-sqlite3"
)

type Storage interface {
	InitDB(filepath string)
	Create(table string, data interface{}) (int, error)
	Read(table string, conditions string, dest interface{}) error
	Update(table string, data interface{}, conditions string) error
	Delete(table string, conditions string) error
}

type DefaultStorage struct {
	db *sql.DB
}

func (ds *DefaultStorage) InitDB(filepath string) {
	var err error

	ds.db, err = sql.Open("sqlite3", filepath)
	if err != nil {
		log.Fatalf("Failed to open database: %v", err)
	}

	if err := ds.db.Ping(); err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	fmt.Println("Database connection established")

	ds.runMigrations()
}

func (ds *DefaultStorage) runMigrations() {
	if ds.db == nil {
		log.Fatal("Database is not initialized")
	}

	driver, err := sqlite3.WithInstance(ds.db, &sqlite3.Config{})
	if err != nil {
		log.Fatalf("Could not start SQL driver: %v", err)
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://db/migrations",
		"sqlite3", driver)

	if err != nil {
		log.Fatalf("Migration setup failed: %v", err)
	}

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatalf("Could not run up migrations: %v", err)
	}

	fmt.Println("Migrations ran successfully")
}

func (ds *DefaultStorage) Create(table string, data interface{}) (int, error) {
	query, values := buildInsertQuery(table, data)
	result, err := ds.db.Exec(query, values...)
	if err != nil {
		return 0, fmt.Errorf("failed to insert into %s: %v", table, err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("failed to get last insert id: %v", err)
	}

	return int(id), nil
}

func (ds *DefaultStorage) Read(table string, conditions string, dest interface{}) error {
	query := fmt.Sprintf("SELECT * FROM %s WHERE %s", table, conditions)
	rows, err := ds.db.Query(query)
	if err != nil {
		return fmt.Errorf("failed to query %s: %v", table, err)
	}
	defer rows.Close()

	if err := scanRows(rows, dest); err != nil {
		return fmt.Errorf("failed to scan rows for %s: %v", table, err)
	}

	return nil
}

func (ds *DefaultStorage) Update(table string, data interface{}, conditions string) error {
	query, values := buildUpdateQuery(table, data, conditions)
	_, err := ds.db.Exec(query, values...)
	if err != nil {
		return fmt.Errorf("failed to update %s: %v", table, err)
	}

	return nil
}

func (ds *DefaultStorage) Delete(table string, conditions string) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE %s", table, conditions)
	_, err := ds.db.Exec(query)
	if err != nil {
		return fmt.Errorf("failed to delete from %s: %v", table, err)
	}

	return nil
}

func buildInsertQuery(table string, data interface{}) (string, []interface{}) {
	v := reflect.ValueOf(data).Elem()
	typeOfS := v.Type()

	var columns []string
	var placeholders []string
	var values []interface{}

	for i := 0; i < v.NumField(); i++ {
		field := typeOfS.Field(i)
		columns = append(columns, field.Tag.Get("db"))
		placeholders = append(placeholders, "?")
		values = append(values, v.Field(i).Interface())
	}

	query := fmt.Sprintf("INSERT INTO %s (%s) VALUES (%s)", table, strings.Join(columns, ", "), strings.Join(placeholders, ", "))
	return query, values
}

func buildUpdateQuery(table string, data interface{}, conditions string) (string, []interface{}) {
	v := reflect.ValueOf(data).Elem()
	typeOfS := v.Type()

	var setClauses []string
	var values []interface{}

	for i := 0; i < v.NumField(); i++ {
		field := typeOfS.Field(i)
		columnName := field.Tag.Get("db")
		setClauses = append(setClauses, fmt.Sprintf("%s = ?", columnName))
		values = append(values, v.Field(i).Interface())
	}

	query := fmt.Sprintf("UPDATE %s SET %s WHERE %s", table, strings.Join(setClauses, ", "), conditions)
	return query, values
}

func scanRows(rows *sql.Rows, dest interface{}) error {
	slice := reflect.ValueOf(dest).Elem()
	elemType := slice.Type().Elem()

	for rows.Next() {
		elem := reflect.New(elemType).Elem()
		fields := make([]interface{}, elem.NumField())
		for i := 0; i < elem.NumField(); i++ {
			fields[i] = elem.Field(i).Addr().Interface()
		}
		if err := rows.Scan(fields...); err != nil {
			return err
		}
		slice.Set(reflect.Append(slice, elem))
	}

	return rows.Err()
}
