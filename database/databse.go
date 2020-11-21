package database

import(
	"fmt"
	"database/sql"
	"os"
	_ "github.com/lib/pq"
)
var DB *sql.DB
func Connect()  {
	db, err := sql.Open("postgres",os.Getenv("DATABASE_URL"))
	if err != nil {
		panic("Failed to connect database")
	}
	fmt.Println("Database Conected!")
	DB = db
}