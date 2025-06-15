package controllers

import (
	"github.com/labstack/echo"
	"net/http"
	"eman/passport/game/src/api"
	"eman/passport/game/src/services"
	"eman/passport/game/src/games"
	"github.com/globalsign/mgo/bson"
)

func (h *Handler) Start(c echo.Context) error {

	auth := c.Get("user").(api.User)

	// Rule: Check if parameters resolve a valid game/level record
	resolver, err := games.NewGameResolver(h.Container.GamesRepository, c.Param("game"), c.Param("diff"))

	if err != nil {
		return c.JSON(http.StatusBadRequest, api.NewCustomError("No such game", 12100))
	}

	// Rule: Check if user has played more than the allowed number of times
	playsToday, err := h.Container.PlaysRepository.CountToday(resolver.GameUid, resolver.LevelUid, auth.Id)
	if err != nil {
		return err;
	}

	if !resolver.Game.CanPlayMore(playsToday) {
		return c.JSON(http.StatusBadRequest, api.NewCustomError("Free play limit exceeded", 12100))
	}

	// Find: The User record
	user, err := h.Container.UsersRepository.Find(auth.Id)

	if err != nil {
		return c.JSON(http.StatusForbidden, api.NewUnauthorized())
	}

	// Rule: Check if user has enough miles to play this city and level
	city, err := h.Container.GamesRepository.FindCityByDifficulty(resolver.LevelUid)

	if err != nil {
		return err;
	}

	if city.Requires > user.Miles {
		return c.JSON(http.StatusBadRequest, api.NewCustomError("Not enough miles", 12100))
	}

	// @TODO OPTIMIZE Check if player has another game in progress. Decide what to do in this case

	// Action: Create a new Play record
	play, err := services.CreateNewPlay(h.Container, auth.Id, resolver.GameUid, resolver.LevelUid)

	if err != nil {
		return err;
	}

	return c.JSON(http.StatusCreated, api.Data{
		Data: struct {
			Id         bson.ObjectId `json:"id"`
			Type       string        `json:"type"`
			Difficulty int           `json:"difficulty"`
		}{
			Id:         play.Id,
			Type:       play.Type,
			Difficulty: play.Difficulty,
		},
	})
}
