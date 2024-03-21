package controller

import (
	"net/http"
	"strconv"

	m "modul/model"

	"github.com/labstack/echo/v4"
)

func GetGames(c echo.Context) error {
	db := InitDB()
	defer db.Close()

	rows, err := db.Query("SELECT id, name, genre, price FROM games")
	if err != nil {
		return err
	}
	defer rows.Close()

	var games []m.Game
	for rows.Next() {
		var game m.Game
		err := rows.Scan(&game.ID, &game.Name, &game.Genre, &game.Price)
		if err != nil {
			return err
		}
		games = append(games, game)
	}
	return c.JSON(http.StatusOK, games)
}

func GetGame(c echo.Context) error {
	db := InitDB()
	defer db.Close()
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}

	var game m.Game
	err = db.QueryRow("SELECT id, name, genre, price FROM games WHERE id = ?", id).Scan(&game.ID, &game.Name, &game.Genre, &game.Price)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, game)
}

func CreateGame(c echo.Context) error {
	db := InitDB()
	defer db.Close()

	name := c.Param("name")
	genre := c.Param("genre")
	price := c.Param("price")

	newGame := m.Game{
		Name:  name,
		Genre: genre,
		Price: price,
	}

	_, err := db.Exec("INSERT INTO games (name, genre, price) VALUES (?, ?, ?)", name, genre, price)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, newGame)
}

func UpdateGame(c echo.Context) error {
	db := InitDB()
	defer db.Close()

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}
	name := c.Param("name")
	genre := c.Param("genre")
	price := c.Param("price")

	_, err = db.Exec("UPDATE games SET name = ?, genre = ?, price = ? WHERE id = ?", name, genre, price, id)
	if err != nil {
		return err
	}

	updatedGame := m.Game{
		ID:    id,
		Name:  name,
		Genre: genre,
		Price: price,
	}

	return c.JSON(http.StatusOK, updatedGame)
}

func DeleteGame(c echo.Context) error {
	db := InitDB()
	defer db.Close()
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}

	_, err = db.Exec("DELETE FROM games WHERE id = ?", id)
	if err != nil {
		return err
	}

	return c.NoContent(http.StatusOK)
}
