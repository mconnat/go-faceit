package client

import (
	"fmt"
	"github.com/mconnat/go-faceit/pkg/models"
)

type GetPlayerParams struct {
	Nickname     string `query:"nickname"`
	Game         string `query:"game"`
	GamePlayerID string `query:"game_player_id"`
}

func (c *FaceITClient) GetPlayer(params *GetPlayerParams) (models.Player, error) {
	player, err := Get(models.Player{}, c, "/players", params)
	if err != nil {
		return models.Player{}, err
	}
	return player, nil
}

func (c *FaceITClient) GetPlayerByID(playerID string) (models.Player, error) {
	player, err := Get(models.Player{}, c, fmt.Sprintf("/players/%s", playerID), nil)
	if err != nil {
		return models.Player{}, err
	}
	return player, nil
}

type GetPlayerHistoryParams struct {
	Game   string `query:"game"`
	From   int64  `query:"from"`
	To     int64  `query:"to"`
	Offset int    `query:"offset"`
	Limit  int    `query:"limit"`
}

func (c *FaceITClient) GetPlayerHistory(playerID string, params *GetPlayerHistoryParams) (models.PlayerHistory, error) {
	player, err := Get(models.PlayerHistory{}, c, fmt.Sprintf("/players/%s/history", playerID), params)
	if err != nil {
		return models.PlayerHistory{}, err
	}
	return player, nil
}

type GetPlayerHubsParams struct {
	Offset int `query:"offset"`
	Limit  int `query:"limit"`
}

func (c *FaceITClient) GetPlayerHubs(playerID string, params *GetPlayerHubsParams) (models.Hubs, error) {
	playerHubs, err := Get(models.Hubs{}, c, fmt.Sprintf("/players/%s/hubs", playerID), params)
	if err != nil {
		return models.Hubs{}, err
	}
	return playerHubs, nil
}

func (c *FaceITClient) GetPlayerStatsByGame(playerID string, game string) (models.PlayerStats, error) {
	stats, err := Get(models.PlayerStats{}, c, fmt.Sprintf("/players/%s/stats/%s", playerID, game), nil)
	if err != nil {
		return models.PlayerStats{}, err
	}
	return stats, nil
}

type GetPlayerTournamentsParams struct {
	Offset int `query:"offset"`
	Limit  int `query:"limit"`
}

func (c *FaceITClient) GetPlayerTournaments(playerID string, params *GetPlayerTournamentsParams) (models.Tournaments, error) {
	tournaments, err := Get(models.Tournaments{}, c, fmt.Sprintf("/players/%s/tournaments", playerID), params)
	if err != nil {
		return models.Tournaments{}, err
	}
	return tournaments, nil
}
