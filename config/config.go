package config

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"golang.org/x/crypto/bcrypt"
)

// Init initializes the environment variables
func init() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Couldn't load local variables, err:", err)
	}
}

// ConnectToDb connects to the database
func ConnectToDb() (*sql.DB, error) {
	dbSource := os.Getenv("MYSQL_SOURCE")
	db, err := sql.Open("mysql", dbSource)
	if err != nil {
		return nil, err
	}
	return db, nil
}

// TablesInit initializes the tables
func TablesInit() error {
	db, err := ConnectToDb()
	if err != nil {
		return err
	}
	defer db.Close()
	// Create tables if they don't exist
	_, err = db.Exec("CREATE TABLE IF NOT EXISTS users (id INT AUTO_INCREMENT PRIMARY KEY, username VARCHAR(50), password VARCHAR(50))")
	if err != nil {
		return err
	}

	_, err = db.Exec("CREATE TABLE IF NOT EXISTS users_logins (id INT AUTO_INCREMENT PRIMARY KEY,username VARCHAR(50),password VARCHAR(120))")
	if err != nil {
		return err
	}

	_, err = db.Exec("CREATE TABLE IF NOT EXISTS images (id INT AUTO_INCREMENT PRIMARY KEY, user_id INT,image_path VARCHAR(120), image_url VARCHAR(120))")
	if err != nil {
		return err
	}
	return nil
}

// CreateMockUser creates a mock user for login
func CreateMockUser() {
	db, err := ConnectToDb()
	if err != nil {
		log.Println(err)
	}
	defer db.Close()
	_, err = db.Exec("INSERT INTO users (username, password) VALUES ('test', '123456')")
	if err != nil {
		log.Println(err)
	}
}

// HashPassword hashes the user's password
func hashPassword(password string) ([]byte, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		log.Println("Couldn't hash user's password", err)
		return []byte(""), err
	}
	return hashedPassword, nil
}

// CreateMockLogs creates a mock log for login
func CreateMockLogs() {
	db, err := ConnectToDb()
	if err != nil {
		log.Println(err)
	}
	defer db.Close()
	userPass, err := hashPassword("123456")
	if err != nil {
		log.Println(err)
	}
	_, err = db.Exec("INSERT INTO users_logins (username, password) VALUES ('test1', ?)", userPass)
	if err != nil {
		log.Println(err)
	}
	_, err = db.Exec("INSERT INTO users_logins (username, password) VALUES ('test2', ?)", userPass)
	if err != nil {
		log.Println(err)
	}
}
