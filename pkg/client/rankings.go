package client

import (
	"fmt"
	"github.com/mconnat/go-faceit/pkg/models"
)

func (c *FaceITClient) GetGameRankingByRegion(gameID string, region string, params map[string]interface{}) (models.Rankings, error) {
	rankings, err := Get(models.Rankings{}, c, fmt.Sprintf("/rankings/games/%s/regions/%s", gameID, region), params)
	if err != nil {
		return models.Rankings{}, err
	}
	return rankings, nil
}

func (c *FaceITClient) GetPlayerRankingByRegion(gameID string, region string, playerID string, params map[string]interface{}) (models.Rankings, error) {
	rankings, err := Get(models.Rankings{}, c, fmt.Sprintf("/rankings/games/%s/regions/%s/players/%s", gameID, region, playerID), params)
	if err != nil {
		return models.Rankings{}, err
	}
	return rankings, nil
}
