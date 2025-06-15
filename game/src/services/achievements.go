package services

import (
	"eman/passport/game/src/container"
	"eman/passport/game/src/models"
	"errors"
	"strings"
)

func UnlockAchievements(container *container.Container, user *models.User, state *UserState) ([]string, error) {

	var achievements []string

	if state.Milestone != nil && state.Milestone.Unlock != "" {
		badge := strings.Replace(state.Milestone.Unlock, "city_", "unlock_", 1)
		achievements = append(achievements, badge)
	}

	if state.Perfect {
		badge, err := preparePerfectBadgeName(state.Game, state.Difficulty)
		if err == nil {
			achievements = append(achievements, badge)
		}
	}

	// Always add this badge it wll be filtered if already achieved
	achievements = append(achievements, "win_first_game_ever")

	achievements = filterAlreadyGiven(achievements, state.Badges)
	if len(achievements) > 0 {
		err := container.UsersRepository.AddBadges(user, achievements)
		if err != nil {
			return nil, err
		}
	}

	return achievements, nil;
}

func UnlockTokenAchievements(container *container.Container, user *models.User) ([]string, error) {
	var achievements []string

	total, err := container.UsersRepository.GetSpentTokens(user.Id)
	if err != nil {
		return nil, err
	}

	badge := prepareSpentTokenBadgeName(total)
	if badge != "" {
		achievements = append(achievements, badge)
	}

	achievements = filterAlreadyGiven(achievements, user.Badges)
	if len(achievements) > 0 {
		err := container.UsersRepository.AddBadges(user, achievements)
		if err != nil {
			return nil, err
		}
	}

	return achievements, nil;
}

func preparePerfectBadgeName(game string, difficulty int) (string, error) {

	var level string

	switch difficulty {
	case 1, 2, 3, 4, 5, 6:
		level = "easy"
	case 7, 8, 9, 10, 11, 12:
		level = "medium"
	case 13, 14, 15, 16, 17, 18:
		level = "hard"
	default:
		return "", errors.New("invalid difficulty")
	}
	return game + "_perfect_" + level, nil
}

func prepareSpentTokenBadgeName(total int) string {

	if total > 1000 {
		return "spent_tokens_500"
	} else if total > 500 {
		return "spent_tokens_200"
	} else if total > 100 {
		return "spent_tokens_100"
	} else if total > 30 {
		return "spent_tokens_30"
	}

	return ""
}

func filterAlreadyGiven(achievements []string, badges []string) []string {

	var filtered []string

	for _, a := range achievements {
		found := false
		for _, b := range badges {
			if a == b {
				found = true
				break
			}
		}
		if !found {
			filtered = append(filtered, a);
		}
	}

	return filtered
}
