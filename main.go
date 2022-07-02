package main

import (
	"database/sql"
	"fmt"
	"log"
	"main/models"

	_ "github.com/lib/pq"
)

//lint:ignore U1000 example
func createUserTable(db *sql.DB) {
	q := `CREATE TABLE IF NOT EXISTS users(
  			id bigserial primary key,
  			name text not null,
  			email text not null,
  			created_at timestamp(0) with time zone not null default now()
		);`

	_, err := db.Exec(q)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(`Created table "users"`)
}

//lint:ignore U1000 example
func createUser(db *sql.DB, name string, email string) {
	um := models.UsersModel{DB: db}
	user := models.User{Name: name, Email: email}

	if err := um.Insert(&user); err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("Created user with ID %d\n", user.ID)
}

//lint:ignore U1000 example
func getAllUsers(db *sql.DB) {
	um := models.UsersModel{DB: db}
	f := models.Filter{
		PageSize: 5,
		Page:     3,
	}
	users, _, err := um.GetAll(f)
	if err != nil {
		log.Fatalln(err)
	}
	for _, user := range users {
		fmt.Printf("%d: %s %s\n", user.ID, user.Name, user.Email)
	}

}

type Application struct {
	Models models.Models
}

func main() {
	dsn := "postgres://docker:docker@localhost:5432/go_sql?sslmode=disable"

	db, err := connectToDb(dsn)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("Connected to db")

	app := Application{
		Models: models.NewModel(db),
	}

	fmt.Println("Starting application...")
	if err = app.serve(); err != nil {
		log.Fatalln(err)
	}
}

func connectToDb(dsn string) (*sql.DB, error) {
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
