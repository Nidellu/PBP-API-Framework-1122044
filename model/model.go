package model

type Game struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Genre string `json:"genre"`
	Price string `json:"price"`
}

type Transaction struct {
	ID     int `json:"id"`
	UserID int `json:"user_id"`
	GameID int `json:"game_id"`
}

type User struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
