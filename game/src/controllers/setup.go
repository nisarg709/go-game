package controllers

import (
	"github.com/labstack/echo"
	"net/http"
	"eman/passport/game/src/api"
	"eman/passport/game/src/games"
)

func (self *Handler) Setup(c echo.Context) error {

	auth := c.Get("user").(api.User)
	id := c.Param("id");

	// Find: The Play record
	play, err := self.Container.PlaysRepository.Find(id)
	if err != nil {
		return err;
	}

	// Rule: Check if Play belongs to User
	if !play.BelongsTo(auth.Id) {
		return c.JSON(http.StatusForbidden, api.NewForbidden())
	}

	// Rule: Check if Play already completed
	if play.IsCompleted() {
		return c.JSON(http.StatusBadRequest, api.NewCustomError("Game already completed", 12100))
	}

	// Action: Respond with setup data
	rules, _ := games.NewGameRules(play, self.Container)

	data, err := rules.SetupData()
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, api.Data{Data: data})
}
