package services

import (
	"eman/passport/daily/src/api"
	"eman/passport/daily/src/models"
	"eman/passport/daily/src/repositories"
)

func GetCurrentLaunch(caller *api.Caller, launches *repositories.LaunchesRepository) (*models.Launch, error) {

	last, err := launches.GetLastLaunch(caller.Id)
	if err != nil {
		return nil, err
	}

	//Case: First ever launch by this user
	if last == nil {
		return launches.AddLaunch(caller.Id, 1)
	}

	//Case: Already logged a launch for current period
	if last.IsFromCurrentPeriod() {
		return last, nil
	}

	//Case: New starting point
	if last.IsFromExpiredPeriod() {
		return launches.AddLaunch(caller.Id, 1)
	}

	//Case: Consecutive period
	return launches.AddLaunch(caller.Id, last.GetConsecutivePeriods()+1)
}
