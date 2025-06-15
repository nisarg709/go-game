package games

import "eman/passport/game/src/models"

type ContinuousGame struct {
	game *models.Game
}

func (self *ContinuousGame) Resume(play *models.Play, payload map[string]interface{}) error {
	return nil
}

func (self *ContinuousGame) CostOfResume() int {
	return self.game.ResumeCost
}

func (self *ContinuousGame) CostOfHelp() int {
	return self.game.HelpCost
}

func (self *ContinuousGame) CanUseHelp(payload map[string]interface{}) bool {
	return false
}

func (self *ContinuousGame) Checkpoint(play *models.Play, payload interface{}) (bool, error) {
	return true, nil
}

func (self *ContinuousGame) HelpData(payload map[string]interface{}) (interface{}, error) {
	return nil, nil
}

func (self *ContinuousGame) SetupData() (interface{}, error) {
	return nil, nil
}
