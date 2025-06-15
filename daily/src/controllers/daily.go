package controllers

import (
	"eman/passport/daily/src/api"
	"eman/passport/daily/src/models"
	"eman/passport/daily/src/services"
	"github.com/labstack/echo"
	"net/http"
)

type Response struct {
	Status bool  `json:"status"`
	Data   *Data `json:"data"`
}

type Data struct {
	Consecutive int              `json:"consecutive"`
	AvailableAt int              `json:"available_at"`
	Rewards     []*models.Reward `json:"rewards,omitempty"`
	Popup       *models.Event    `json:"popup,omitempty"`
}


func (self *Handler) Daily(c echo.Context) error {

	self.caller = c.Get("caller").(*api.Caller)

	launch, err := services.GetCurrentLaunch(self.caller, self.Container.LaunchesRepository)
	if err != nil {
		return err
	}

	var rewards []*models.Reward
	if launch.IsFresh() {
		rewards = services.GetRewards(launch.GetConsecutivePeriods())
		services.GiveRewards(self.Container.RewardsRepository, rewards, self.caller.Id)
	}

	event, err := self.Container.EventsRepository.GetEvent()
	if err != nil {
		return err
	}

	response := &Response{
		Status: launch.IsFresh(),
		Data: &Data{
			Consecutive: launch.GetConsecutivePeriods(),
			AvailableAt: launch.NextAvailableAt(),
			Rewards:     rewards,
			Popup:       event,
		},
	}

	return c.JSON(http.StatusOK, response)
}
