package client

import (
	"fmt"
	"github.com/mconnat/go-faceit/pkg/models"
)

func (c *FaceITClient) GetChampionshipsByGameID(gameID string, params map[string]interface{}) (models.Championships, error) {
	if gameID == "" {
		return models.Championships{}, fmt.Errorf("gameID undefined")
	}
	if params == nil {
		params = make(map[string]interface{}, 1)
	}
	params["game"] = gameID
	championships, err := Get(models.Championships{}, c, "/championships", params)
	if err != nil {
		return models.Championships{}, err
	}
	return championships, nil
}

type GetChampionshipByIDParams struct {
	Expanded []string `query:"expanded"`
}

func (c *FaceITClient) GetChampionshipByID(championshipID string, params *GetChampionshipByIDParams) (models.Championship, error) {
	if championshipID == "" {
		return models.Championship{}, fmt.Errorf("championshipID undefined")
	}
	championships, err := Get(models.Championship{}, c, fmt.Sprintf("/championships/%s", championshipID), params)
	if err != nil {
		return models.Championship{}, err
	}
	return championships, nil
}

func (c *FaceITClient) GetMatchesByChampionshipID(championshipID string, params map[string]interface{}) (models.Matches, error) {
	if championshipID == "" {
		return models.Matches{}, fmt.Errorf("championshipID undefined")
	}
	matches, err := Get(models.Matches{}, c, fmt.Sprintf("/championships/%s/matches", championshipID), params)
	if err != nil {
		return models.Matches{}, err
	}
	return matches, nil
}

func (c *FaceITClient) GetResultsByChampionshipID(championshipID string, params map[string]interface{}) (models.ChampionshipResults, error) {
	if championshipID == "" {
		return models.ChampionshipResults{}, fmt.Errorf("championshipID undefined")
	}
	results, err := Get(models.ChampionshipResults{}, c, fmt.Sprintf("/championships/%s/results", championshipID), params)
	if err != nil {
		return models.ChampionshipResults{}, err
	}
	return results, nil
}

func (c *FaceITClient) GetSubscriptionsByChampionshipID(championshipID string, params map[string]interface{}) (models.ChampionshipSubscriptions, error) {
	if championshipID == "" {
		return models.ChampionshipSubscriptions{}, fmt.Errorf("championshipID undefined")
	}
	results, err := Get(models.ChampionshipSubscriptions{}, c, fmt.Sprintf("/championships/%s/subscriptions", championshipID), params)
	if err != nil {
		return models.ChampionshipSubscriptions{}, err
	}
	return results, nil
}
