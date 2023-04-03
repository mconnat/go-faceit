package client

import (
	"fmt"
	"github.com/mconnat/faceit-api-client/pkg/models"
)

func (c *FaceITClient) GetLeaderboardsByChampionshipID(championshipID string, params map[string]interface{}) (models.GroupLeaderboards, error) {
	leaderboards, err := Get(models.GroupLeaderboards{}, c, fmt.Sprintf("/leaderboards/championships/%s", championshipID), params)
	if err != nil {
		return models.GroupLeaderboards{}, err
	}
	return leaderboards, nil
}

func (c *FaceITClient) GetGroupLeaderboardByChampionshipID(championshipID string, group int64, params map[string]interface{}) (models.Leaderboards, error) {
	groupLeaderboard, err := Get(models.Leaderboards{}, c, fmt.Sprintf("/leaderboards/championships/%s/groups/%d", championshipID, group), params)
	if err != nil {
		return models.Leaderboards{}, err
	}
	return groupLeaderboard, nil
}

func (c *FaceITClient) GetLeaderboardsByHubID(hubID string, params map[string]interface{}) (models.GroupLeaderboards, error) {
	leaderboards, err := Get(models.GroupLeaderboards{}, c, fmt.Sprintf("/leaderboards/hubs/%s", hubID), params)
	if err != nil {
		return models.GroupLeaderboards{}, err
	}
	return leaderboards, nil
}

func (c *FaceITClient) GetGeneralLeaderboardsByHubID(hubID string, params map[string]interface{}) (models.Leaderboards, error) {
	groupLeaderboard, err := Get(models.Leaderboards{}, c, fmt.Sprintf("/leaderboards/hubs/%s/general", hubID), params)
	if err != nil {
		return models.Leaderboards{}, err
	}
	return groupLeaderboard, nil
}

func (c *FaceITClient) GetSeasonLeaderboardsByHubID(hubID string, season int64, params map[string]interface{}) (models.Leaderboards, error) {
	groupLeaderboard, err := Get(models.Leaderboards{}, c, fmt.Sprintf("/leaderboards/hubs/%s/seasons/%d", hubID, season), params)
	if err != nil {
		return models.Leaderboards{}, err
	}
	return groupLeaderboard, nil
}

func (c *FaceITClient) GetLeaderboardByID(leaderboardID string, params map[string]interface{}) (models.Leaderboards, error) {
	groupLeaderboard, err := Get(models.Leaderboards{}, c, fmt.Sprintf("/leaderboards/%s", leaderboardID), params)
	if err != nil {
		return models.Leaderboards{}, err
	}
	return groupLeaderboard, nil
}
