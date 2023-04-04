package client

import (
	"fmt"
	"github.com/mconnat/go-faceit/pkg/models"
)

func (c *FaceITClient) GetGames(params map[string]interface{}) (models.Games, error) {
	games, err := Get(models.Games{}, c, "/games", params)
	if err != nil {
		return models.Games{}, err
	}
	return games, nil
}

func (c *FaceITClient) GetGameByID(gameID string, params map[string]interface{}) (models.Game, error) {
	game, err := Get(models.Game{}, c, fmt.Sprintf("/games/%s", gameID), params)
	if err != nil {
		return models.Game{}, err
	}
	return game, nil
}

func (c *FaceITClient) GetGameParentByGameID(gameID string, params map[string]interface{}) (models.Game, error) {
	game, err := Get(models.Game{}, c, fmt.Sprintf("/games/%s/parent", gameID), params)
	if err != nil {
		return models.Game{}, err
	}
	return game, nil
}
