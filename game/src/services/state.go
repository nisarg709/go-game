package services

import (
	"eman/passport/game/src/container"
	"eman/passport/game/src/models"
	"eman/passport/game/src/repositories"
)

type UserState struct {
	InitialMiles int
	UpdatedMiles int
	Badges       []string
	Perfect      bool
	Game         string
	Difficulty   int

	Milestone *models.Milestone
	Rewards   []models.Reward
}

func (self *UserState) CalculateState(container *container.Container, user *models.User, play *models.Play) (error) {

	boosts := container.UsersRepository.GetBoosts(user.Id)
	game, err := container.GamesRepository.FindGame(play.Type, play.Difficulty)

	if err != nil {
		return err
	}

	self.Badges = user.Badges
	self.Game = play.Type
	self.Difficulty = play.Difficulty
	self.Perfect = play.HelpsCount() == 0

	awardedMiles := miles(game, boosts)

	if awardedMiles > 0 {
		self.InitialMiles = user.Miles
		self.UpdatedMiles = user.Miles + awardedMiles

		self.Rewards = append(self.Rewards, models.Reward{"miles", awardedMiles})
	}

	milestone := findMilestoneReached(container.MilestonesRepository, self.InitialMiles, self.UpdatedMiles)

	if milestone != nil {
		self.Milestone = milestone

		additional := rewards(milestone, boosts)
		self.Rewards = append(self.Rewards, additional...)
	}

	return nil
}

func rewards(milestone *models.Milestone, boosts models.BoostCollection) []models.Reward {
	var rewards []models.Reward

	for i, v := range milestone.Rewards {

		rewards = append(rewards,  models.Reward{v.Type, v.Qty})

		if "ticket" == v.Type && boosts.Contains("EntriesX2") {
			rewards[i].Qty *= 2
		}
	}

	return rewards
}

func miles(game *models.Game, boosts models.BoostCollection) int {

	milesAwarded := game.GetRewardedMiles();
	if boosts.Contains("MilesX2") {
		milesAwarded = milesAwarded * 2
	}

	return milesAwarded
}

func findMilestoneReached(repository *repositories.MilestonesRepository, oldUserMiles int, newUserMiles int) *models.Milestone {

	milestone, err := repository.Achieved(oldUserMiles, newUserMiles)

	if err != nil {
		return nil
	}

	return milestone;
}
