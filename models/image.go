package models

type Image struct {
	ID     int    `json:"id"`
	UserID int    `json:"user_id"`
	Path   string `json:"image_path"`
	URL    string `json:"image_url"`
}
