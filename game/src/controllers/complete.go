package controllers

import (
	"eman/passport/game/src/api"
	"eman/passport/game/src/models"
	"eman/passport/game/src/services"
	"eman/passport/game/src/services/checksum"
	"github.com/labstack/echo"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"net/http"
	"os"

	"github.com/go-ozzo/ozzo-validation"
)

type CompleteRequest struct {
	Success bool
	Score   int
	Helps   int
	Payload string
}

func (r CompleteRequest) Validate() error {
	return validation.ValidateStruct(&r,
		validation.Field(&r.Success, validation.NotNil),
		validation.Field(&r.Score, validation.NotNil),
		validation.Field(&r.Helps, validation.NotNil),
		validation.Field(&r.Payload, validation.Required),
	)
}

func (self *Handler) Complete(c echo.Context) error {

	auth := c.Get("user").(api.User)
	id := c.Param("id");

	//Action: Validate Request
	request := new(CompleteRequest)
	if err := c.Bind(request); err != nil {
		return c.JSON(http.StatusForbidden, api.NewValidationError())
	}

	if err := request.Validate(); err != nil {
		return c.JSON(http.StatusForbidden, api.NewValidationError())
	}

	// Find: the Play record
	play, err := self.Container.PlaysRepository.Find(id)

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

	// Rule: Check if game score has the correct checksum
	checksum := checksum.NewChecksum(viper.GetString("games.checksum_key"))
	if !checksum.Verify(request.Payload, id, request.Success, request.Score, play.HelpsCount()) {
		// Checksum is not valid, but we don't want any attacker to know this,
		// so we are just returning a successful status

		log.SetOutput(os.Stdout)
		log.SetFormatter(&log.JSONFormatter{})
		log.WithFields(log.Fields{
			"payload": request.Payload,
			"success": request.Success,
			"score":   request.Score,
			"helps":   request.Helps,
			"id":      id,
		}).Warning("Invalid Checksum sent")

		//return c.JSON(http.StatusOK, api.Status{Status: true})
	}

	// Action: Complete the Play record
	if err = self.Container.PlaysRepository.Complete(play, request.Success, request.Score); err != nil {
		return err
	}

	// Check: If user failed the game -> skip all rewards
	if !request.Success {
		return c.JSON(http.StatusOK, api.Status{Status: true})
	}

	// Find: The User record
	user, err := self.Container.UsersRepository.Find(auth.Id)

	if err != nil {
		return err;
	}

	// Action: Calculate New State
	state := &services.UserState{}
	state.CalculateState(self.Container, user, play)

	// Action: Give User rewards
	err = services.GiveReward(self.Container, user, state)
	if err != nil {
		return err
	}

	// Action: Check for Achievements
	achievements, err := services.UnlockAchievements(self.Container, user, state)

	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK,
		api.Status{
			Status: true, Data: struct {
				Rewards      []models.Reward `json:"rewards"`
				Achievements []string        `json:"achievements"`
			}{
				Rewards:      state.Rewards,
				Achievements: achievements,
			},
		})
}
