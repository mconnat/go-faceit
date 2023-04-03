package client

import (
	"github.com/mconnat/faceit-api-client/pkg/models"
)

func (c *FaceITClient) SearchChampionships(name string, params map[string]interface{}) (models.Championships, error) {
	if params == nil {
		params = make(map[string]interface{}, 1)
	}
	params["name"] = name
	championships, err := Get(models.Championships{}, c, "/search/championships", params)
	if err != nil {
		return models.Championships{}, err
	}
	return championships, nil
}

func (c *FaceITClient) SearchHubs(name string, params map[string]interface{}) (models.Hubs, error) {
	if params == nil {
		params = make(map[string]interface{}, 1)
	}
	params["name"] = name
	hubs, err := Get(models.Hubs{}, c, "/search/hubs", params)
	if err != nil {
		return models.Hubs{}, err
	}
	return hubs, nil
}

func (c *FaceITClient) SearchOrganizers(name string, params map[string]interface{}) (models.Organizers, error) {
	if params == nil {
		params = make(map[string]interface{}, 1)
	}
	params["name"] = name
	organizers, err := Get(models.Organizers{}, c, "/search/organizers", params)
	if err != nil {
		return models.Organizers{}, err
	}
	return organizers, nil
}

func (c *FaceITClient) SearchPlayers(nickname string, params map[string]interface{}) (models.Players, error) {
	if params == nil {
		params = make(map[string]interface{}, 1)
	}
	params["nickname"] = nickname
	players, err := Get(models.Players{}, c, "/search/players", params)
	if err != nil {
		return models.Players{}, err
	}
	return players, nil
}

func (c *FaceITClient) SearchTeams(nickname string, params map[string]interface{}) (models.Teams, error) {
	if params == nil {
		params = make(map[string]interface{}, 1)
	}
	params["nickname"] = nickname
	teams, err := Get(models.Teams{}, c, "/search/teams", params)
	if err != nil {
		return models.Teams{}, err
	}
	return teams, nil
}

func (c *FaceITClient) SearchTournaments(name string, params map[string]interface{}) (models.Tournaments, error) {
	if params == nil {
		params = make(map[string]interface{}, 1)
	}
	params["nickname"] = name
	tournaments, err := Get(models.Tournaments{}, c, "/search/tournaments", params)
	if err != nil {
		return models.Tournaments{}, err
	}
	return tournaments, nil
}
