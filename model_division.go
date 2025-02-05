/*
 * Data API
 *
 * # This API provide access to FACEIT's data
 *
 * API version: 4
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package faceit

type Division struct {
	Assets *Assets `json:"assets,omitempty"`
	// The type of the division. Can be nested or classic. Nested means that the division has tiers, classic is without tiers.
	ConfigType string `json:"config_type,omitempty"`
	LeaderboardConfig *LeaderboardConfig `json:"leaderboard_config,omitempty"`
	// The leaderboards of the division
	Leaderboards []string `json:"leaderboards,omitempty"`
	// Max ELO for a user to be placed in this division after placement matches
	MaxElo int64 `json:"max_elo,omitempty"`
	// Min ELO for a user to be placed in this division after placement matches
	MinElo int64 `json:"min_elo,omitempty"`
	// The name of the division.
	Name string `json:"name,omitempty"`
	// The tiers of the division
	Tiers []Tier `json:"tiers,omitempty"`
	// The type of the division.
	Type_ string `json:"type,omitempty"`
}
