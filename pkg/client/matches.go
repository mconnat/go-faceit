package client

import (
	"fmt"
	"github.com/mconnat/go-faceit/pkg/models"
)

func (c *FaceITClient) GetMatchByID(matchID string, params map[string]interface{}) (models.Match, error) {
	match, err := Get(models.Match{}, c, fmt.Sprintf("/matches/%s", matchID), params)
	if err != nil {
		return models.Match{}, err
	}
	return match, nil
}

func (c *FaceITClient) GetMatchStats(matchID string, params map[string]interface{}) (models.MatchStats, error) {
	stats, err := Get(models.MatchStats{}, c, fmt.Sprintf("/matches/%s/stats", matchID), params)
	if err != nil {
		return models.MatchStats{}, err
	}
	return stats, nil
}
