package services

import (
	"eman/passport/game/src/container"
	"eman/passport/game/src/models"
)

func GiveReward(container *container.Container, user *models.User, state *UserState) (error) {

	for _, v := range state.Rewards {
		//TODO: Create a record for the reward being given
		switch v.Type {
		case "miles":
			container.UsersRepository.AddMiles(user, v.Qty)
		case "ticket":
			container.UsersRepository.AddTicket(user, v.Qty)
		}
	}

	return nil;
}
