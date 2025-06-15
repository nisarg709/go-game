package services

import (
	"eman/passport/daily/src/models"
	"eman/passport/daily/src/repositories"
)

func GetRewards(periods int) []*models.Reward {

	var rewards []*models.Reward

	if periods <= 0 {
		return rewards
	}

	if periods%5 == 0 {
		return append(rewards, &models.Reward{"ticket", 1})
	}

	var qty int

	// CASE 1: ceil( x/y ) * 5
	//qty = (1 + ((periods - 1) / 5)) * 5

	// CASE 2: ( ceil( Log1.75(periods) ) + 1 ) * 5
	//qty = (int(math.Ceil(math.Log(float64(periods))/math.Log(1.75))) + 1) * 5;

	// CASE 3: ( ceil( Log1.75(periods) ) + 1 ) * 5
	switch check := periods; {
	case check < 5:
		qty = periods
	case check < 10:
		qty = 4;
	case check < 25:
		qty = 5
	case check < 40:
		qty = 6
	case check < 60:
		qty = 7
	case check < 80:
		qty = 8
	case check < 100:
		qty = 9
	default:
		qty = 10
	}

	return append(rewards, &models.Reward{"token", qty})
}

func GiveRewards(repo *repositories.RewardsRepository, rewards []*models.Reward, userId string) {
	for _, reward := range rewards {

		if reward.Type == "ticket" {
			repo.GiveTickets(userId, reward.Qty)
			continue
		}

		if reward.Type == "token" {
			repo.GiveTokens(userId, reward.Qty)
			continue
		}
	}
}
