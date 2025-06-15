package controllers

import (
	"github.com/labstack/echo"
	"net/http"
	"eman/passport/game/src/api"
	"eman/passport/game/src/games"
	"gopkg.in/go-playground/validator.v9"
)

type CheckRequest struct {
	Payload interface{}   `validate:"required"`
}

func (h *Handler) Check(c echo.Context) error {

	auth := c.Get("user").(api.User)
	id := c.Param("id");

	//Action: Validate Request
	validator := validator.New()

	request := new(CheckRequest)
	if err := c.Bind(request); err != nil {
		return err
	}

	if err := validator.Struct(request); err != nil {
		return c.JSON(http.StatusForbidden, api.NewValidationError())
	}


	// Find: The Play record
	play, err := h.Container.PlaysRepository.Find(id)
	if err != nil {
		return c.JSON(http.StatusForbidden, api.NewNotFoundError())
	}

	// Rule: Check if Play belongs to User
	if !play.BelongsTo(auth.Id) {
		return c.JSON(http.StatusForbidden, api.NewForbidden())
	}

	// Rule: Check if Play already completed
	if play.IsCompleted() {
		return c.JSON(http.StatusBadRequest, api.NewCustomError("Game already completed", 12100))
	}

	// Action: Respond to checkpoint
	rules, _ := games.NewGameRules(play, h.Container)
	status, err := rules.Checkpoint(play, request.Payload);

	if err != nil {
		return err;
	}

	return c.JSON(http.StatusOK, api.Status{Status: status})
}
