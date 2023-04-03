package client

import (
	"fmt"
	"github.com/mconnat/faceit-api-client/pkg/models"
)

func (c *FaceITClient) GetOrganizer(params map[string]interface{}) (models.Organizer, error) {
	organizer, err := Get(models.Organizer{}, c, "/organizers", params)
	if err != nil {
		return models.Organizer{}, err
	}
	return organizer, nil
}

func (c *FaceITClient) GetOrganizerByID(organizerID string, params map[string]interface{}) (models.Organizer, error) {
	organizer, err := Get(models.Organizer{}, c, fmt.Sprintf("/organizers/%s", organizerID), params)
	if err != nil {
		return models.Organizer{}, err
	}
	return organizer, nil
}

func (c *FaceITClient) GetOrganizerChampionships(organizerID string, params map[string]interface{}) (models.Championships, error) {
	championships, err := Get(models.Championships{}, c, fmt.Sprintf("/organizers/%s/championships", organizerID), params)
	if err != nil {
		return models.Championships{}, err
	}
	return championships, nil
}

func (c *FaceITClient) GetOrganizerGames(organizerID string, params map[string]interface{}) (models.Games, error) {
	games, err := Get(models.Games{}, c, fmt.Sprintf("/organizers/%s/games", organizerID), params)
	if err != nil {
		return models.Games{}, err
	}
	return games, nil
}

func (c *FaceITClient) GetOrganizerHubs(organizerID string, params map[string]interface{}) (models.Hubs, error) {
	hubs, err := Get(models.Hubs{}, c, fmt.Sprintf("/organizers/%s/hubs", organizerID), params)
	if err != nil {
		return models.Hubs{}, err
	}
	return hubs, nil
}

func (c *FaceITClient) GetOrganizerTournaments(organizerID string, params map[string]interface{}) (models.Tournaments, error) {
	tournaments, err := Get(models.Tournaments{}, c, fmt.Sprintf("/organizers/%s/tournaments", organizerID), params)
	if err != nil {
		return models.Tournaments{}, err
	}
	return tournaments, nil
}
