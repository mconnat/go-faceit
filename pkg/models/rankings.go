package models

type Rankings struct {
	End   int `json:"end"`
	Items []struct {
		Country        string `json:"country"`
		FaceitElo      int    `json:"faceit_elo"`
		GameSkillLevel int    `json:"game_skill_level"`
		Nickname       string `json:"nickname"`
		PlayerId       string `json:"player_id"`
		Position       int    `json:"position"`
	} `json:"items"`
	Position int `json:"position,omitempty"`
	Start    int `json:"start"`
}
