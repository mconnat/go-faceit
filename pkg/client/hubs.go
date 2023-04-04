package client

import (
	"fmt"
	"github.com/mconnat/go-faceit/pkg/models"
)

func (c *FaceITClient) GetHubsByID(hubID string, params map[string]interface{}) (models.Hub, error) {
	if hubID == "" {
		return models.Hub{}, fmt.Errorf("hubID undefined")
	}
	hub, err := Get(models.Hub{}, c, fmt.Sprintf("/hubs/%s", hubID), params)
	if err != nil {
		return models.Hub{}, err
	}
	return hub, nil
}

func (c *FaceITClient) GetHubMatches(hubID string, params map[string]interface{}) (models.Matches, error) {
	hub, err := Get(models.Matches{}, c, fmt.Sprintf("/hubs/%s/matches", hubID), params)
	if err != nil {
		return models.Matches{}, err
	}
	return hub, nil
}

func (c *FaceITClient) GetHubMembers(hubID string, params map[string]interface{}) (models.HubMembers, error) {
	members, err := Get(models.HubMembers{}, c, fmt.Sprintf("/hubs/%s/members", hubID), params)
	if err != nil {
		return models.HubMembers{}, err
	}
	return members, nil
}

func (c *FaceITClient) GetHubRoles(hubID string, params map[string]interface{}) (models.HubRoles, error) {
	roles, err := Get(models.HubRoles{}, c, fmt.Sprintf("/hubs/%s/roles", hubID), params)
	if err != nil {
		return models.HubRoles{}, err
	}
	return roles, nil
}

func (c *FaceITClient) GetHubRules(hubID string, params map[string]interface{}) (models.HubRules, error) {
	rules, err := Get(models.HubRules{}, c, fmt.Sprintf("/hubs/%s/rules", hubID), params)
	if err != nil {
		return models.HubRules{}, err
	}
	return rules, nil
}

func (c *FaceITClient) GetHubStats(hubID string, params map[string]interface{}) (models.HubStats, error) {
	rules, err := Get(models.HubStats{}, c, fmt.Sprintf("/hubs/%s/stats", hubID), params)
	if err != nil {
		return models.HubStats{}, err
	}
	return rules, nil
}
