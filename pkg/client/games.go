package client

import (
	"fmt"

	"github.com/mconnat/go-faceit/pkg/models"
)

type GetGamesParams struct {
	Offset int `query:"offset"`
	Limit  int `query:"limit"`
}

func (c *FaceITClient) GetGames(params *GetGamesParams) (models.Games, error) {
	games, err := Get(models.Games{}, c, "/games", params)
	if err != nil {
		return models.Games{}, err
	}
	return games, nil
}

func (c *FaceITClient) GetGameByID(gameID string) (models.Game, error) {
	game, err := Get(models.Game{}, c, fmt.Sprintf("/games/%s", gameID), nil)
	if err != nil {
		return models.Game{}, err
	}
	return game, nil
}

func (c *FaceITClient) GetGameParentByGameID(gameID string) (models.Game, error) {
	game, err := Get(models.Game{}, c, fmt.Sprintf("/games/%s/parent", gameID), nil)
	if err != nil {
		return models.Game{}, err
	}
	return game, nil
}
