package services

import (
	"eman/passport/daily/src/models"
	"testing"
)

func TestInvalidPeriod(t *testing.T) {

	var rewards []*models.Reward

	rewards = GetRewards(0);
	if len(rewards) > 0 {
		t.Errorf("Unexpected number of rewards, got: %d, expect: %d.", len(rewards), 0)
	}
}

func TestDay3(t *testing.T) {

	var rewards []*models.Reward
	expected := models.Reward{"token", 3};

	rewards = GetRewards(3);
	if len(rewards) != 1 {
		t.Errorf("Unexpected number of rewards, got: %d, expect: %d.", len(rewards), 1)
	}

	if *rewards[0] != expected {
		t.Errorf("Wrong type of reward, got: %+v, expected: %+v", *rewards[0], expected)
	}
}

func TestDay5(t *testing.T) {

	var rewards []*models.Reward
	expected := models.Reward{"ticket", 1};

	rewards = GetRewards(5);
	if len(rewards) != 1 {
		t.Errorf("Unexpected number of rewards, got: %d, expect: %d.", len(rewards), 1)
	}

	if *rewards[0] != expected {
		t.Errorf("Wrong type of reward, got: %+v, expected: %+v", *rewards[0], expected)
	}
}

func TestDay7(t *testing.T) {

	var rewards []*models.Reward
	expected := models.Reward{"token", 4};

	rewards = GetRewards(7);
	if len(rewards) != 1 {
		t.Errorf("Unexpected number of rewards, got: %d, expect: %d.", len(rewards), 1)
	}

	if *rewards[0] != expected {
		t.Errorf("Wrong type of reward, got: %+v, expected: %+v", *rewards[0], expected)
	}
}

func TestDay15(t *testing.T) {

	var rewards []*models.Reward
	expected := models.Reward{"ticket", 1};

	rewards = GetRewards(15);
	if len(rewards) != 1 {
		t.Errorf("Unexpected number of rewards, got: %d, expect: %d.", len(rewards), 1)
	}

	if *rewards[0] != expected {
		t.Errorf("Wrong type of reward, got: %+v, expected: %+v", *rewards[0], expected)
	}
}

func TestDay77(t *testing.T) {

	var rewards []*models.Reward
	expected := models.Reward{"token", 8};

	rewards = GetRewards(77);
	if len(rewards) != 1 {
		t.Errorf("Unexpected number of rewards, got: %d, expect: %d.", len(rewards), 1)
	}

	if *rewards[0] != expected {
		t.Errorf("Wrong type of reward, got: %+v, expected: %+v", *rewards[0], expected)
	}
}

func TestDay253(t *testing.T) {

	var rewards []*models.Reward
	expected := models.Reward{"token", 10};

	rewards = GetRewards(253);
	if len(rewards) != 1 {
		t.Errorf("Unexpected number of rewards, got: %d, expect: %d.", len(rewards), 1)
	}

	if *rewards[0] != expected {
		t.Errorf("Wrong type of reward, got: %+v, expected: %+v", *rewards[0], expected)
	}
}

func TestDay366(t *testing.T) {

	var rewards []*models.Reward
	expected := models.Reward{"token", 10};

	rewards = GetRewards(366);
	if len(rewards) != 1 {
		t.Errorf("Unexpected number of rewards, got: %d, expect: %d.", len(rewards), 1)
	}

	if *rewards[0] != expected {
		t.Errorf("Wrong type of reward, got: %+v, expected: %+v", *rewards[0], expected)
	}
}