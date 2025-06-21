package model

type ErrorResponse struct {
	Error string `json:"error"`
}

type Campsite struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Address     string  `json:"address"`
	Description string  `json:"description"`
	Facilities  string  `json:"facilities"`
	Price       int     `json:"price"`
	ImageURL    string  `json:"image_url"`
	Latitude    float64 `json:"latitude"`
	Longitude   float64 `json:"longitude"`
	CreatedAt   string  `json:"created_at"`
}
