package controllers

import (
	"eman/passport/game/src/api"
	"eman/passport/game/src/games"
	"eman/passport/game/src/services"
	"github.com/labstack/echo"
	"net/http"
)

func (h *Handler) Help(c echo.Context) error {

	auth := c.Get("user").(api.User)
	id := c.Param("id");

	m := echo.Map{}
	if err := c.Bind(&m); err != nil {
		//Ignore empty POST body
	}

	var payload map[string]interface{}
	if m["payload"] != nil {
		payload = m["payload"].(map[string]interface{})
	}

	// Find: The Play record
	play, err := h.Container.PlaysRepository.Find(id)
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

	// Rule: Check if game rules allow using help
	rules, _ := games.NewGameRules(play, h.Container)

	if !rules.CanUseHelp(payload) {
		return c.JSON(http.StatusBadRequest, api.NewCustomError("Can't use help", 12100))
	}

	// Find: The User record
	user, err := h.Container.UsersRepository.Find(auth.Id)
	if err != nil {
		return err;
	}

	// Rule: Check if User has enough tokens
	if !user.HasEnoughTokens(rules.CostOfHelp()) {
		return c.JSON(http.StatusBadRequest, api.NewCustomError("Not enough tokens", 12100))
	}

	data, err := rules.HelpData(payload)
	if err != nil {
		return err
	}

	// Action: Charge tokens
	if err := h.Container.UsersRepository.Charge(user, rules.CostOfHelp(), "help:"+play.Id.Hex()); err != nil {
		return err
	}

	// Action: Note the Help usage
	if err := h.Container.PlaysRepository.Help(play); err != nil {
		return err
	}

	// Action: Check for Achievements
	achievements, err := services.UnlockTokenAchievements(h.Container, user)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK,
		api.Status{
			Status: true,
			Data: struct {
				Help         interface{} `json:"help"`
				Achievements []string    `json:"achievements"`
			}{
				Help:         data,
				Achievements: achievements,
			},
		},
	)
}
