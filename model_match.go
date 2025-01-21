/*
 * Data API
 *
 * # This API provide access to FACEIT's data
 *
 * API version: 4
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package go-faceit

type Match struct {
	BestOf int64 `json:"best_of,omitempty"`
	BroadcastStartTime int64 `json:"broadcast_start_time,omitempty"`
	BroadcastStartTimeLabel string `json:"broadcast_start_time_label,omitempty"`
	CalculateElo bool `json:"calculate_elo,omitempty"`
	ChatRoomId string `json:"chat_room_id,omitempty"`
	CompetitionId string `json:"competition_id,omitempty"`
	CompetitionName string `json:"competition_name,omitempty"`
	CompetitionType string `json:"competition_type,omitempty"`
	ConfiguredAt int64 `json:"configured_at,omitempty"`
	DemoUrl []string `json:"demo_url,omitempty"`
	DetailedResults []DetailedMatchResult `json:"detailed_results,omitempty"`
	FaceitUrl string `json:"faceit_url,omitempty"`
	FinishedAt int64 `json:"finished_at,omitempty"`
	Game string `json:"game,omitempty"`
	Group int64 `json:"group,omitempty"`
	MatchId string `json:"match_id,omitempty"`
	OrganizerId string `json:"organizer_id,omitempty"`
	Region string `json:"region,omitempty"`
	Results *MatchResult `json:"results,omitempty"`
	Round int64 `json:"round,omitempty"`
	ScheduledAt int64 `json:"scheduled_at,omitempty"`
	StartedAt int64 `json:"started_at,omitempty"`
	Status string `json:"status,omitempty"`
	Teams map[string]Faction `json:"teams,omitempty"`
	Version int64 `json:"version,omitempty"`
	Voting interface{} `json:"voting,omitempty"`
}
