/*
 * Data API
 *
 * # This API provide access to FACEIT's data
 *
 * API version: 4
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package faceit

type MatchHistory struct {
	CompetitionId string `json:"competition_id,omitempty"`
	CompetitionName string `json:"competition_name,omitempty"`
	CompetitionType string `json:"competition_type,omitempty"`
	FaceitUrl string `json:"faceit_url,omitempty"`
	FinishedAt int64 `json:"finished_at,omitempty"`
	GameId string `json:"game_id,omitempty"`
	GameMode string `json:"game_mode,omitempty"`
	MatchId string `json:"match_id,omitempty"`
	MatchType string `json:"match_type,omitempty"`
	MaxPlayers int64 `json:"max_players,omitempty"`
	OrganizerId string `json:"organizer_id,omitempty"`
	PlayingPlayers []string `json:"playing_players,omitempty"`
	Region string `json:"region,omitempty"`
	Results *MatchResult `json:"results,omitempty"`
	StartedAt int64 `json:"started_at,omitempty"`
	Status string `json:"status,omitempty"`
	Teams map[string]HistoryFaction `json:"teams,omitempty"`
	TeamsSize int64 `json:"teams_size,omitempty"`
}
