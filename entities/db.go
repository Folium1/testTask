package entities

import (
	"database/sql"
	"test/config"
	"test/models"
)

type Storage struct {
	db *sql.DB
}

// New returns a new Storage instance
func New() (*Storage, error) {
	db, err := config.ConnectToDb()
	if err != nil {
		return nil, err
	}
	return &Storage{db}, nil
}

// SaveImage saves an image's data to the database
func (s *Storage) SaveImage(image models.Image) error {
	query := "INSERT INTO images (user_id, image_path, image_url) VALUES (?, ?, ?)"
	_, err := s.db.Exec(query, image.UserID, image.Path, image.URL)
	if err != nil {
		return err
	}
	return nil
}

// GetImagesByUserID returns all user's images
func (s *Storage) GetImagesByUserID(userID int) ([]models.Image, error) {
	var images []models.Image
	query := "SELECT id, user_id, image_path, image_url FROM images WHERE user_id = ?"
	rows, err := s.db.Query(query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// Iterate over the rows
	for rows.Next() {
		var image models.Image
		err := rows.Scan(&image.ID, &image.UserID, &image.Path, &image.URL)
		if err != nil {
			return nil, err
		}
		images = append(images, image)
	}
	return images, nil
}

// GetUserByUsername returns a user data from db by username
func (s *Storage) GetUserByUsername(username string) (*models.User, error) {
	var user models.User
	query := "SELECT id, username, password FROM users WHERE username = ?"
	row := s.db.QueryRow(query, username)
	err := row.Scan(&user.ID, &user.Username, &user.Password)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (s *Storage) SaveUserLog(user models.User) error {
	query := "INSERT INTO users_logins (username, password) VALUES (?, ?)"
	_, err := s.db.Exec(query, user.Username, user.Password)
	if err != nil {
		return err
	}
	return nil
}
