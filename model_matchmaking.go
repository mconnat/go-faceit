/*
 * Data API
 *
 * # This API provide access to FACEIT's data
 *
 * API version: 4
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package faceit

// Matchmaking A detailed representation of a matchmaking
type Matchmaking struct {
	Game string `json:"game,omitempty"`
	Icon string `json:"icon,omitempty"`
	Id string `json:"id,omitempty"`
	LeagueId string `json:"league_id,omitempty"`
	LongDescription string `json:"long_description,omitempty"`
	Name string `json:"name,omitempty"`
	Queues []MatchmakingQueue `json:"queues,omitempty"`
	Region string `json:"region,omitempty"`
	ShortDescription string `json:"short_description,omitempty"`
}
