package client

import (
	"fmt"
	"github.com/mconnat/faceit-api-client/pkg/models"
)

func (c *FaceITClient) GetTournaments(params map[string]interface{}) (models.Tournaments, error) {
	tournaments, err := Get(models.Tournaments{}, c, "/tournaments", params)
	if err != nil {
		return models.Tournaments{}, err
	}
	return tournaments, nil
}

func (c *FaceITClient) GetTournamentByID(tournamentID string, params map[string]interface{}) (models.Tournaments, error) {
	tournaments, err := Get(models.Tournaments{}, c, fmt.Sprintf("/tournaments/%s", tournamentID), params)
	if err != nil {
		return models.Tournaments{}, err
	}
	return tournaments, nil
}

func (c *FaceITClient) GetTournamentBrackets(tournamentID string, params map[string]interface{}) (models.TournamentBrackets, error) {
	brackets, err := Get(models.TournamentBrackets{}, c, fmt.Sprintf("/tournaments/%s/brackets", tournamentID), params)
	if err != nil {
		return models.TournamentBrackets{}, err
	}
	return brackets, nil
}

func (c *FaceITClient) GetTournamentMatches(tournamentID string, params map[string]interface{}) (models.Matches, error) {
	matches, err := Get(models.Matches{}, c, fmt.Sprintf("/tournaments/%s/matches", tournamentID), params)
	if err != nil {
		return models.Matches{}, err
	}
	return matches, nil
}

func (c *FaceITClient) GetTournamentTeams(tournamentID string, params map[string]interface{}) (models.TournamentTeams, error) {
	teams, err := Get(models.TournamentTeams{}, c, fmt.Sprintf("/tournaments/%s/teams", tournamentID), params)
	if err != nil {
		return models.TournamentTeams{}, err
	}
	return teams, nil
}
