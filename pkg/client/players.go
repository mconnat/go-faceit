package client

import (
	"fmt"
	"github.com/mconnat/go-faceit/pkg/models"
)

func (c *FaceITClient) GetPlayer(params map[string]interface{}) (models.Player, error) {
	player, err := Get(models.Player{}, c, "/players", params)
	if err != nil {
		return models.Player{}, err
	}
	return player, nil
}

func (c *FaceITClient) GetPlayerByID(playerID string, params map[string]interface{}) (models.Player, error) {
	player, err := Get(models.Player{}, c, fmt.Sprintf("/players/%s", playerID), params)
	if err != nil {
		return models.Player{}, err
	}
	return player, nil
}

func (c *FaceITClient) GetPlayerHistory(playerID string, game string, params map[string]interface{}) (models.PlayerHistory, error) {
	if params == nil {
		params = make(map[string]interface{}, 1)
	}
	params["game"] = game
	player, err := Get(models.PlayerHistory{}, c, fmt.Sprintf("/players/%s/history", playerID), params)
	if err != nil {
		return models.PlayerHistory{}, err
	}
	return player, nil
}

func (c *FaceITClient) GetPlayerHubs(playerID string, params map[string]interface{}) (models.Hubs, error) {
	playerHubs, err := Get(models.Hubs{}, c, fmt.Sprintf("/players/%s/hubs", playerID), params)
	if err != nil {
		return models.Hubs{}, err
	}
	return playerHubs, nil
}

func (c *FaceITClient) GetPlayerStatsByGame(playerID string, game string, params map[string]interface{}) (models.PlayerStats, error) {
	stats, err := Get(models.PlayerStats{}, c, fmt.Sprintf("/players/%s/stats/%s", playerID, game), params)
	if err != nil {
		return models.PlayerStats{}, err
	}
	return stats, nil
}

func (c *FaceITClient) GetPlayersLastMatchesStats(playerID string, game string, params map[string]interface{}) (models.PlayerMatchResponse, error) {
	stats, err := Get(models.PlayerMatchResponse{}, c, fmt.Sprintf("/players/%s/games/%s/stats", playerID, game), params)
	if err != nil {
		return models.PlayerMatchResponse{}, err
	}
	return stats, nil
}

func (c *FaceITClient) GetPlayerTournaments(playerID string, params map[string]interface{}) (models.Tournaments, error) {
	tournaments, err := Get(models.Tournaments{}, c, fmt.Sprintf("/players/%s/tournaments", playerID), params)
	if err != nil {
		return models.Tournaments{}, err
	}
	return tournaments, nil
}
